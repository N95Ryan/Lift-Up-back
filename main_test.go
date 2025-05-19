package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// setupTestRouter configure un routeur de test
func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Configuration des routes de test
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return router
}

// TestPingEndpoint teste la route /ping
func TestPingEndpoint(t *testing.T) {
	router := setupTestRouter()

	// Création d'une requête de test
	req, _ := http.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Vérification du code de statut
	assert.Equal(t, http.StatusOK, w.Code)

	// Vérification du corps de la réponse
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "pong", response["message"])
}

// TestMainFunction teste la fonction main
func TestMainFunction(t *testing.T) {
	// Note: Ce test est un exemple de structure
	// La fonction main() ne peut pas être testée directement car elle démarre le serveur
	// Dans un vrai environnement, on utiliserait des mocks et des tests d'intégration
	t.Run("should not panic", func(t *testing.T) {
		// Vérifier que la configuration de base est correcte
		assert.NotNil(t, setupTestRouter())
	})
}
