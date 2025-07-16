package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// -- ApplyDefaultMiddlewares aplica middlewares útiles al router de Gin --//
func ApplyDefaultMiddlewares(router *gin.Engine) {
	router.Use(gin.Logger())   // Log de todas las peticiones
	router.Use(gin.Recovery()) // Manejo de errores/panic

	//-- Middleware CORS: permite peticiones desde frontend --//
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
}
