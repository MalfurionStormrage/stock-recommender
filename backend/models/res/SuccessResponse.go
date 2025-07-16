package models

// SuccessResponse representa una respuesta exitosa de la API.
type SuccessResponse struct {
	Status string      `json:"status"` // Estado de la respuesta, normalmente "success"
	Data   interface{} `json:"data"`   // Datos retornados por la API
}

// ErrorResponse representa una respuesta de error de la API.
type ErrorResponse struct {
	Status  string `json:"status"`  // Estado de la respuesta, normalmente "error"
	Message string `json:"message"` // Mensaje descriptivo del error
}

// Constantes para los estados de respuesta
const (
	StatusSuccess = "success"
	StatusError   = "error"
)
