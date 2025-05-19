package controllers

import (
	"net/http"

	"lift-up-back/models"

	"github.com/gin-gonic/gin"
)

// UserController gère les opérations liées aux utilisateurs
type UserController struct{}

// NewUserController crée une nouvelle instance de UserController
func NewUserController() *UserController {
	return &UserController{}
}

// GetProfile récupère le profil de l'utilisateur connecté
func (uc *UserController) GetProfile(c *gin.Context) {
	userID := c.GetString("user_id")
	// TODO: Récupérer les données de l'utilisateur depuis Supabase
	user := models.User{
		ID:       userID,
		Username: "Test User",
		Email:    "test@example.com",
	}
	c.JSON(http.StatusOK, user)
}

// UpdateProfile met à jour le profil de l'utilisateur
func (uc *UserController) UpdateProfile(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: Mettre à jour les données de l'utilisateur dans Supabase
	// userID sera utilisé une fois la connexion à Supabase établie
	_ = c.GetString("user_id")

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}

// CalculateBMI calcule l'IMC de l'utilisateur
func (uc *UserController) CalculateBMI(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Calcul de l'IMC : poids (kg) / (taille (m))²
	heightInMeters := user.Height / 100
	bmi := user.Weight / (heightInMeters * heightInMeters)

	c.JSON(http.StatusOK, gin.H{
		"bmi":      bmi,
		"category": getBMICategory(bmi),
	})
}

// getBMICategory retourne la catégorie d'IMC
func getBMICategory(bmi float64) string {
	switch {
	case bmi < 18.5:
		return "Insuffisance pondérale"
	case bmi < 25:
		return "Corpulence normale"
	case bmi < 30:
		return "Surpoids"
	default:
		return "Obésité"
	}
}
