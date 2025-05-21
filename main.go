package main

import (
	"log"

	"lift-up-back/config"
	"lift-up-back/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Chargement des variables d'environnement
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Erreur: fichier .env non trouvé. Veuillez créer ce fichier avec vos variables d'environnement Supabase.")
	}

	// Initialisation de Supabase
	if _, err := config.InitSupabase(); err != nil {
		log.Fatal("Failed to initialize Supabase:", err)
	}

	// Configuration de Gin
	router := gin.Default()

	// Configuration des routes
	routes.SetupUserRoutes(router)

	// Démarrage du serveur
	log.Println("Server starting on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
