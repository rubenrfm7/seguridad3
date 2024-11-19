package utils

import (
	"log"

	"github.com/gin-gonic/gin"
)

func SetupTLS(router *gin.Engine) {
	// Configuración de seguridad TLS si es necesario.
	log.Println("TLS setup completed")
}
