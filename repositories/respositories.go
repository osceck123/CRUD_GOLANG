package repositories

import (
	"crud-gin/database"
	"crud-gin/models"
	"database/sql"
	"fmt"
)

// Crear usuario
func CreateUser(usuario models.User) (sql.Result, error) {
	query := "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"
	result, err := database.DB.Exec(query, usuario.Name, usuario.Email, usuario.Password)
	if err != nil {
		return nil, fmt.Errorf("error al insertar usuario: %v", err)
	}
	return result, nil
}

// Obtener todos los usuarios
func GetUsers() ([]models.User, error) {
	var users []models.User
	query := "SELECT id, name, email FROM users"
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener usuarios: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, fmt.Errorf("error al escanear usuario: %v", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error al iterar sobre usuarios: %v", err)
	}

	return users, nil
}

// Obtener usuario por ID
func GetUserByID(id int) (*models.User, error) {
	var user models.User
	query := "SELECT id, name, email FROM users WHERE id = ?"
	err := database.DB.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Usuario no encontrado
		}
		return nil, fmt.Errorf("error al obtener usuario: %v", err)
	}
	return &user, nil
}

// Actualizar usuario por ID
func UpdateUser(id int, updateUser models.User) (sql.Result, error) {
	query := "UPDATE users SET name = ?, email = ?, password = ? WHERE id = ?"
	result, err := database.DB.Exec(query, updateUser.Name, updateUser.Email, updateUser.Password, id)
	if err != nil {
		return nil, fmt.Errorf("error al actualizar usuario: %v", err)
	}
	return result, nil
}

// Eliminar usuario por ID
func DeleteUserByID(id int) (sql.Result, error) {
	query := "DELETE FROM users WHERE id = ?"
	result, err := database.DB.Exec(query, id)
	if err != nil {
		return nil, fmt.Errorf("error al eliminar usuario: %v", err)
	}
	return result, nil
}
