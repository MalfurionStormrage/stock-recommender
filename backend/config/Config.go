package config

import (
	"log"
	"os"
	"strings"
)

func GetDSN() string {
	//establecer cadena de conexion
	dsn := os.ExpandEnv("postgresql://${DB_USER}:${DB_PASSWORD}@${DB_NAME}-12891.j77." +
		"aws-us-east-1.cockroachlabs.cloud:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSLMODE}")
	// Validar que no hayan valores sin reemplazar
	if strings.Contains(dsn, "${") {
		log.Fatal("❌ Faltan variables de entorno en el DSN")
	}
	return dsn
}
