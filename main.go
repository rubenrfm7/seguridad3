package main

import (
	"myserver/handlers"
	"myserver/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	// Inicialización del servidor
	r := gin.Default()

	// Configurar HTTPS
	utils.SetupTLS(r)

	// Rutas principales
	r.GET("/version", handlers.GetVersion)
	r.POST("/signup", handlers.SignUp)
	r.POST("/login", handlers.Login)

	// Rutas protegidas con autenticación
	auth := r.Group("/", utils.AuthMiddleware())
	auth.GET("/:username/:doc_id", handlers.GetDocument)
	auth.POST("/:username/:doc_id", handlers.CreateDocument)
	auth.PUT("/:username/:doc_id", handlers.UpdateDocument)
	auth.DELETE("/:username/:doc_id", handlers.DeleteDocument)
	auth.GET("/:username/_all_docs", handlers.GetAllDocuments)

	// Servir
	r.RunTLS(":5000", "cert.pem", "key.pem") // HTTPS
}
