package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mremperor-atwork/rpg1/api1/internal/handlers"
	"github.com/mremperor-atwork/rpg1/api1/internal/middleware"
)

func SetupUserRoutes(r *gin.Engine) {
	userHandlers := handlers.NewUserHandlers()

	// User API group - all user operations require authentication
	userAPI := r.Group("/api/users")
	userAPI.Use(middleware.AuthMiddleware())
	{
		// User CRUD operations
		userAPI.POST("", userHandlers.CreateUser)
		userAPI.GET("", userHandlers.GetAllUsers)
		userAPI.GET("/search", userHandlers.SearchUsers)
		userAPI.GET("/me", userHandlers.GetCurrentUser)
		userAPI.GET("/:id", userHandlers.GetUser)
		userAPI.PUT("/:id", userHandlers.UpdateUser)
		userAPI.DELETE("/:id", userHandlers.DeleteUser)

		// User profile and statistics
		userAPI.GET("/:id/profile", userHandlers.GetUserProfile)
		userAPI.GET("/:id/stats", userHandlers.GetUserStats)
		userAPI.GET("/:id/activity", userHandlers.GetUserActivity)

		// User preferences
		userAPI.GET("/:id/preferences", userHandlers.GetUserPreferences)
		userAPI.PUT("/:id/preferences", userHandlers.UpdateUserPreferences)

		// User authentication (password change)
		userAPI.POST("/:id/change-password", userHandlers.ChangePassword)
	}
}
