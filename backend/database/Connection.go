package database

import (
	"backend/config"
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func Connect() {
	//-- Contexto de conexion --//
	ctx := context.Background()
	//-- Obtener url de conexion a base de datos --//
	dsn := config.GetDSN()
	//-- Validacion de conexion a base de datos --//
	var err error
	Pool, err = pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatal("❌ Error al conectar al pool de PostgreSQL:", err)
	}
	//-- confirmacion de conexion --//
	log.Println("✅ Conectado al pool de PostgreSQL correctamente")
}

// -- Metodo para cerrar conexion a base de datos --//
func Close() {
	if Pool != nil {
		Pool.Close()
		log.Println("✅ Pool de conexiones cerrado correctamente")
	}
}
