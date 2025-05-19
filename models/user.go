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

// Exercise représente un exercice dans un programme
type Exercise struct {
	ID         string    `json:"id"`
	ProgramID  string    `json:"program_id"`
	Name       string    `json:"name"`
	Sets       int       `json:"sets"`
	Reps       int       `json:"reps"`
	Weight     float64   `json:"weight"`      // en kg
	WeightUnit string    `json:"weight_unit"` // "kg" ou "lbs"
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
