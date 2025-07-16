package database

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	models "backend/models/stock"
)

func MigrarStockDesdeAPI() {
	//Url de api externa
	baseURL := "https://8j5baasof2.execute-api.us-west-2.amazonaws.com/production/swechallenge/list"
	nextPage := ""
	//Cliente http
	client := &http.Client{}

	// Crear tabla una sola vez
	if _, err := Pool.Exec(context.Background(), CreateTableStock); err != nil {
		log.Fatal("❌ Error al crear la tabla stock:", err)
	}

	for {
		// Construir la URL
		var url string
		if nextPage == "" {
			url = baseURL
		} else {
			url = fmt.Sprintf("%s?next_page=%s", baseURL, nextPage)
		}
		log.Println("🌐 Solicitando:", url)

		// Preparar request
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatalf("❌ Error al crear request HTTP: %v", err)
		}
		//Headers
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+os.ExpandEnv("${AUTHORIZATION}"))

		// Ejecutar request
		resp, err := client.Do(req)
		if err != nil {
			log.Fatalf("❌ Error en la solicitud HTTP: %v", err)
		}
		defer resp.Body.Close()

		//Validar estado de respuesta de api externa
		if resp.StatusCode != http.StatusOK {
			log.Fatalf("❌ Código de estado no OK: %d", resp.StatusCode)
		}

		// Decodificar respuesta
		var apiResp models.APIResponse
		if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
			log.Fatalf("❌ Error al parsear JSON: %v", err)
		}

		// Verificar si ya migramos todos los registros de esta página (next_pages)
		var countEnBD int
		err = Pool.QueryRow(context.Background(),
			`SELECT COUNT(*) FROM stock WHERE next_page = $1`, apiResp.NextPage,
		).Scan(&countEnBD)
		if err != nil {
			log.Fatal("❌ Error al contar registros en BD:", err)
		}

		cantidadEnAPI := len(apiResp.Items)
		if countEnBD == cantidadEnAPI && cantidadEnAPI > 0 {
			log.Printf("ℹ️ Página %s ya migrada (%d registros). Saltando.\n", apiResp.NextPage, countEnBD)
			nextPage = apiResp.NextPage
			if nextPage == "" {
				log.Println("✅ No hay más páginas. Migración finalizada.")
				break
			}
			continue
		}

		// Insertar datos
		insertados := 0
		for _, item := range apiResp.Items {
			_, err := Pool.Exec(context.Background(), Create_stock,
				item.Ticker,
				strings.ReplaceAll(strings.ReplaceAll(item.TargetFrom, "$", ""), ",", ""),
				strings.ReplaceAll(strings.ReplaceAll(item.TargetTo, "$", ""), ",", ""),
				item.Company, item.Action, item.Brokerage,
				item.RatingFrom, item.RatingTo, item.Time, nextPage,
				apiResp.NextPage,
			)
			if err != nil {
				log.Printf("⚠️ Error al insertar %s: %v\n", item.Ticker, err)
				continue
			}
			insertados++
		}

		log.Printf("✅ Insertados %d registros desde página %s\n", insertados, apiResp.NextPage)

		// Terminar si no hay siguiente página
		if apiResp.NextPage == "" {
			log.Println("✅ No hay más páginas. Migración finalizada.")
			break
		}

		nextPage = apiResp.NextPage
	}

	log.Println("🎉 Migración completada exitosamente.")
}
