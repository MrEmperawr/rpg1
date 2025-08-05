package repository

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/mremperor-atwork/rpg1/api1/internal/database"
	"github.com/mremperor-atwork/rpg1/api1/internal/features/game"
	"github.com/mremperor-atwork/rpg1/api1/internal/models"
)

type CharacterRepository struct {
	db *gorm.DB
}

func NewCharacterRepository() *CharacterRepository {
	return &CharacterRepository{
		db: database.GetDB(),
	}
}

func (r *CharacterRepository) CreateCharacter(character *models.Character) error {
	return r.db.Create(character).Error
}

func (r *CharacterRepository) GetCharacterByID(id uuid.UUID) (*models.Character, error) {
	var character models.Character
	err := r.db.Preload("User").Preload("Campaign").First(&character, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &character, nil
}

func (r *CharacterRepository) GetCharactersByUserID(userID uuid.UUID) ([]models.Character, error) {
	var characters []models.Character
	err := r.db.Preload("Campaign").Where("user_id = ?", userID).Find(&characters).Error
	return characters, err
}

func (r *CharacterRepository) GetCharactersByCampaignID(campaignID uuid.UUID) ([]models.Character, error) {
	var characters []models.Character
	err := r.db.Preload("User").Where("campaign_id = ?", campaignID).Find(&characters).Error
	return characters, err
}

func (r *CharacterRepository) UpdateCharacter(character *models.Character) error {
	return r.db.Save(character).Error
}

func (r *CharacterRepository) DeleteCharacter(id uuid.UUID) error {
	return r.db.Delete(&models.Character{}, "id = ?", id).Error
}

func (r *CharacterRepository) GetCharacterAttributes(characterID uuid.UUID) ([]game.CharacterAttribute, error) {
	var attributes []game.CharacterAttribute
	err := r.db.Preload("Attribute").Where("character_id = ?", characterID).Find(&attributes).Error
	return attributes, err
}

func (r *CharacterRepository) SetCharacterAttribute(characterAttribute *game.CharacterAttribute) error {
	var existing game.CharacterAttribute
	err := r.db.Where("character_id = ? AND attribute_id = ?", characterAttribute.CharacterID, characterAttribute.AttributeID).First(&existing).Error

	if err == gorm.ErrRecordNotFound {
		return r.db.Create(characterAttribute).Error
	} else if err != nil {
		return err
	}

	existing.Value = characterAttribute.Value
	return r.db.Save(&existing).Error
}

func (r *CharacterRepository) DeleteCharacterAttribute(characterID, attributeID uuid.UUID) error {
	return r.db.Where("character_id = ? AND attribute_id = ?", characterID, attributeID).Delete(&game.CharacterAttribute{}).Error
}

func (r *CharacterRepository) GetCharacterSkills(characterID uuid.UUID) ([]game.CharacterSkill, error) {
	var skills []game.CharacterSkill
	err := r.db.Preload("Skill").Where("character_id = ?", characterID).Find(&skills).Error
	return skills, err
}

func (r *CharacterRepository) SetCharacterSkill(characterSkill *game.CharacterSkill) error {
	var existing game.CharacterSkill
	err := r.db.Where("character_id = ? AND skill_id = ?", characterSkill.CharacterID, characterSkill.SkillID).First(&existing).Error

	if err == gorm.ErrRecordNotFound {
		return r.db.Create(characterSkill).Error
	} else if err != nil {
		return err
	}

	existing.Value = characterSkill.Value
	return r.db.Save(&existing).Error
}

func (r *CharacterRepository) DeleteCharacterSkill(characterID, skillID uuid.UUID) error {
	return r.db.Where("character_id = ? AND skill_id = ?", characterID, skillID).Delete(&game.CharacterSkill{}).Error
}

func (r *CharacterRepository) GetCharacterSkillSpecialties(characterID uuid.UUID) ([]game.CharacterSkillSpecialty, error) {
	var specialties []game.CharacterSkillSpecialty
	err := r.db.Preload("SkillSpecialty").Preload("SkillSpecialty.Skill").Where("character_id = ?", characterID).Find(&specialties).Error
	return specialties, err
}

func (r *CharacterRepository) AddCharacterSkillSpecialty(characterID, skillSpecialtyID uuid.UUID) error {
	specialty := game.CharacterSkillSpecialty{
		CharacterID:      characterID,
		SkillSpecialtyID: skillSpecialtyID,
	}
	return r.db.Create(&specialty).Error
}

func (r *CharacterRepository) RemoveCharacterSkillSpecialty(characterID, skillSpecialtyID uuid.UUID) error {
	return r.db.Where("character_id = ? AND skill_specialty_id = ?", characterID, skillSpecialtyID).Delete(&game.CharacterSkillSpecialty{}).Error
}

func (r *CharacterRepository) GetCharacterQualities(characterID uuid.UUID) ([]game.CharacterQuality, error) {
	var qualities []game.CharacterQuality
	err := r.db.Preload("Quality").Where("character_id = ?", characterID).Find(&qualities).Error
	return qualities, err
}

func (r *CharacterRepository) SetCharacterQuality(characterQuality *game.CharacterQuality) error {
	var existing game.CharacterQuality
	err := r.db.Where("character_id = ? AND quality_id = ?", characterQuality.CharacterID, characterQuality.QualityID).First(&existing).Error

	if err == gorm.ErrRecordNotFound {
		return r.db.Create(characterQuality).Error
	} else if err != nil {
		return err
	}

	existing.Rating = characterQuality.Rating
	return r.db.Save(&existing).Error
}

func (r *CharacterRepository) DeleteCharacterQuality(characterID, qualityID uuid.UUID) error {
	return r.db.Where("character_id = ? AND quality_id = ?", characterID, qualityID).Delete(&game.CharacterQuality{}).Error
}

func (r *CharacterRepository) GetCharacterEquipment(characterID uuid.UUID) ([]game.CharacterEquipment, error) {
	var equipment []game.CharacterEquipment
	err := r.db.Preload("Equipment").Where("character_id = ?", characterID).Find(&equipment).Error
	return equipment, err
}

func (r *CharacterRepository) AddCharacterEquipment(characterEquipment *game.CharacterEquipment) error {
	var existing game.CharacterEquipment
	err := r.db.Where("character_id = ? AND equipment_id = ?", characterEquipment.CharacterID, characterEquipment.EquipmentID).First(&existing).Error

	if err == gorm.ErrRecordNotFound {
		return r.db.Create(characterEquipment).Error
	} else if err != nil {
		return err
	}

	existing.Quantity += characterEquipment.Quantity
	return r.db.Save(&existing).Error
}

func (r *CharacterRepository) UpdateCharacterEquipmentQuantity(characterID, equipmentID uuid.UUID, quantity int) error {
	return r.db.Model(&game.CharacterEquipment{}).Where("character_id = ? AND equipment_id = ?", characterID, equipmentID).Update("quantity", quantity).Error
}

func (r *CharacterRepository) RemoveCharacterEquipment(characterID, equipmentID uuid.UUID) error {
	return r.db.Where("character_id = ? AND equipment_id = ?", characterID, equipmentID).Delete(&game.CharacterEquipment{}).Error
}

func (r *CharacterRepository) GetCharacterPersonalEquipment(characterID uuid.UUID) ([]game.CharacterPersonalEquipment, error) {
	var equipment []game.CharacterPersonalEquipment
	err := r.db.Preload("PersonalEquipment").Where("character_id = ?", characterID).Find(&equipment).Error
	return equipment, err
}

func (r *CharacterRepository) AddCharacterPersonalEquipment(characterEquipment *game.CharacterPersonalEquipment) error {
	var existing game.CharacterPersonalEquipment
	err := r.db.Where("character_id = ? AND personal_equipment_id = ?", characterEquipment.CharacterID, characterEquipment.PersonalEquipmentID).First(&existing).Error

	if err == gorm.ErrRecordNotFound {
		return r.db.Create(characterEquipment).Error
	} else if err != nil {
		return err
	}

	existing.Quantity += characterEquipment.Quantity
	return r.db.Save(&existing).Error
}

func (r *CharacterRepository) UpdateCharacterPersonalEquipmentQuantity(characterID, equipmentID uuid.UUID, quantity int) error {
	return r.db.Model(&game.CharacterPersonalEquipment{}).Where("character_id = ? AND personal_equipment_id = ?", characterID, equipmentID).Update("quantity", quantity).Error
}

func (r *CharacterRepository) RemoveCharacterPersonalEquipment(characterID, equipmentID uuid.UUID) error {
	return r.db.Where("character_id = ? AND personal_equipment_id = ?", characterID, equipmentID).Delete(&game.CharacterPersonalEquipment{}).Error
}

func (r *CharacterRepository) GetCharacterDerivedStats(characterID uuid.UUID) (*game.CharacterDerivedStats, error) {
	var stats game.CharacterDerivedStats
	err := r.db.Where("character_id = ?", characterID).First(&stats).Error
	if err != nil {
		return nil, err
	}
	return &stats, nil
}

func (r *CharacterRepository) SetCharacterDerivedStats(stats *game.CharacterDerivedStats) error {
	var existing game.CharacterDerivedStats
	err := r.db.Where("character_id = ?", stats.CharacterID).First(&existing).Error

	if err == gorm.ErrRecordNotFound {
		return r.db.Create(stats).Error
	} else if err != nil {
		return err
	}

	return r.db.Model(&existing).Updates(stats).Error
}

func (r *CharacterRepository) SearchCharacters(userID uuid.UUID, query string) ([]models.Character, error) {
	var characters []models.Character
	err := r.db.Preload("Campaign").Preload("Species").
		Where("user_id = ? AND (name ILIKE ? OR concept ILIKE ?)", userID, "%"+query+"%", "%"+query+"%").
		Find(&characters).Error
	return characters, err
}

func (r *CharacterRepository) ValidateCharacterCreation(character *models.Character) error {
	var count int64
	err := r.db.Model(&models.Character{}).Where("user_id = ? AND name = ?", character.UserID, character.Name).Count(&count).Error
	if err != nil {
		return fmt.Errorf("failed to validate character name: %w", err)
	}

	if count > 0 {
		return fmt.Errorf("character name '%s' already exists for this user", character.Name)
	}

	return nil
}
