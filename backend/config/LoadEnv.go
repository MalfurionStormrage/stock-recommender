package config

import (
	"log"
	"sync"

	"github.com/joho/godotenv"
)

// -- Declaración de la variable once para garantizar la carga única del archivo .env. --//
var once sync.Once

// -- Cargar variables de desarrollo --//
func LoadEnv() {
	once.Do(func() {
		// Intenta cargar el .env desde la ruta actual
		err := godotenv.Load()
		if err != nil {
			// Si falla, intenta cargar desde el directorio superior
			err = godotenv.Load("../.env")
			if err != nil {
				log.Println("⚠️ No se pudo cargar el archivo .env, usando variables del sistema.")
			}
		}
	})
}
