package config

import (
	"os"
)

// SupabaseConfig contient les informations de configuration pour Supabase
type SupabaseConfig struct {
	URL       string
	APIKey    string
	JWTSecret string
}

// NewSupabaseConfig crée une nouvelle instance de configuration Supabase
func NewSupabaseConfig() *SupabaseConfig {
	return &SupabaseConfig{
		URL:       os.Getenv("SUPABASE_URL"),
		APIKey:    os.Getenv("SUPABASE_API_KEY"),
		JWTSecret: os.Getenv("SUPABASE_JWT_SECRET"),
	}
}

// InitSupabase initialise la connexion à Supabase
func InitSupabase() (*SupabaseConfig, error) {
	config := NewSupabaseConfig()
	// TODO: Ajouter la logique de connexion à Supabase
	return config, nil
}
