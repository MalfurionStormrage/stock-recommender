package utils

import (
	models "backend/models/res"

	"github.com/gin-gonic/gin"
)

//--------------------------------------------------------------//
//--- METODOS UTILIZADOS PARA RESPONDER EN LOS CONTROLADORES --//
//------------------------------------------------------------//

// RespondSuccess envía una respuesta exitosa con el estado "success".
func RespondSuccess(c *gin.Context, data interface{}) {
	response := models.SuccessResponse{
		Status: models.StatusSuccess,
		Data:   data,
	}
	c.JSON(200, response)
}

// RespondError envía una respuesta de error con el estado "error".
func RespondError(c *gin.Context, code int, message string) {
	response := models.ErrorResponse{
		Status:  models.StatusError,
		Message: message,
	}
	c.JSON(code, response)
}
