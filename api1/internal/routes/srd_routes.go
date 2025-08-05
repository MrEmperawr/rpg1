package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mremperor-atwork/rpg1/api1/internal/handlers"
)

func SetupSRDRoutes(router *gin.Engine) {
	srdHandlers := handlers.NewSRDHandlers()

	srd := router.Group("/api/srd")
	{
		srd.GET("/entries", srdHandlers.GetSRDEntries)
		srd.GET("/entries/:id", srdHandlers.GetSRDEntryByID)
		srd.GET("/entries/category/:category", srdHandlers.GetSRDEntriesByCategory)

		srd.GET("/categories", srdHandlers.GetSRDCategories)

		srd.GET("/search", srdHandlers.SearchSRDEntries)

		srd.GET("/content", srdHandlers.GetSRDContent)
		srd.GET("/content/:title", srdHandlers.GetSRDContentByTitle)
	}
}
