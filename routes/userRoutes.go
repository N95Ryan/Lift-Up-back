package routes

import (
	"lift-up-back/controllers"
	"lift-up-back/middlewares"

	"github.com/gin-gonic/gin"
)

// SetupUserRoutes configure les routes liées aux utilisateurs
func SetupUserRoutes(router *gin.Engine) {
	userController := controllers.NewUserController()

	// Groupe de routes protégées par l'authentification
	authorized := router.Group("/api")
	authorized.Use(middlewares.AuthMiddleware())
	{
		// Routes du profil utilisateur
		authorized.GET("/profile", userController.GetProfile)
		authorized.PUT("/profile", userController.UpdateProfile)
		authorized.POST("/bmi", userController.CalculateBMI)
	}
}
