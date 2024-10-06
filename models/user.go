package models

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"-"` // Nota: Usamos `-` para no enviar el campo en respuestas JSON
	Age       int    `json:"age"`
	CreatedAt string `json:"created_at"`
}
