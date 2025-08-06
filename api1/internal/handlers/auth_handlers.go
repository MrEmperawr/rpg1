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

// Register godoc
// @Summary Register a new user
// @Description Register a new user with email, password, and optional profile information
// @Tags auth
// @Accept json
// @Produce json
// @Param user body RegisterRequest true "User registration data"
// @Success 201 {object} RegisterResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/auth/register [post]
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

// Login godoc
// @Summary Login user
// @Description Authenticate user with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body LoginRequest true "User credentials"
// @Success 200 {object} LoginResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /api/auth/login [post]
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

// Logout godoc
// @Summary Logout user
// @Description Logout the current user and invalidate their session
// @Tags auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} LogoutResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/auth/logout [post]
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

// RefreshToken godoc
// @Summary Refresh access token
// @Description Refresh the access token using a refresh token
// @Tags auth
// @Accept json
// @Produce json
// @Param refresh body RefreshTokenRequest true "Refresh token"
// @Success 200 {object} RefreshTokenResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /api/auth/refresh [post]
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

// GetCurrentUser godoc
// @Summary Get current user
// @Description Get the current authenticated user's information
// @Tags auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} GetCurrentUserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /api/auth/me [get]
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

// Request/Response models for Swagger documentation
type RegisterRequest struct {
	Email       string `json:"email" example:"user@example.com" binding:"required,email"`
	Password    string `json:"password" example:"password123" binding:"required,min=6"`
	FirstName   string `json:"first_name" example:"John"`
	LastName    string `json:"last_name" example:"Doe"`
	DisplayName string `json:"display_name" example:"johndoe"`
}

type RegisterResponse struct {
	Message string      `json:"message" example:"User registered successfully"`
	User    interface{} `json:"user"`
	Session interface{} `json:"session"`
}

type LoginRequest struct {
	Email    string `json:"email" example:"user@example.com" binding:"required,email"`
	Password string `json:"password" example:"password123" binding:"required"`
}

type LoginResponse struct {
	Message string      `json:"message" example:"Login successful"`
	User    interface{} `json:"user"`
	Session interface{} `json:"session"`
}

type LogoutResponse struct {
	Message string `json:"message" example:"Logout successful"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" example:"refresh_token_here" binding:"required"`
}

type RefreshTokenResponse struct {
	Message string      `json:"message" example:"Token refreshed successfully"`
	User    interface{} `json:"user"`
	Session interface{} `json:"session"`
}

type GetCurrentUserResponse struct {
	User interface{} `json:"user"`
}

type ErrorResponse struct {
	Error string `json:"error" example:"Error message"`
}
