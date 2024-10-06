package main

import (
	"mi-api-go/database"
	"mi-api-go/handlers"

	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Conectar a la base de datos

	log.Println("Iniciando el servidor...")

	database.Connect()

	log.Println("Conexión a la base de datos establecida")

	// Crear un enrutador Gin
	r := gin.Default()

	// Rutas de la API
	r.GET("/users", handlers.GetUsers)
	r.POST("/users", handlers.CreateUser)
	r.PUT("/users/:id", handlers.UpdateUser)
	r.DELETE("/users/:id", handlers.DeleteUser)

	// Correr el servidor en el puerto 8080
	r.Run(":8081")
}
