package controllers

import (
	models "backend/models/stock"
	services "backend/services/stock"
	"backend/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//---------------------------//
//-- CONTROLADOR DE STOCK --//
//-------------------------//

// -- ENDPOINT PARA OBTENER TODO LOS REGISTROS ---//
// Get_all_stocks godoc
// @Summary Obtener todos los registros
// @Description Retorna todos los registros de la tabla stock
// @Tags Stock
// @Produce json
// @Success 200 {object} models.SuccessResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/stock [get]
func Get_all_stocks(c *gin.Context) {
	//-- RETORNAMOS REGISTROS DEL SERVICIO --//
	stock, err := services.Get_all_stocks()
	//-- RETORNAMOS ERRORES EN CASO DE HABERLOS --//
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, fmt.Sprintf("Error al obtener stock: %v", err))
		return
	}
	//-- RETORNAR REGISTROS --//
	utils.RespondSuccess(c, stock)
}

// -- ENDPOINT PARA FILTRAR POR TICKER O GRUPO POR NEXT_PAGE --//
// SearchStock godoc
// @Summary Filtrar registros de stock
// @Description Filtra los registros por ticker o next_page usando el query param "filter"
// @Tags Stock
// @Produce json
// @Param filter query string true "Valor a buscar en ticker o next_page"
// @Success 200 {object} models.SuccessResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /stock/search [get]
func SearchStock(c *gin.Context) {
	//-- RETORNAMOS REGISTROS PASANDO COMO PARAMETRO EL VALOR DE BUSQUEDA --//
	stock, err := services.SearchStock(c.Query("filter"))
	//-- RETORNAMOS ERRORES EN CASO DE HABERLOS --//
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, fmt.Sprintf("Error al obtener stock: %v", err))
		return
	}
	//-- RETORNAR REGISTROS --//
	utils.RespondSuccess(c, stock)
}

// -- ENDPOINT PARA CREAR NUEVO REGISTRO --//
// CrearStock godoc
// @Summary Crear nuevo stock
// @Description Crea un nuevo registro en la base de datos
// @Tags Stock
// @Accept json
// @Produce json
// @Param stock body models.StockItemCreateRequest true "Datos del nuevo stock"
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/stock [post]
func CrearStock(c *gin.Context) {

	//-- CREAMOS ESTRUCTURA DEL MODELO, PAR ESTRUCTURAR LOS VALORES RECIBIDOS --//
	var req models.StockItemCreateRequest

	//-- RETORNAMOS ERRORES EN CASO DE HABERLOS --//
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "JSON inválido")
		return
	}

	//-- UTILIZAMOS EL SERVICIO PARA GUARDAR LOS DATOS RECIBIDOS --//
	if err := services.CrearStock(req); err != nil {
		utils.RespondError(c, http.StatusInternalServerError, fmt.Sprintf("Error al insertar stock: %v", err))
		return
	}

	//-- CONFIRMAMOS QUE SE GUARDARON LOS DATOS --//
	utils.RespondSuccess(c, gin.H{"message": "Stock creado exitosamente"})
}

// -- ENDPOINT PARA ACTUALIZAR REGISTROS EXISTENTES --//
// ActualizarStock godoc
// @Summary Actualizar un registro existente
// @Description Actualiza un registro de stock existente por ID
// @Tags Stock
// @Accept json
// @Produce json
// @Param id path int true "ID del stock a actualizar"
// @Param body body models.StockItemCreateRequest true "Datos del stock a actualizar"
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.ErrorResponse "ID inválido o JSON malformado"
// @Failure 500 {object} models.ErrorResponse "Error al actualizar el stock"
// @Router /api/v1/stock/{id} [put]
func ActualizarStock(c *gin.Context) {
	//-- OBTENEMOS EL ID DEL REGISTROS QUE ACTUALIZAREMOS --//
	id, err := strconv.Atoi(c.Param("id"))
	//-- RETORNAMOS ERRORES EN CASO DE HABERLOS --//
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "ID inválido")
		return
	}
	//-- CREAMOS ESTRUCTURA DEL MODELO, PAR ESTRUCTURAR LOS VALORES RECIBIDOS --//
	var req models.StockItemCreateRequest
	//-- RETORNAMOS ERRORES EN CASO DE HABERLOS --//
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "JSON inválido")
		return
	}
	//-- UTILIZAMOS EL SERVICIO PARA ACTUALIZAR LOS DATOS, PASANDO COMO PARAMETROS EL ID Y LOS DATOS RECIBIDOS EN EL BODY --//
	//-- EN CASO DE ERROR, RETORNAMOS MENSAJE CON LA INFORMACION --//
	if err := services.ActualizarStock(id, req); err != nil {
		utils.RespondError(c, http.StatusInternalServerError, fmt.Sprintf("Error al actualizar stock: %v", err))
		return
	}
	//-- RETORNAMOS MENSAJE DE CONFIRMACION DE LA ACTUALIZACION --//
	utils.RespondSuccess(c, gin.H{"message": "Stock actualizado correctamente"})

}

// -- ENDPOINT PARA ELIMINAR REGISTROS EXISTENTES --//
// EliminarStock godoc
// @Summary Eliminar un registro de stock
// @Description Elimina un registro existente usando su ID
// @Tags Stock
// @Produce json
// @Param id path int true "ID del stock a eliminar"
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.ErrorResponse "ID inválido"
// @Failure 500 {object} models.ErrorResponse "Error al eliminar el stock"
// @Router /api/v1/stock/{id} [delete]
func EliminarStock(c *gin.Context) {
	//-- RECIBIMOS ID DEL REGISTRO QUE VAMOS A ELIMINAR --//
	id, err := strconv.Atoi(c.Param("id"))
	//-- RETORNAMOS ERRORES EN CASO DE HABERLOS --//
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "ID inválido")
		return
	}
	//-- UTILIZAMOS EL SERVICIO PARA ELIMINAR UN REGISTRO EXISTENTA PASANDO POR PARAMETRO EL ID --//
	//-- EN CASO DE ERROR, RETORNAMOS MENSAJE CON LA INFORMACION --//
	if err := services.EliminarStock(id); err != nil {
		utils.RespondError(c, http.StatusInternalServerError, fmt.Sprintf("Error al eliminar stock: %v", err))
		return
	}
	//-- RETORNAMOS MENSAJE DE CONFIRMACION DE LA ELIMINACION --//
	utils.RespondSuccess(c, gin.H{"message": "Stock eliminado correctamente"})
}

func GetBestStockRecommendation(c *gin.Context) {
	//-- OBTENEMOS LA MEJOR RECOMENDACION DE STOCK --//
	stocks, err := services.GetBestStockRecommendation()
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, fmt.Sprintf("Error al obtener recomendaciones de stock: %v", err))
		return
	}
	//-- RETORNAMOS LAS RECOMENDACIONES --//
	utils.RespondSuccess(c, stocks)
}
