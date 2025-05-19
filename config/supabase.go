package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	supa "github.com/supabase-community/supabase-go"
)

type SupabaseConfig struct {
	URL       string
	APIKey    string
	JWTSecret string
	Client    *supa.Client
}

// NewSupabaseConfig crée une nouvelle instance de configuration
func NewSupabaseConfig() *SupabaseConfig {
	if err := godotenv.Load(".env.local"); err != nil {
		fmt.Printf("Erreur lors du chargement du fichier .env.local: %v\n", err)
	}

	return &SupabaseConfig{
		URL:    os.Getenv("EXPO_PUBLIC_SUPABASE_URL"),
		APIKey: os.Getenv("EXPO_PUBLIC_SUPABASE_ANON_KEY"),
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
	if s.Client == nil {
		return fmt.Errorf("client Supabase non initialisé")
	}
	return nil
}
