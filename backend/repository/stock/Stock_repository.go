package repository

import (
	"backend/database"
	models "backend/models/stock"
	"context"
	"strconv"

	"github.com/jackc/pgx/v5"
)

//----------------------------//
//---- REPOSITORIO  STOCK  --//
//--------------------------//

// Get_all_stocks obtiene todos los registros de la tabla stock de forma dinámica y compatible con pgx
func Get_all_stocks() ([]map[string]interface{}, error) {
	rows, err := database.Pool.Query(context.Background(), database.Get_all_stocks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return rowsToMapSlice(rows)
}

// SearchStock filtra los registros por ticker o next_page en la tabla stock de forma dinámica y compatible con pgx
func SearchStock(filter string) ([]map[string]interface{}, error) {
	rows, err := database.Pool.Query(context.Background(), database.FilterStocks, filter, filter)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return rowsToMapSlice(rows)
}

// CrearStock inserta un nuevo registro en la tabla stock
func CrearStock(s models.StockItemCreateRequest) error {
	_, err := database.Pool.Exec(
		context.Background(), database.Create_stock,
		s.Ticker, s.TargetFrom, s.TargetTo,
		s.Company, s.Action, s.Brokerage,
		s.RatingFrom, s.RatingTo, s.Time, s.Origin, s.NextPage,
	)
	return err
}

// ActualizarStock actualiza un registro existente identificado por su ID
func ActualizarStock(id int, s models.StockItemCreateRequest) error {
	_, err := database.Pool.Exec(
		context.Background(), database.Update_stock, s.TargetFrom, s.TargetTo,
		s.Company, s.Action, s.Brokerage,
		s.RatingFrom, s.RatingTo, s.Time, s.Origin, s.NextPage, id,
	)
	return err
}

// EliminarStock elimina un registro existente por su ID
func EliminarStock(id int) error {
	_, err := database.Pool.Exec(context.Background(), database.Delete_stock, id)
	return err
}

// GetBestStockRecommendation obtiene la mejor recomendación de stock
func GetBestStockRecommendation() ([]map[string]interface{}, error) {
	rows, err := database.Pool.Query(context.Background(), database.GetBestStockRecommendation)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return rowsToMapSlice(rows)
}

// rowsToMapSlice convierte los resultados de pgx.Rows en un slice de mapas dinámicos
func rowsToMapSlice(rows pgx.Rows) ([]map[string]interface{}, error) {
	fieldDescriptions := rows.FieldDescriptions()
	columns := make([]string, len(fieldDescriptions))
	for i, fd := range fieldDescriptions {
		columns[i] = string(fd.Name)
	}

	var results []map[string]interface{}

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return nil, err
		}
		rowMap := make(map[string]interface{})
		for i, col := range columns {
			if col == "id" {
				if v, ok := values[i].(int64); ok {
					rowMap[col] = strconv.FormatInt(v, 10)
				} else {
					rowMap[col] = values[i]
				}
			} else {
				rowMap[col] = values[i]
			}
		}
		results = append(results, rowMap)
	}

	return results, nil
}

func GetStockByID(id int) (map[string]interface{}, error) {
	rows, err := database.Pool.Query(context.Background(), "SELECT * FROM stock WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	fieldDescriptions := rows.FieldDescriptions()
	columns := make([]string, len(fieldDescriptions))
	for i, fd := range fieldDescriptions {
		columns[i] = string(fd.Name)
	}

	if rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return nil, err
		}
		rowMap := make(map[string]interface{})
		for i, col := range columns {
			rowMap[col] = values[i]
		}
		return rowMap, nil
	}
	return nil, nil // No encontrado
}
