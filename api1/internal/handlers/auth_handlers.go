package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mremperor-atwork/rpg1/api1/internal/supabase"
	"github.com/supabase-community/gotrue-go/types"
	supa "github.com/supabase-community/supabase-go"
)

type AuthHandlers struct {
	supabaseClient *supa.Client
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	userMetadata := map[string]interface{}{
		"first_name":   registerData.FirstName,
		"last_name":    registerData.LastName,
		"display_name": registerData.DisplayName,
	}

	authResponse, err := h.supabaseClient.Auth.Signup(types.SignupRequest{
		Email:    registerData.Email,
		Password: registerData.Password,
		Data:     userMetadata,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user":    authResponse.User,
		"session": authResponse.Session,
	})
}

func (h *AuthHandlers) Login(c *gin.Context) {
	var loginData struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	authResponse, err := h.supabaseClient.Auth.SignInWithEmailPassword(loginData.Email, loginData.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user":    authResponse.User,
		"session": authResponse,
	})
}

func (h *AuthHandlers) Logout(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization header required"})
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	authClient := h.supabaseClient.Auth.WithToken(token)

	err := authClient.Logout()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to logout: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
}

func (h *AuthHandlers) RefreshToken(c *gin.Context) {
	var refreshData struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}

	if err := c.ShouldBindJSON(&refreshData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	authResponse, err := h.supabaseClient.Auth.RefreshToken(refreshData.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Token refreshed successfully",
		"user":    authResponse.User,
		"session": authResponse,
	})
}

func (h *AuthHandlers) GetCurrentUser(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization header required"})
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	authClient := h.supabaseClient.Auth.WithToken(token)

	user, err := authClient.GetUser()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
