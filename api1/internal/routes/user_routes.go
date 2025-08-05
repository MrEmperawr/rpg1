package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mremperor-atwork/rpg1/api1/internal/handlers"
)

func SetupUserRoutes(r *gin.Engine) {
	userHandlers := handlers.NewUserHandlers()

	// User API group
	userAPI := r.Group("/api/users")
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

		// User authentication
		userAPI.POST("/register", userHandlers.Register)
		userAPI.POST("/login", userHandlers.Login)
		userAPI.POST("/logout", userHandlers.Logout)
		userAPI.POST("/:id/change-password", userHandlers.ChangePassword)
	}
}
