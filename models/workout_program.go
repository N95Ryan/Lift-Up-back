package models

import "time"

// WorkoutProgram représente un programme d'entraînement
type WorkoutProgram struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"` // "SPLIT" ou "PPL"
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
