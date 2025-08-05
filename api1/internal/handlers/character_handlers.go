package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/mremperor-atwork/rpg1/api1/internal/features/game"
	"github.com/mremperor-atwork/rpg1/api1/internal/models"
	"github.com/mremperor-atwork/rpg1/api1/internal/repository"
)

type CharacterHandlers struct {
	characterRepo *repository.CharacterRepository
}

func NewCharacterHandlers() *CharacterHandlers {
	return &CharacterHandlers{
		characterRepo: repository.NewCharacterRepository(),
	}
}

func (h *CharacterHandlers) CreateCharacter(c *gin.Context) {
	var character models.Character
	if err := c.ShouldBindJSON(&character); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	// For now, we'll use a placeholder - in production this would come from JWT token
	userID := uuid.MustParse("00000000-0000-0000-0000-000000000001") // Placeholder
	character.UserID = &userID

	if err := h.characterRepo.ValidateCharacterCreation(&character); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.characterRepo.CreateCharacter(&character); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create character: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, character)
}

func (h *CharacterHandlers) GetCharacter(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid character ID"})
		return
	}

	character, err := h.characterRepo.GetCharacterByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Character not found"})
		return
	}

	c.JSON(http.StatusOK, character)
}

func (h *CharacterHandlers) GetUserCharacters(c *gin.Context) {
	// TODO: Get user ID from authentication context
	userID := uuid.MustParse("00000000-0000-0000-0000-000000000001") // Placeholder

	characters, err := h.characterRepo.GetCharactersByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch characters: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, characters)
}

func (h *CharacterHandlers) GetCampaignCharacters(c *gin.Context) {
	campaignIDStr := c.Param("campaignId")
	campaignID, err := uuid.Parse(campaignIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid campaign ID"})
		return
	}

	characters, err := h.characterRepo.GetCharactersByCampaignID(campaignID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch campaign characters: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, characters)
}

func (h *CharacterHandlers) UpdateCharacter(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid character ID"})
		return
	}

	var character models.Character
	if err := c.ShouldBindJSON(&character); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	character.ID = id

	if err := h.characterRepo.UpdateCharacter(&character); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update character: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, character)
}

func (h *CharacterHandlers) DeleteCharacter(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid character ID"})
		return
	}

	if err := h.characterRepo.DeleteCharacter(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete character: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Character deleted successfully"})
}

func (h *CharacterHandlers) SearchCharacters(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Search query is required"})
		return
	}

	// TODO: Get user ID from authentication context
	userID := uuid.MustParse("00000000-0000-0000-0000-000000000001") // Placeholder

	characters, err := h.characterRepo.SearchCharacters(userID, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search characters: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, characters)
}

func (h *CharacterHandlers) GetCharacterAttributes(c *gin.Context) {
	characterIDStr := c.Param("id")
	characterID, err := uuid.Parse(characterIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid character ID"})
		return
	}

	attributes, err := h.characterRepo.GetCharacterAttributes(characterID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch character attributes: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, attributes)
}

func (h *CharacterHandlers) SetCharacterAttribute(c *gin.Context) {
	characterIDStr := c.Param("id")
	characterID, err := uuid.Parse(characterIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid character ID"})
		return
	}

	var attribute game.CharacterAttribute
	if err := c.ShouldBindJSON(&attribute); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	attribute.CharacterID = characterID

	if err := h.characterRepo.SetCharacterAttribute(&attribute); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set character attribute: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, attribute)
}

func (h *CharacterHandlers) DeleteCharacterAttribute(c *gin.Context) {
	characterIDStr := c.Param("id")
	characterID, err := uuid.Parse(characterIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid character ID"})
		return
	}

	attributeIDStr := c.Param("attributeId")
	attributeID, err := uuid.Parse(attributeIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid attribute ID"})
		return
	}

	if err := h.characterRepo.DeleteCharacterAttribute(characterID, attributeID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete character attribute: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Character attribute deleted successfully"})
}

func (h *CharacterHandlers) GetCharacterSkills(c *gin.Context) {
	characterIDStr := c.Param("id")
	characterID, err := uuid.Parse(characterIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid character ID"})
		return
	}

	skills, err := h.characterRepo.GetCharacterSkills(characterID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch character skills: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, skills)
}

func (h *CharacterHandlers) SetCharacterSkill(c *gin.Context) {
	characterIDStr := c.Param("id")
	characterID, err := uuid.Parse(characterIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid character ID"})
		return
	}

	var skill game.CharacterSkill
	if err := c.ShouldBindJSON(&skill); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	skill.CharacterID = characterID

	if err := h.characterRepo.SetCharacterSkill(&skill); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set character skill: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, skill)
}

func (h *CharacterHandlers) DeleteCharacterSkill(c *gin.Context) {
	characterIDStr := c.Param("id")
	characterID, err := uuid.Parse(characterIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid character ID"})
		return
	}

	skillIDStr := c.Param("skillId")
	skillID, err := uuid.Parse(skillIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid skill ID"})
		return
	}

	if err := h.characterRepo.DeleteCharacterSkill(characterID, skillID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete character skill: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Character skill deleted successfully"})
}

func (h *CharacterHandlers) GetCharacterSkillSpecialties(c *gin.Context) {
	characterIDStr := c.Param("id")
	characterID, err := uuid.Parse(characterIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid character ID"})
		return
	}

	specialties, err := h.characterRepo.GetCharacterSkillSpecialties(characterID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch character skill specialties: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, specialties)
}

func (h *CharacterHandlers) AddCharacterSkillSpecialty(c *gin.Context) {
	characterIDStr := c.Param("id")
	characterID, err := uuid.Parse(characterIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid character ID"})
		return
	}

	var request struct {
		SkillSpecialtyID uuid.UUID `json:"skill_specialty_id"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	if err := h.characterRepo.AddCharacterSkillSpecialty(characterID, request.SkillSpecialtyID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add character skill specialty: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Character skill specialty added successfully"})
}

func (h *CharacterHandlers) RemoveCharacterSkillSpecialty(c *gin.Context) {
	characterIDStr := c.Param("id")
	characterID, err := uuid.Parse(characterIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid character ID"})
		return
	}

	specialtyIDStr := c.Param("specialtyId")
	specialtyID, err := uuid.Parse(specialtyIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid specialty ID"})
		return
	}

	if err := h.characterRepo.RemoveCharacterSkillSpecialty(characterID, specialtyID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove character skill specialty: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Character skill specialty removed successfully"})
}

func (h *CharacterHandlers) GetCharacterQualities(c *gin.Context) {
	characterIDStr := c.Param("id")
	characterID, err := uuid.Parse(characterIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid character ID"})
		return
	}

	qualities, err := h.characterRepo.GetCharacterQualities(characterID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch character qualities: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, qualities)
}

func (h *CharacterHandlers) SetCharacterQuality(c *gin.Context) {
	characterIDStr := c.Param("id")
	characterID, err := uuid.Parse(characterIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid character ID"})
		return
	}

	var quality game.CharacterQuality
	if err := c.ShouldBindJSON(&quality); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	quality.CharacterID = characterID

	if err := h.characterRepo.SetCharacterQuality(&quality); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set character quality: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, quality)
}

func (h *CharacterHandlers) DeleteCharacterQuality(c *gin.Context) {
	characterIDStr := c.Param("id")
	characterID, err := uuid.Parse(characterIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid character ID"})
		return
	}

	qualityIDStr := c.Param("qualityId")
	qualityID, err := uuid.Parse(qualityIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quality ID"})
		return
	}

	if err := h.characterRepo.DeleteCharacterQuality(characterID, qualityID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete character quality: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Character quality deleted successfully"})
}

func (h *CharacterHandlers) GetCharacterEquipment(c *gin.Context) {
	characterIDStr := c.Param("id")
	characterID, err := uuid.Parse(characterIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid character ID"})
		return
	}

	equipment, err := h.characterRepo.GetCharacterEquipment(characterID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch character equipment: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, equipment)
}

func (h *CharacterHandlers) AddCharacterEquipment(c *gin.Context) {
	characterIDStr := c.Param("id")
	characterID, err := uuid.Parse(characterIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid character ID"})
		return
	}

	var equipment game.CharacterEquipment
	if err := c.ShouldBindJSON(&equipment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	equipment.CharacterID = characterID

	if err := h.characterRepo.AddCharacterEquipment(&equipment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add character equipment: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, equipment)
}

func (h *CharacterHandlers) UpdateCharacterEquipmentQuantity(c *gin.Context) {
	characterIDStr := c.Param("id")
	characterID, err := uuid.Parse(characterIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid character ID"})
		return
	}

	equipmentIDStr := c.Param("equipmentId")
	equipmentID, err := uuid.Parse(equipmentIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid equipment ID"})
		return
	}

	quantityStr := c.Param("quantity")
	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quantity"})
		return
	}

	if err := h.characterRepo.UpdateCharacterEquipmentQuantity(characterID, equipmentID, quantity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update character equipment quantity: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Character equipment quantity updated successfully"})
}

func (h *CharacterHandlers) RemoveCharacterEquipment(c *gin.Context) {
	characterIDStr := c.Param("id")
	characterID, err := uuid.Parse(characterIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid character ID"})
		return
	}

	equipmentIDStr := c.Param("equipmentId")
	equipmentID, err := uuid.Parse(equipmentIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid equipment ID"})
		return
	}

	if err := h.characterRepo.RemoveCharacterEquipment(characterID, equipmentID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove character equipment: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Character equipment removed successfully"})
}

func (h *CharacterHandlers) GetCharacterPersonalEquipment(c *gin.Context) {
	characterIDStr := c.Param("id")
	characterID, err := uuid.Parse(characterIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid character ID"})
		return
	}

	equipment, err := h.characterRepo.GetCharacterPersonalEquipment(characterID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch character personal equipment: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, equipment)
}

func (h *CharacterHandlers) AddCharacterPersonalEquipment(c *gin.Context) {
	characterIDStr := c.Param("id")
	characterID, err := uuid.Parse(characterIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid character ID"})
		return
	}

	var equipment game.CharacterPersonalEquipment
	if err := c.ShouldBindJSON(&equipment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	equipment.CharacterID = characterID

	if err := h.characterRepo.AddCharacterPersonalEquipment(&equipment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add character personal equipment: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, equipment)
}

func (h *CharacterHandlers) UpdateCharacterPersonalEquipmentQuantity(c *gin.Context) {
	characterIDStr := c.Param("id")
	characterID, err := uuid.Parse(characterIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid character ID"})
		return
	}

	equipmentIDStr := c.Param("equipmentId")
	equipmentID, err := uuid.Parse(equipmentIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid equipment ID"})
		return
	}

	quantityStr := c.Param("quantity")
	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quantity"})
		return
	}

	if err := h.characterRepo.UpdateCharacterPersonalEquipmentQuantity(characterID, equipmentID, quantity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update character personal equipment quantity: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Character personal equipment quantity updated successfully"})
}

func (h *CharacterHandlers) RemoveCharacterPersonalEquipment(c *gin.Context) {
	characterIDStr := c.Param("id")
	characterID, err := uuid.Parse(characterIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid character ID"})
		return
	}

	equipmentIDStr := c.Param("equipmentId")
	equipmentID, err := uuid.Parse(equipmentIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid equipment ID"})
		return
	}

	if err := h.characterRepo.RemoveCharacterPersonalEquipment(characterID, equipmentID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove character personal equipment: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Character personal equipment removed successfully"})
}

func (h *CharacterHandlers) GetCharacterDerivedStats(c *gin.Context) {
	characterIDStr := c.Param("id")
	characterID, err := uuid.Parse(characterIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid character ID"})
		return
	}

	stats, err := h.characterRepo.GetCharacterDerivedStats(characterID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch character derived stats: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, stats)
}

func (h *CharacterHandlers) SetCharacterDerivedStats(c *gin.Context) {
	characterIDStr := c.Param("id")
	characterID, err := uuid.Parse(characterIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid character ID"})
		return
	}

	var stats game.CharacterDerivedStats
	if err := c.ShouldBindJSON(&stats); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	stats.CharacterID = characterID

	if err := h.characterRepo.SetCharacterDerivedStats(&stats); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set character derived stats: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, stats)
}
