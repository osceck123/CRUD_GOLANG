package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3" // Importa el driver de SQLite
)

var DB *sql.DB

// Conectar a la base de datos SQLite
func ConnectDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./data/users.db")
	if err != nil {
		log.Fatal("Error al conectar a SQLite:", err)
	}

	// Verificar la conexión
	err = DB.Ping()
	if err != nil {
		log.Fatal("No se pudo conectar a la base de datos:", err)
	}

	fmt.Println("Conectado a SQLite!")

	// Crear la tabla 'users' si no existe
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	);
	`
	_, err = DB.Exec(query)
	if err != nil {
		log.Fatal("Error al crear la tabla de usuarios:", err)
	}
}
