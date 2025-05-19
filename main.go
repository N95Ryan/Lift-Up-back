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
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	// Initialisation de Supabase
	supabaseConfig, err := config.InitSupabase()
	if err != nil {
		log.Fatal("Failed to initialize Supabase:", err)
	}

	// Configuration de Gin
	router := gin.Default()

	// Configuration des routes
	routes.SetupUserRoutes(router)

	// Route de test
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Démarrage du serveur
	log.Println("Server starting on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
