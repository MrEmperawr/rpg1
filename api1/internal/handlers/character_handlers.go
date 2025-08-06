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

// CreateCharacter godoc
// @Summary Create a new character
// @Description Create a new character for the authenticated user
// @Tags characters
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param character body models.Character true "Character data"
// @Success 201 {object} models.Character
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/characters [post]
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

// GetCharacter godoc
// @Summary Get character by ID
// @Description Retrieve a character by its unique identifier
// @Tags characters
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Character ID" format(uuid)
// @Success 200 {object} models.Character
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/characters/{id} [get]
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

// GetUserCharacters godoc
// @Summary Get user's characters
// @Description Retrieve all characters belonging to the authenticated user
// @Tags characters
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Character
// @Failure 500 {object} ErrorResponse
// @Router /api/characters [get]
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

// GetCampaignCharacters godoc
// @Summary Get campaign characters
// @Description Retrieve all characters in a specific campaign
// @Tags characters
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param campaignId path string true "Campaign ID" format(uuid)
// @Success 200 {array} models.Character
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/characters/campaign/{campaignId} [get]
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

// UpdateCharacter godoc
// @Summary Update character
// @Description Update an existing character's information
// @Tags characters
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Character ID" format(uuid)
// @Param character body models.Character true "Updated character data"
// @Success 200 {object} models.Character
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/characters/{id} [put]
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

// DeleteCharacter godoc
// @Summary Delete character
// @Description Delete a character by its ID
// @Tags characters
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Character ID" format(uuid)
// @Success 200 {object} DeleteResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/characters/{id} [delete]
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

// SearchCharacters godoc
// @Summary Search characters
// @Description Search characters by name or other criteria
// @Tags characters
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param q query string true "Search query"
// @Success 200 {object} CharacterSearchResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/characters/search [get]
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

	c.JSON(http.StatusOK, gin.H{
		"characters": characters,
		"count":      len(characters),
		"query":      query,
	})
}

// GetCharacterAttributes godoc
// @Summary Get character attributes
// @Description Retrieve all attributes for a specific character
// @Tags characters
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Character ID" format(uuid)
// @Success 200 {array} game.CharacterAttribute
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/characters/{id}/attributes [get]
func (h *CharacterHandlers) GetCharacterAttributes(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid character ID"})
		return
	}

	attributes, err := h.characterRepo.GetCharacterAttributes(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Character not found"})
		return
	}

	c.JSON(http.StatusOK, attributes)
}

// SetCharacterAttribute godoc
// @Summary Set character attribute
// @Description Set or update a character's attribute value
// @Tags characters
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Character ID" format(uuid)
// @Param attribute body game.CharacterAttribute true "Attribute data"
// @Success 200 {object} game.CharacterAttribute
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/characters/{id}/attributes [post]
func (h *CharacterHandlers) SetCharacterAttribute(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid character ID"})
		return
	}

	var attribute game.CharacterAttribute
	if err := c.ShouldBindJSON(&attribute); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	attribute.CharacterID = id

	if err := h.characterRepo.SetCharacterAttribute(&attribute); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set attribute: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, attribute)
}

// GetCharacterSkills godoc
// @Summary Get character skills
// @Description Retrieve all skills for a specific character
// @Tags characters
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Character ID" format(uuid)
// @Success 200 {array} game.CharacterSkill
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/characters/{id}/skills [get]
func (h *CharacterHandlers) GetCharacterSkills(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid character ID"})
		return
	}

	skills, err := h.characterRepo.GetCharacterSkills(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Character not found"})
		return
	}

	c.JSON(http.StatusOK, skills)
}

// SetCharacterSkill godoc
// @Summary Set character skill
// @Description Set or update a character's skill value
// @Tags characters
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Character ID" format(uuid)
// @Param skill body game.CharacterSkill true "Skill data"
// @Success 200 {object} game.CharacterSkill
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/characters/{id}/skills [post]
func (h *CharacterHandlers) SetCharacterSkill(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid character ID"})
		return
	}

	var skill game.CharacterSkill
	if err := c.ShouldBindJSON(&skill); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	skill.CharacterID = id

	if err := h.characterRepo.SetCharacterSkill(&skill); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set skill: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, skill)
}

// GetCharacterEquipment godoc
// @Summary Get character equipment
// @Description Retrieve all equipment for a specific character
// @Tags characters
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Character ID" format(uuid)
// @Success 200 {array} game.CharacterEquipment
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/characters/{id}/equipment [get]
func (h *CharacterHandlers) GetCharacterEquipment(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid character ID"})
		return
	}

	equipment, err := h.characterRepo.GetCharacterEquipment(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Character not found"})
		return
	}

	c.JSON(http.StatusOK, equipment)
}

// AddCharacterEquipment godoc
// @Summary Add character equipment
// @Description Add equipment to a character's inventory
// @Tags characters
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Character ID" format(uuid)
// @Param equipment body game.CharacterEquipment true "Equipment data"
// @Success 200 {object} game.CharacterEquipment
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/characters/{id}/equipment [post]
func (h *CharacterHandlers) AddCharacterEquipment(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid character ID"})
		return
	}

	var equipment game.CharacterEquipment
	if err := c.ShouldBindJSON(&equipment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	equipment.CharacterID = id

	if err := h.characterRepo.AddCharacterEquipment(&equipment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add equipment: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, equipment)
}

// DeleteCharacterAttribute godoc
// @Summary Delete character attribute
// @Description Delete a character's attribute
// @Tags characters
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Character ID" format(uuid)
// @Param attributeId path string true "Attribute ID" format(uuid)
// @Success 200 {object} DeleteResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/characters/{id}/attributes/{attributeId} [delete]
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

// DeleteCharacterSkill godoc
// @Summary Delete character skill
// @Description Delete a character's skill
// @Tags characters
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Character ID" format(uuid)
// @Param skillId path string true "Skill ID" format(uuid)
// @Success 200 {object} DeleteResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/characters/{id}/skills/{skillId} [delete]
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

// GetCharacterSkillSpecialties godoc
// @Summary Get character skill specialties
// @Description Retrieve all skill specialties for a specific character
// @Tags characters
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Character ID" format(uuid)
// @Success 200 {array} game.CharacterSkillSpecialty
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/characters/{id}/skill-specialties [get]
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

// AddCharacterSkillSpecialty godoc
// @Summary Add character skill specialty
// @Description Add a skill specialty to a character
// @Tags characters
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Character ID" format(uuid)
// @Param specialty body AddSkillSpecialtyRequest true "Skill specialty data"
// @Success 200 {object} AddSkillSpecialtyResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/characters/{id}/skill-specialties [post]
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

// RemoveCharacterSkillSpecialty godoc
// @Summary Remove character skill specialty
// @Description Remove a skill specialty from a character
// @Tags characters
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Character ID" format(uuid)
// @Param specialtyId path string true "Specialty ID" format(uuid)
// @Success 200 {object} DeleteResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/characters/{id}/skill-specialties/{specialtyId} [delete]
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

// GetCharacterQualities godoc
// @Summary Get character qualities
// @Description Retrieve all qualities for a specific character
// @Tags characters
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Character ID" format(uuid)
// @Success 200 {array} game.CharacterQuality
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/characters/{id}/qualities [get]
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

// SetCharacterQuality godoc
// @Summary Set character quality
// @Description Set or update a character's quality
// @Tags characters
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Character ID" format(uuid)
// @Param quality body game.CharacterQuality true "Quality data"
// @Success 200 {object} game.CharacterQuality
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/characters/{id}/qualities [post]
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

// DeleteCharacterQuality godoc
// @Summary Delete character quality
// @Description Delete a character's quality
// @Tags characters
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Character ID" format(uuid)
// @Param qualityId path string true "Quality ID" format(uuid)
// @Success 200 {object} DeleteResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/characters/{id}/qualities/{qualityId} [delete]
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

// UpdateCharacterEquipmentQuantity godoc
// @Summary Update character equipment quantity
// @Description Update the quantity of equipment for a character
// @Tags characters
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Character ID" format(uuid)
// @Param equipmentId path string true "Equipment ID" format(uuid)
// @Param quantity path int true "New quantity"
// @Success 200 {object} UpdateQuantityResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/characters/{id}/equipment/{equipmentId}/quantity/{quantity} [put]
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

// RemoveCharacterEquipment godoc
// @Summary Remove character equipment
// @Description Remove equipment from a character's inventory
// @Tags characters
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Character ID" format(uuid)
// @Param equipmentId path string true "Equipment ID" format(uuid)
// @Success 200 {object} DeleteResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/characters/{id}/equipment/{equipmentId} [delete]
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

// GetCharacterPersonalEquipment godoc
// @Summary Get character personal equipment
// @Description Retrieve all personal equipment for a specific character
// @Tags characters
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Character ID" format(uuid)
// @Success 200 {array} game.CharacterPersonalEquipment
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/characters/{id}/personal-equipment [get]
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

// AddCharacterPersonalEquipment godoc
// @Summary Add character personal equipment
// @Description Add personal equipment to a character's inventory
// @Tags characters
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Character ID" format(uuid)
// @Param equipment body game.CharacterPersonalEquipment true "Personal equipment data"
// @Success 200 {object} game.CharacterPersonalEquipment
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/characters/{id}/personal-equipment [post]
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

// UpdateCharacterPersonalEquipmentQuantity godoc
// @Summary Update character personal equipment quantity
// @Description Update the quantity of personal equipment for a character
// @Tags characters
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Character ID" format(uuid)
// @Param equipmentId path string true "Equipment ID" format(uuid)
// @Param quantity path int true "New quantity"
// @Success 200 {object} UpdateQuantityResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/characters/{id}/personal-equipment/{equipmentId}/quantity/{quantity} [put]
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

// RemoveCharacterPersonalEquipment godoc
// @Summary Remove character personal equipment
// @Description Remove personal equipment from a character's inventory
// @Tags characters
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Character ID" format(uuid)
// @Param equipmentId path string true "Equipment ID" format(uuid)
// @Success 200 {object} DeleteResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/characters/{id}/personal-equipment/{equipmentId} [delete]
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

// GetCharacterDerivedStats godoc
// @Summary Get character derived stats
// @Description Retrieve derived statistics for a specific character
// @Tags characters
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Character ID" format(uuid)
// @Success 200 {object} game.CharacterDerivedStats
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/characters/{id}/derived-stats [get]
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

// SetCharacterDerivedStats godoc
// @Summary Set character derived stats
// @Description Set or update a character's derived statistics
// @Tags characters
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Character ID" format(uuid)
// @Param stats body game.CharacterDerivedStats true "Derived stats data"
// @Success 200 {object} game.CharacterDerivedStats
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/characters/{id}/derived-stats [post]
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

// Character Response models for Swagger documentation
type CharacterSearchResponse struct {
	Characters []models.Character `json:"characters"`
	Count      int                `json:"count" example:"5"`
	Query      string             `json:"query" example:"fighter"`
}

type DeleteResponse struct {
	Message string `json:"message" example:"Character deleted successfully"`
}

type UpdateQuantityResponse struct {
	Message string `json:"message" example:"Quantity updated successfully"`
}

type AddSkillSpecialtyRequest struct {
	SkillSpecialtyID uuid.UUID `json:"skill_specialty_id" example:"123e4567-e89b-12d3-a456-426614174000"`
}

type AddSkillSpecialtyResponse struct {
	Message string `json:"message" example:"Character skill specialty added successfully"`
}
