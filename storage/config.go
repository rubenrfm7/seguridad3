package storage

import (
	"os"
)

var RootPath = "./data" // Directorio raíz para almacenamiento.

func InitStorage() error {
	// Crear el directorio raíz si no existe.
	return os.MkdirAll(RootPath, os.ModePerm)
}
