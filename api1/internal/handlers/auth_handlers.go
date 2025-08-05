package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mremperor-atwork/rpg1/api1/internal/supabase"
)

type AuthHandlers struct {
	supabaseClient interface{}
}

func NewAuthHandlers() *AuthHandlers {
	return &AuthHandlers{
		supabaseClient: supabase.GetClient(),
	}
}

func (h *AuthHandlers) Register(c *gin.Context) {
	var registerData struct {
		Email       string `json:"email" binding:"required,email"`
		Password    string `json:"password" binding:"required,min=6"`
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		DisplayName string `json:"display_name"`
	}

	if err := c.ShouldBindJSON(&registerData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body: " + err.Error()})
		return
	}

	// For now, return a mock response
	// TODO: Implement actual Supabase Auth registration
	c.JSON(http.StatusCreated, gin.H{
		"message": "user registered successfully (mock response)",
		"user": gin.H{
			"id":    "mock-user-id",
			"email": registerData.Email,
		},
		"access_token":  "mock-access-token",
		"refresh_token": "mock-refresh-token",
	})
}

func (h *AuthHandlers) Login(c *gin.Context) {
	var loginData struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body: " + err.Error()})
		return
	}

	// For now, return a mock response
	// TODO: Implement actual Supabase Auth login
	c.JSON(http.StatusOK, gin.H{
		"message": "login successful (mock response)",
		"session": gin.H{
			"access_token":  "mock-access-token",
			"refresh_token": "mock-refresh-token",
			"user": gin.H{
				"id":    "mock-user-id",
				"email": loginData.Email,
			},
		},
	})
}

func (h *AuthHandlers) Logout(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no authorization token provided"})
		return
	}

	// Extract token from "Bearer <token>"
	token := strings.TrimPrefix(authHeader, "Bearer ")
	if token == authHeader {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid authorization header format"})
		return
	}

	// For now, return a mock response
	// TODO: Implement actual Supabase Auth logout
	c.JSON(http.StatusOK, gin.H{"message": "logout successful (mock response)"})
}

// RefreshToken handles token refresh
func (h *AuthHandlers) RefreshToken(c *gin.Context) {
	var refreshData struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}

	if err := c.ShouldBindJSON(&refreshData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body: " + err.Error()})
		return
	}

	// For now, return a mock response
	// TODO: Implement actual Supabase Auth token refresh
	c.JSON(http.StatusOK, gin.H{
		"message": "token refreshed successfully (mock response)",
		"session": gin.H{
			"access_token":  "new-mock-access-token",
			"refresh_token": "new-mock-refresh-token",
			"user": gin.H{
				"id":    "mock-user-id",
				"email": "user@example.com",
			},
		},
	})
}

func (h *AuthHandlers) GetCurrentUser(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no authorization token provided"})
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	if token == authHeader {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid authorization header format"})
		return
	}

	// TODO: Implement actual Supabase Auth user retrieval
	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"id":       "mock-user-id",
			"email":    "user@example.com",
			"metadata": map[string]interface{}{},
		},
	})
}
