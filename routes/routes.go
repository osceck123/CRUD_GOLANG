package routes

import (
	"crud-gin/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	// Rutas del CRUD de usuarios
	router.POST("/users", controllers.CreateUser)
	router.GET("/users", controllers.GetUsers)
	router.GET("/users/:id", controllers.GetUserByID)
	router.PUT("/users/:id", controllers.UpdateUser)
	router.DELETE("/users/:id", controllers.DeleteUser)

	// Ruta de WebSocket
	router.GET("/ws", controllers.HandleConnections)

	// Iniciar la gestión de mensajes
	go controllers.HandleMessages()
}
