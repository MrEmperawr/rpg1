package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mremperor-atwork/rpg1/api1/internal/handlers"
)

func SetupCharacterRoutes(r *gin.Engine) {
	characterHandlers := handlers.NewCharacterHandlers()

	// Character API group - main character operations
	characterAPI := r.Group("/api/characters")
	{
		// Character CRUD operations
		characterAPI.POST("", characterHandlers.CreateCharacter)
		characterAPI.GET("", characterHandlers.GetUserCharacters)
		characterAPI.GET("/search", characterHandlers.SearchCharacters)
		characterAPI.GET("/:id", characterHandlers.GetCharacter)
		characterAPI.PUT("/:id", characterHandlers.UpdateCharacter)
		characterAPI.DELETE("/:id", characterHandlers.DeleteCharacter)

		// Campaign characters
		characterAPI.GET("/campaign/:campaignId", characterHandlers.GetCampaignCharacters)
	}

	// Character details API group (separate to avoid route conflicts)
	characterDetailsAPI := r.Group("/api/characters/:id")
	{
		// Character attributes
		characterDetailsAPI.GET("/attributes", characterHandlers.GetCharacterAttributes)
		characterDetailsAPI.POST("/attributes", characterHandlers.SetCharacterAttribute)
		characterDetailsAPI.DELETE("/attributes/:attributeId", characterHandlers.DeleteCharacterAttribute)

		// Character skills
		characterDetailsAPI.GET("/skills", characterHandlers.GetCharacterSkills)
		characterDetailsAPI.POST("/skills", characterHandlers.SetCharacterSkill)
		characterDetailsAPI.DELETE("/skills/:skillId", characterHandlers.DeleteCharacterSkill)

		// Character skill specialties
		characterDetailsAPI.GET("/skill-specialties", characterHandlers.GetCharacterSkillSpecialties)
		characterDetailsAPI.POST("/skill-specialties", characterHandlers.AddCharacterSkillSpecialty)
		characterDetailsAPI.DELETE("/skill-specialties/:specialtyId", characterHandlers.RemoveCharacterSkillSpecialty)

		// Character qualities
		characterDetailsAPI.GET("/qualities", characterHandlers.GetCharacterQualities)
		characterDetailsAPI.POST("/qualities", characterHandlers.SetCharacterQuality)
		characterDetailsAPI.DELETE("/qualities/:qualityId", characterHandlers.DeleteCharacterQuality)

		// Character equipment
		characterDetailsAPI.GET("/equipment", characterHandlers.GetCharacterEquipment)
		characterDetailsAPI.POST("/equipment", characterHandlers.AddCharacterEquipment)
		characterDetailsAPI.PUT("/equipment/:equipmentId/quantity/:quantity", characterHandlers.UpdateCharacterEquipmentQuantity)
		characterDetailsAPI.DELETE("/equipment/:equipmentId", characterHandlers.RemoveCharacterEquipment)

		// Character personal equipment
		characterDetailsAPI.GET("/personal-equipment", characterHandlers.GetCharacterPersonalEquipment)
		characterDetailsAPI.POST("/personal-equipment", characterHandlers.AddCharacterPersonalEquipment)
		characterDetailsAPI.PUT("/personal-equipment/:equipmentId/quantity/:quantity", characterHandlers.UpdateCharacterPersonalEquipmentQuantity)
		characterDetailsAPI.DELETE("/personal-equipment/:equipmentId", characterHandlers.RemoveCharacterPersonalEquipment)

		// Character derived stats
		characterDetailsAPI.GET("/derived-stats", characterHandlers.GetCharacterDerivedStats)
		characterDetailsAPI.POST("/derived-stats", characterHandlers.SetCharacterDerivedStats)
	}
}
