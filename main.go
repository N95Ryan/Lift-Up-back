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
	supabase, err := config.InitSupabase()
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

	// Route de test Supabase
	router.GET("/test-supabase", func(c *gin.Context) {
		if err := supabase.TestConnection(); err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "Connexion à Supabase réussie !",
			"config": gin.H{
				"url": supabase.URL,
			},
		})
	})

	// Démarrage du serveur
	log.Println("Server starting on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
