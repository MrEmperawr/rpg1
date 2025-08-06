package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mremperor-atwork/rpg1/api1/internal/handlers"
	"github.com/mremperor-atwork/rpg1/api1/internal/middleware"
)

func SetupAuthRoutes(r *gin.Engine) {
	authHandlers := handlers.NewAuthHandlers()

	authAPI := r.Group("/api/auth")
	{
		authAPI.POST("/register", authHandlers.Register)
		authAPI.POST("/login", authHandlers.Login)
		authAPI.POST("/refresh", authHandlers.RefreshToken)

		// Standard protected routes (manual refresh required)
		protected := authAPI.Group("")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.POST("/logout", authHandlers.Logout)
			protected.GET("/me", authHandlers.GetCurrentUser)
		}

		// Auto-refresh protected routes (automatic token refresh)
		autoRefreshProtected := authAPI.Group("/auto-refresh")
		autoRefreshProtected.Use(middleware.RequireAuthWithAutoRefresh())
		{
			autoRefreshProtected.GET("/me", authHandlers.GetCurrentUser)
		}
	}
}
