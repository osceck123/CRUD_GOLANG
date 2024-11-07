package models

// Estructura del mensaje
type Message struct {
	Sender  string `json:"sender"`
	Content string `json:"content"`
}
