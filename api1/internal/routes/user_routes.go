package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mremperor-atwork/rpg1/api1/internal/handlers"
	"github.com/mremperor-atwork/rpg1/api1/internal/middleware"
)

func SetupUserRoutes(r *gin.Engine) {
	userHandlers := handlers.NewUserHandlers()

	// User API group
	userAPI := r.Group("/api/users")
	{
		// Public routes (no authentication required)
		userAPI.POST("/register", userHandlers.Register)
		userAPI.POST("/login", userHandlers.Login)

		// Protected routes (authentication required)
		protected := userAPI.Group("")
		protected.Use(middleware.AuthMiddleware())
		{
			// User CRUD operations
			protected.POST("", userHandlers.CreateUser)
			protected.GET("", userHandlers.GetAllUsers)
			protected.GET("/search", userHandlers.SearchUsers)
			protected.GET("/me", userHandlers.GetCurrentUser)
			protected.GET("/:id", userHandlers.GetUser)
			protected.PUT("/:id", userHandlers.UpdateUser)
			protected.DELETE("/:id", userHandlers.DeleteUser)

			// User profile and statistics
			protected.GET("/:id/profile", userHandlers.GetUserProfile)
			protected.GET("/:id/stats", userHandlers.GetUserStats)
			protected.GET("/:id/activity", userHandlers.GetUserActivity)

			// User preferences
			protected.GET("/:id/preferences", userHandlers.GetUserPreferences)
			protected.PUT("/:id/preferences", userHandlers.UpdateUserPreferences)

			// User authentication
			protected.POST("/logout", userHandlers.Logout)
			protected.POST("/:id/change-password", userHandlers.ChangePassword)
		}
	}
}
