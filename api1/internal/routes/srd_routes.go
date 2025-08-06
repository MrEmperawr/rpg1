package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mremperor-atwork/rpg1/api1/internal/handlers"
	"github.com/mremperor-atwork/rpg1/api1/internal/middleware"
)

func SetupSRDRoutes(router *gin.Engine) {
	srdHandlers := handlers.NewSRDHandlers()

	srd := router.Group("/api/srd")
	{
		// Read-only endpoints (public)
		srd.GET("/entries", srdHandlers.GetSRDEntries)
		srd.GET("/entries/:id", srdHandlers.GetSRDEntryByID)
		srd.GET("/entries/category/:category", srdHandlers.GetSRDEntriesByCategory)
		srd.GET("/categories", srdHandlers.GetSRDCategories)
		srd.GET("/search", srdHandlers.SearchSRDEntries)
		srd.GET("/content", srdHandlers.GetSRDContent)
		srd.GET("/content/:title", srdHandlers.GetSRDContentByTitle)

		// Protected CRUD endpoints (require authentication)
		protected := srd.Group("")
		protected.Use(middleware.RequireAuth())
		{
			// SRD Entry CRUD
			protected.POST("/entries", srdHandlers.CreateSRDEntry)
			protected.PUT("/entries/:id", srdHandlers.UpdateSRDEntry)
			protected.DELETE("/entries/:id", srdHandlers.DeleteSRDEntry)

			// SRD Content CRUD
			protected.POST("/content", srdHandlers.CreateSRDContent)
			protected.PUT("/content/:title", srdHandlers.UpdateSRDContent)
			protected.DELETE("/content/:title", srdHandlers.DeleteSRDContent)
		}
	}
}
