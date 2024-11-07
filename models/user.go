package models

// User representa el perfil del usuario
type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"` // Se omite en la respuesta para seguridad
}
