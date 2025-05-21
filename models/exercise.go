package models

import "time"

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
