package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/mremperor-atwork/rpg1/api1/internal/middleware"
	"github.com/mremperor-atwork/rpg1/api1/internal/models"
	"github.com/mremperor-atwork/rpg1/api1/internal/repository"
)

type UserHandlers struct {
	userRepo *repository.UserRepository
}

func NewUserHandlers() *UserHandlers {
	return &UserHandlers{
		userRepo: repository.NewUserRepository(),
	}
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user account
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user body models.User true "User data"
// @Success 201 {object} CreateUserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/users [post]
func (h *UserHandlers) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	if err := h.userRepo.ValidateUserCreation(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.userRepo.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user":    user,
	})
}

// GetUser godoc
// @Summary Get user by ID
// @Description Retrieve a user by their unique identifier
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "User ID" format(uuid)
// @Success 200 {object} GetUserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/users/{id} [get]
func (h *UserHandlers) GetUser(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := h.userRepo.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// GetCurrentUser godoc
// @Summary Get current user
// @Description Retrieve the currently authenticated user's information
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} GetUserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/users/me [get]
func (h *UserHandlers) GetCurrentUser(c *gin.Context) {
	authUser, exists := middleware.GetUserFromContext(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userID, err := uuid.Parse(authUser.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := h.userRepo.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// UpdateUser godoc
// @Summary Update user
// @Description Update an existing user's information
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "User ID" format(uuid)
// @Param user body models.User true "Updated user data"
// @Success 200 {object} models.User
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/users/{id} [put]
func (h *UserHandlers) UpdateUser(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	user.ID = userID

	if err := h.userRepo.UpdateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// DeleteUser godoc
// @Summary Delete user
// @Description Delete a user by their ID
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "User ID" format(uuid)
// @Success 200 {object} DeleteResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/users/{id} [delete]
func (h *UserHandlers) DeleteUser(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := h.userRepo.DeleteUser(userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

// GetAllUsers godoc
// @Summary Get all users
// @Description Retrieve all users in the system
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.User
// @Failure 500 {object} ErrorResponse
// @Router /api/users [get]
func (h *UserHandlers) GetAllUsers(c *gin.Context) {
	users, err := h.userRepo.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// SearchUsers godoc
// @Summary Search users
// @Description Search users by name, email, or other criteria
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param q query string true "Search query"
// @Success 200 {object} UserSearchResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/users/search [get]
func (h *UserHandlers) SearchUsers(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Search query is required"})
		return
	}

	users, err := h.userRepo.SearchUsers(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search users: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
		"count": len(users),
		"query": query,
	})
}

// GetUserStats godoc
// @Summary Get user statistics
// @Description Retrieve statistics for a specific user
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "User ID" format(uuid)
// @Success 200 {object} UserStatsResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/users/{id}/stats [get]
func (h *UserHandlers) GetUserStats(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	stats, err := h.userRepo.GetUserStats(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"stats": stats})
}

// GetUserActivity godoc
// @Summary Get user activity
// @Description Retrieve a user's activity history
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "User ID" format(uuid)
// @Param limit query int false "Number of activities to return (default: 10)"
// @Success 200 {array} map[string]interface{}
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/users/{id}/activity [get]
func (h *UserHandlers) GetUserActivity(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	limitStr := c.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}

	activities, err := h.userRepo.GetUserActivity(userID, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user activity: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"activities": activities})
}

// GetUserPreferences godoc
// @Summary Get user preferences
// @Description Retrieve a user's preferences and settings
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "User ID" format(uuid)
// @Success 200 {object} UserPreferencesResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/users/{id}/preferences [get]
func (h *UserHandlers) GetUserPreferences(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	preferences, err := h.userRepo.GetUserPreferences(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"preferences": preferences})
}

// UpdateUserPreferences godoc
// @Summary Update user preferences
// @Description Update a user's preferences and settings
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "User ID" format(uuid)
// @Param preferences body interface{} true "User preferences"
// @Success 200 {object} UserPreferencesResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/users/{id}/preferences [put]
func (h *UserHandlers) UpdateUserPreferences(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var preferences map[string]interface{}
	if err := c.ShouldBindJSON(&preferences); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	if err := h.userRepo.UpdateUserPreferences(userID, preferences); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update preferences: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"preferences": preferences})
}

// Login godoc
// @Summary User login
// @Description Authenticate user and return session token
// @Tags auth
// @Accept json
// @Produce json
// @Param login_data body LoginRequest true "Login data"
// @Success 200 {object} LoginResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Router /api/auth/login [post]
func (h *UserHandlers) Login(c *gin.Context) {
	var loginData struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful (use Supabase Auth for actual authentication)",
		"email":   loginData.Email,
	})
}

// Logout godoc
// @Summary User logout
// @Description Invalidate session token
// @Tags auth
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} LogoutResponse
// @Failure 400 {object} ErrorResponse
// @Router /api/auth/logout [post]
func (h *UserHandlers) Logout(c *gin.Context) {
	sessionToken := c.GetHeader("Authorization")
	if sessionToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No session token provided"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logout successful (use Supabase Auth for actual logout)"})
}

// ChangePassword godoc
// @Summary Change user password
// @Description Change a user's password
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "User ID" format(uuid)
// @Param password body ChangePasswordRequest true "Password change data"
// @Success 200 {object} ChangePasswordResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/users/{id}/change-password [post]
func (h *UserHandlers) ChangePassword(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var request struct {
		CurrentPassword string `json:"current_password" binding:"required"`
		NewPassword     string `json:"new_password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	if err := h.userRepo.ChangePassword(userID, request.CurrentPassword, request.NewPassword); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to change password: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password changed successfully"})
}

// Register godoc
// @Summary User registration
// @Description Register a new user account
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.User true "User data"
// @Success 201 {object} RegisterResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/auth/register [post]
func (h *UserHandlers) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	if err := h.userRepo.ValidateUserCreation(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.userRepo.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully (use Supabase Auth for actual registration)",
		"user":    user,
	})
}

// GetUserProfile godoc
// @Summary Get user profile
// @Description Retrieve a user's public profile information
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "User ID" format(uuid)
// @Success 200 {object} UserProfileResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/users/{id}/profile [get]
func (h *UserHandlers) GetUserProfile(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := h.userRepo.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"profile": user})
}

// User Response models for Swagger documentation
type CreateUserResponse struct {
	Message string      `json:"message" example:"User created successfully"`
	User    models.User `json:"user"`
}

type GetUserResponse struct {
	User models.User `json:"user"`
}

type UserSearchResponse struct {
	Users []models.User `json:"users"`
	Count int           `json:"count" example:"5"`
	Query string        `json:"query" example:"john"`
}

type UserStatsResponse struct {
	Stats map[string]interface{} `json:"stats"`
}

type UserProfileResponse struct {
	Profile models.User `json:"profile"`
}

type UserPreferencesResponse struct {
	Preferences map[string]interface{} `json:"preferences"`
}

type ChangePasswordRequest struct {
	CurrentPassword string `json:"current_password" example:"oldpassword123" binding:"required"`
	NewPassword     string `json:"new_password" example:"newpassword123" binding:"required,min=6"`
}

type ChangePasswordResponse struct {
	Message string `json:"message" example:"Password changed successfully"`
}
