package main

import (
	"crud-gin/database"
	"crud-gin/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// Registrar rutas
	routes.RegisterRoutes(router)

	//definimos la conexion a la db como llamado

	database.ConnectDB()

	// Iniciar el servidor en el puerto 8080
	router.Run(":8181")
}
