package routes

import (
	controllers "backend/controllers/stock"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(r *gin.Engine) {
	//-- IMPLEMENTAMOS SWAGGER --//
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//-------------------------//
	//-- URL base de la api --//
	//-----------------------//
	api := r.Group("/api/v1")
	{
		//------------------------------//
		//-- ENDPOINT DE LA API REST --//
		//----------------------------//

		//-- OBTENER TODO LOS REGISTROS --//
		api.GET("/stock", controllers.Get_all_stocks)

		//-- BUSCAR REGISTROS POR TICKERT O NEXT_PAGE --//
		api.GET("/stock/search", controllers.SearchStock)

		//-- CREAR NUEVO REGISTRO --//
		api.POST("/stock", controllers.CrearStock)

		//-- ACTUALIZAR UN REGISTRO EXISTENTE --//
		api.PUT("/stock/:id", controllers.ActualizarStock)

		//-- ELIMINAR UN REGISTRO EXISTENTE --//
		api.DELETE("/stock/:id", controllers.EliminarStock)

		//-- OBTENER MEJOR RECOMENDACIÓN DE STOCK --//
		api.GET("/stock/best-recommendation", controllers.GetBestStockRecommendation)

	}
}
