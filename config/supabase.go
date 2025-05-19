package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	supa "github.com/supabase-community/supabase-go"
)

// SupabaseConfig contient les informations de configuration pour Supabase
type SupabaseConfig struct {
	URL       string
	APIKey    string
	JWTSecret string
	Client    *supa.Client
}

// NewSupabaseConfig crée une nouvelle instance de configuration Supabase
func NewSupabaseConfig() *SupabaseConfig {
	// Chargement des variables d'environnement
	if err := godotenv.Load(".env.local"); err != nil {
		fmt.Printf("Erreur lors du chargement du fichier .env.local: %v\n", err)
	}

	return &SupabaseConfig{
		URL:       os.Getenv("SUPABASE_URL"),
		APIKey:    os.Getenv("SUPABASE_API_KEY"),
		JWTSecret: os.Getenv("SUPABASE_JWT_SECRET"),
	}
}

// InitSupabase initialise la connexion à Supabase
func InitSupabase() (*SupabaseConfig, error) {
	config := NewSupabaseConfig()

	// Création du client Supabase
	client, err := supa.NewClient(config.URL, config.APIKey, nil)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de la création du client Supabase: %v", err)
	}

	config.Client = client
	return config, nil
}

// TestConnection teste la connexion à Supabase
func (s *SupabaseConfig) TestConnection() error {
	// Test simple de la connexion en récupérant la version de l'API
	_, err := s.Client.DB.From("_test_connection").Select("*").Limit(1).Execute()
	if err != nil {
		return fmt.Errorf("erreur lors du test de connexion: %v", err)
	}
	return nil
}
