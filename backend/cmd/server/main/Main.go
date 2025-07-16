package main

import (
	"backend/config"
	"backend/database"
	_ "backend/docs"
	"backend/middleware"
	"backend/routes"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

// @title Stock Recommender API
// @version 1.0
// @description API REST para manejar datos de stock
// @host localhost:8080
// @BasePath /api/v1
func main() {
	//-- Asegúrate de que se haya cargado el .env --//
	config.LoadEnv()
	//-- Iniciar pool de conexion a base de datos --//
	database.Connect()
	// database.MigrarStockDesdeAPI()
	//-- Configuracion predeterminada de gin --//
	r := gin.Default()
	//-- middleware para las configuraciones basicas de conexion con el forntend --//
	middleware.ApplyDefaultMiddlewares(r)
	//-- Rutas de la api --//
	routes.SetupRoutes(r)
	//-- Cerrar pool al salir --//
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
		<-sig
		database.Close()
		os.Exit(0)
	}()
	//-- Iniciar servidor con gin framework --//
	r.Run(":8080")
}
