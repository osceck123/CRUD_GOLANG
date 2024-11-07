package controllers

import (
	"net/http"
	"strconv"

	"crud-gin/models"
	"crud-gin/services"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// Crear usuario
func CreateUser(c *gin.Context) {
	var user models.User

	// Intentamos obtener el JSON del cuerpo de la petición y lo vinculamos a user
	if err := c.ShouldBindJSON(&user); err != nil {
		// Si hay un error en el cuerpo de la petición, devolvemos un 400
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Llamamos al servicio para crear el usuario
	userID, err := services.CreateUser(user)
	if err != nil {
		// Si ocurre un error al crear el usuario, devolvemos un 500
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Si todo sale bien, devolvemos un 200 con el ID del usuario creado
	c.JSON(http.StatusOK, gin.H{
		"message": "Usuario creado exitosamente",
		"user_id": userID, // Devuelve el ID generado por SQLite
	})
}

// Obtener todos los usuarios
func GetUsers(c *gin.Context) {
	usersResult, err := services.GetUsers()
	if err != nil {
		// Si ocurre un error al obtener los usuarios, devolvemos un 500
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Si todo sale bien, devolvemos un 200 con todos los usuarios
	c.JSON(http.StatusOK, usersResult)
}

// Obtener un usuario por ID
func GetUserByID(c *gin.Context) {
	// Convertimos el parámetro `id` de string a int
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Llamamos al servicio para obtener el usuario
	user, err := services.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Usuario no encontrado"})
		return
	}

	// Si se encuentra el usuario, lo devolvemos en la respuesta
	c.JSON(http.StatusOK, user)
}

// Actualizar un usuario
func UpdateUser(c *gin.Context) {
	// Convertimos el parámetro `id` de string a int
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var updateUser models.User
	// Verificamos si hay errores en la solicitud (JSON inválido)
	if err := c.ShouldBindJSON(&updateUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Llamamos al servicio para actualizar el usuario
	err = services.UpdateUser(id, updateUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar usuario"})
		return
	}

	// Si todo sale bien, devolvemos el mensaje de éxito
	c.JSON(http.StatusOK, gin.H{"message": "Usuario actualizado exitosamente"})
}

// Eliminar un usuario por ID
func DeleteUser(c *gin.Context) {
	// Convertimos el parámetro `id` de string a int
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Llamamos al servicio para eliminar el usuario
	err = services.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar usuario"})
		return
	}

	// Si todo sale bien, devolvemos un mensaje de éxito
	c.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado exitosamente"})
}

// WebSocket

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Permitir conexiones desde cualquier origen (por simplicidad)
	},
}

var clients = make(map[*websocket.Conn]bool) // Lista de clientes conectados
var broadcast = make(chan models.Message)    // Canal de difusión de mensajes

// Gestiona la conexión WebSocket
func HandleConnections(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}
	defer ws.Close()
	clients[ws] = true

	for {
		var msg models.Message
		// Leer mensaje
		err := ws.ReadJSON(&msg)
		if err != nil {
			delete(clients, ws)
			break
		}
		// Enviar mensaje al canal de difusión
		broadcast <- msg
	}
}

// Gestiona la difusión de los mensajes a todos los usuarios conectados
func HandleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				client.Close()
				delete(clients, client)
			}
		}
	}
}
