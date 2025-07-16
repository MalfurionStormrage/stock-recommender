package services

import (
	models "backend/models/stock"
	repository "backend/repository/stock"
	"fmt"
	"log"
)

//---------------------------------//
//---- LÓGICA DE NEGOCIO STOCK  --//
//-------------------------------//

// -- Get_all_stocks retorna todos los registros de la tabla stock --//
func Get_all_stocks() ([]map[string]interface{}, error) {
	return repository.Get_all_stocks()
}

// -- SearchStock busca registros filtrando por ticker o next_page --//
func SearchStock(filter string) ([]map[string]interface{}, error) {
	return repository.SearchStock(filter)
}

// -- CrearStock guarda un nuevo registro de stock en la base de datos --//
func CrearStock(stock models.StockItemCreateRequest) error {
	// Validar que no exista un registro con el mismo ticker y company
	registros, err := repository.SearchStock(stock.Ticker)
	if err != nil {
		return err
	}
	for _, reg := range registros {
		if reg["ticker"] == stock.Ticker && reg["company"] == stock.Company {
			return fmt.Errorf("ya existe un registro con ese ticker")
		}
	}
	return repository.CrearStock(stock)
}

// -- ActualizarStock actualiza un registro existente por su ID --//
func ActualizarStock(id int, stock models.StockItemCreateRequest) error {
	// Validar existencia
	existe, err := GetStockByID(id)
	if err != nil {
		return err
	}
	if existe == nil {
		return fmt.Errorf("no existe un registro con ese ID")
	}
	return repository.ActualizarStock(id, stock)
}

// -- EliminarStock elimina un registro existente por su ID --//
func EliminarStock(id int) error {
	// Validar existencia
	registro, err := GetStockByID(id)
	if err != nil {
		log.Println("Error:", err)
	}
	if registro == nil {
		return fmt.Errorf("no existe un registro con ese ID")
	}

	return repository.EliminarStock(id)
}

// -- GetBestStockRecommendation obtiene la mejor recomendación de stock --//
func GetBestStockRecommendation() ([]map[string]interface{}, error) {
	// Obtener todos los registros
	stocks, err := repository.GetBestStockRecommendation()
	if err != nil {
		return nil, fmt.Errorf("error al obtener recomendaciones de stock: %v", err)
	}

	// Validar que haya registros
	if len(stocks) == 0 {
		return nil, fmt.Errorf("no se encontraron recomendaciones de stock")
	}

	return stocks, nil
}

// -- GetStockByID busca un registro de stock por su ID y lo retorna como un map[string]interface{} --//
func GetStockByID(id int) (map[string]interface{}, error) {
	return repository.GetStockByID(id)
}
