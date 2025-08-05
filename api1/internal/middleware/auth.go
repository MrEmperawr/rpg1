package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthUser struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Bearer token required"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		user, err := validateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: " + err.Error()})
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}

func validateToken(tokenString string) (*AuthUser, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	sub, ok := claims["sub"].(string)
	if !ok {
		return nil, fmt.Errorf("missing user ID in token")
	}

	email, ok := claims["email"].(string)
	if !ok {
		return nil, fmt.Errorf("missing email in token")
	}

	return &AuthUser{
		ID:    sub,
		Email: email,
	}, nil
}

func GetUserFromContext(c *gin.Context) (*AuthUser, bool) {
	user, exists := c.Get("user")
	if !exists {
		return nil, false
	}

	authUser, ok := user.(*AuthUser)
	return authUser, ok
}

func RequireAuth() gin.HandlerFunc {
	return AuthMiddleware()
}

func OptionalAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			if user, err := validateToken(tokenString); err == nil {
				c.Set("user", user)
			}
		}
		c.Next()
	}
}
