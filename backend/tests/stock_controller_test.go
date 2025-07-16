package tests

import (
	"backend/config"
	"backend/database"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	controllers "backend/controllers/stock"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// init se ejecuta una vez al inicio para preparar el entorno de pruebas
func init() {
	gin.SetMode(gin.TestMode)
	config.LoadEnv()   // Carga variables del entorno (.env)
	database.Connect() // Conecta a la base de datos (PostgreSQL con pgx)
}

// setupRouter define las rutas necesarias para las pruebas
func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/api/v1/stock", controllers.Get_all_stocks)
	r.GET("/stock/search", controllers.SearchStock)
	r.POST("/api/v1/stock", controllers.CrearStock)
	r.PUT("/api/v1/stock/:id", controllers.ActualizarStock)
	r.DELETE("/api/v1/stock/:id", controllers.EliminarStock)
	return r
}

// TestGetAllStocks verifica que se puedan obtener todos los registros
func TestGetAllStocks(t *testing.T) {
	r := setupRouter()
	req, _ := http.NewRequest("GET", "/api/v1/stock", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "status")
}

// TestSearchStock prueba el filtro con query param "filter"
func TestSearchStock(t *testing.T) {
	r := setupRouter()
	req, _ := http.NewRequest("GET", "/stock/search?filter=AAPL", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "status")
}

// TestCrearStock prueba la creación de un nuevo stock con datos válidos
func TestCrearStock(t *testing.T) {
	r := setupRouter()
	payload := map[string]interface{}{
		"ticker":      "AAPL",
		"target_from": "150",
		"target_to":   "180",
		"company":     "Apple Inc.",
		"action":      "Buy",
		"brokerage":   "Goldman Sachs",
		"rating_from": "Neutral",
		"rating_to":   "Buy",
		"time":        "2023-07-01T00:00:00Z",
		"next_page":   "page_2",
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/api/v1/stock", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "Stock creado exitosamente")
}

// TestActualizarStock prueba la actualización de un stock existente
func TestActualizarStock(t *testing.T) {
	r := setupRouter()
	payload := map[string]interface{}{
		"ticker":      "AAPL",
		"target_from": "155",
		"target_to":   "185",
		"company":     "Apple Inc.",
		"action":      "Hold",
		"brokerage":   "Morgan Stanley",
		"rating_from": "Buy",
		"rating_to":   "Sell",
		"time":        "2023-07-02T00:00:00Z",
		"next_page":   "page_3",
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("PUT", "/api/v1/stock/1", bytes.NewBuffer(body)) // Asegúrate que el ID 1 existe
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "Stock actualizado correctamente")
}

// TestEliminarStock prueba la eliminación de un registro existente
func TestEliminarStock(t *testing.T) {
	r := setupRouter()
	req, _ := http.NewRequest("DELETE", "/api/v1/stock/1", nil) // Asegúrate que el ID 1 existe
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "Stock eliminado correctamente")
}

// TestCrearStockInvalidJSON verifica el manejo de un JSON mal formado
func TestCrearStockInvalidJSON(t *testing.T) {
	r := setupRouter()
	body := []byte(`{"ticker": 123}`) // ticker debería ser string
	req, _ := http.NewRequest("POST", "/api/v1/stock", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Contains(t, w.Body.String(), "JSON inválido")
}

// TestActualizarStockInvalidID prueba el manejo de un ID inválido en la ruta
func TestActualizarStockInvalidID(t *testing.T) {
	r := setupRouter()
	body := []byte(`{}`)
	req, _ := http.NewRequest("PUT", "/api/v1/stock/abc", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Contains(t, w.Body.String(), "ID inválido")
}
