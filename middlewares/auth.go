package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware vérifie le token JWT de Supabase
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// Le token doit être au format "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}

		// TODO: Vérifier le token avec Supabase
		// Pour l'instant, on accepte tous les tokens
		_ = parts[1] // token sera utilisé une fois la vérification Supabase implémentée
		c.Set("user_id", "test_user_id")
		c.Next()
	}
}
