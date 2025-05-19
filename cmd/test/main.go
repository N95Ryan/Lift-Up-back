package main

import (
	"fmt"
	"log"

	"lift-up-back/config"
)

func main() {
	// Initialisation de la connexion Supabase
	supabase, err := config.InitSupabase()
	if err != nil {
		log.Fatalf("Erreur lors de l'initialisation de Supabase: %v", err)
	}

	// Test de la connexion
	if err := supabase.TestConnection(); err != nil {
		log.Fatalf("Erreur lors du test de connexion: %v", err)
	}

	fmt.Println("✅ Connexion à Supabase réussie !")
}
