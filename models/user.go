package models

import "time"

// User représente un utilisateur dans l'application
type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Height    float64   `json:"height"` // en cm
	Weight    float64   `json:"weight"` // en kg
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
