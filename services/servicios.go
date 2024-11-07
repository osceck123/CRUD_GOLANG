package services

import (
	"crud-gin/models"
	"crud-gin/repositories"
	"fmt"
)

// Crear usuario
func CreateUser(usuario models.User) (int64, error) {
	result, err := repositories.CreateUser(usuario)
	if err != nil {
		return 0, fmt.Errorf("error al crear usuario: %v", err)
	}

	id, err := result.LastInsertId() // Obtener el ID del usuario insertado
	if err != nil {
		return 0, fmt.Errorf("error al obtener ID del usuario: %v", err)
	}

	return id, nil
}

// Obtener todos los usuarios
func GetUsers() ([]models.User, error) {
	users, err := repositories.GetUsers() // Llama a la función de repositorio
	if err != nil {
		return nil, fmt.Errorf("error al obtener usuarios: %v", err)
	}

	return users, nil
}

// Obtener usuario por ID
func GetUserByID(id int) (*models.User, error) {
	user, err := repositories.GetUserByID(id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil // Usuario no encontrado
		}
		return nil, fmt.Errorf("error al obtener usuario: %v", err)
	}

	return user, nil
}

// Actualizar usuario
func UpdateUser(id int, updateUser models.User) error {
	_, err := repositories.UpdateUser(id, updateUser)
	if err != nil {
		return fmt.Errorf("error al actualizar usuario: %v", err)
	}

	return nil
}

// Eliminar usuario por ID
func DeleteUser(id int) error {
	_, err := repositories.DeleteUserByID(id)
	if err != nil {
		return fmt.Errorf("error al eliminar usuario: %v", err)
	}

	return nil
}
