package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mremperor-atwork/rpg1/api1/internal/handlers"
	"github.com/mremperor-atwork/rpg1/api1/internal/middleware"
)

func SetupCharacterRoutes(r *gin.Engine) {
	characterHandlers := handlers.NewCharacterHandlers()

	characterAPI := r.Group("/api/characters")
	characterAPI.Use(middleware.AuthMiddleware())
	{
		characterAPI.POST("", characterHandlers.CreateCharacter)
		characterAPI.GET("", characterHandlers.GetUserCharacters)
		characterAPI.GET("/search", characterHandlers.SearchCharacters)
		characterAPI.GET("/:id", characterHandlers.GetCharacter)
		characterAPI.PUT("/:id", characterHandlers.UpdateCharacter)
		characterAPI.DELETE("/:id", characterHandlers.DeleteCharacter)

		characterAPI.GET("/campaign/:campaignId", characterHandlers.GetCampaignCharacters)

		characterAPI.GET("/:id/attributes", characterHandlers.GetCharacterAttributes)
		characterAPI.POST("/:id/attributes", characterHandlers.SetCharacterAttribute)
		characterAPI.DELETE("/:id/attributes/:attributeId", characterHandlers.DeleteCharacterAttribute)

		characterAPI.GET("/:id/skills", characterHandlers.GetCharacterSkills)
		characterAPI.POST("/:id/skills", characterHandlers.SetCharacterSkill)
		characterAPI.DELETE("/:id/skills/:skillId", characterHandlers.DeleteCharacterSkill)

		characterAPI.GET("/:id/skill-specialties", characterHandlers.GetCharacterSkillSpecialties)
		characterAPI.POST("/:id/skill-specialties", characterHandlers.AddCharacterSkillSpecialty)
		characterAPI.DELETE("/:id/skill-specialties/:specialtyId", characterHandlers.RemoveCharacterSkillSpecialty)

		characterAPI.GET("/:id/qualities", characterHandlers.GetCharacterQualities)
		characterAPI.POST("/:id/qualities", characterHandlers.SetCharacterQuality)
		characterAPI.DELETE("/:id/qualities/:qualityId", characterHandlers.DeleteCharacterQuality)

		characterAPI.GET("/:id/equipment", characterHandlers.GetCharacterEquipment)
		characterAPI.POST("/:id/equipment", characterHandlers.AddCharacterEquipment)
		characterAPI.PUT("/:id/equipment/:equipmentId/quantity/:quantity", characterHandlers.UpdateCharacterEquipmentQuantity)
		characterAPI.DELETE("/:id/equipment/:equipmentId", characterHandlers.RemoveCharacterEquipment)

		characterAPI.GET("/:id/personal-equipment", characterHandlers.GetCharacterPersonalEquipment)
		characterAPI.POST("/:id/personal-equipment", characterHandlers.AddCharacterPersonalEquipment)
		characterAPI.PUT("/:id/personal-equipment/:equipmentId/quantity/:quantity", characterHandlers.UpdateCharacterPersonalEquipmentQuantity)
		characterAPI.DELETE("/:id/personal-equipment/:equipmentId", characterHandlers.RemoveCharacterPersonalEquipment)

		characterAPI.GET("/:id/derived-stats", characterHandlers.GetCharacterDerivedStats)
		characterAPI.POST("/:id/derived-stats", characterHandlers.SetCharacterDerivedStats)
	}
}
