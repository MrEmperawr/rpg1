package repository

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/mremperor-atwork/rpg1/api1/internal/database"
	"github.com/mremperor-atwork/rpg1/api1/internal/models"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: database.GetDB(),
	}
}

// User CRUD Operations
func (r *UserRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) GetUserByID(id uuid.UUID) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) UpdateUser(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepository) DeleteUser(id uuid.UUID) error {
	return r.db.Delete(&models.User{}, "id = ?", id).Error
}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *UserRepository) SearchUsers(query string) ([]models.User, error) {
	var users []models.User
	err := r.db.Where("email ILIKE ? OR first_name ILIKE ? OR last_name ILIKE ? OR display_name ILIKE ?",
		"%"+query+"%", "%"+query+"%", "%"+query+"%", "%"+query+"%").Find(&users).Error
	return users, err
}

// User Validation
func (r *UserRepository) ValidateUserCreation(user *models.User) error {
	// Check if email is unique
	var count int64
	err := r.db.Model(&models.User{}).Where("email = ?", user.Email).Count(&count).Error
	if err != nil {
		return fmt.Errorf("failed to validate email uniqueness: %w", err)
	}
	if count > 0 {
		return fmt.Errorf("email already exists")
	}

	// Check if display name is unique (if provided)
	if user.DisplayName != "" {
		err = r.db.Model(&models.User{}).Where("display_name = ?", user.DisplayName).Count(&count).Error
		if err != nil {
			return fmt.Errorf("failed to validate display name uniqueness: %w", err)
		}
		if count > 0 {
			return fmt.Errorf("display name already exists")
		}
	}

	return nil
}

func (r *UserRepository) ValidateUserUpdate(user *models.User) error {
	// Check if email is unique (excluding current user)
	var count int64
	err := r.db.Model(&models.User{}).Where("email = ? AND id != ?", user.Email, user.ID).Count(&count).Error
	if err != nil {
		return fmt.Errorf("failed to validate email uniqueness: %w", err)
	}
	if count > 0 {
		return fmt.Errorf("email already exists")
	}

	// Check if display name is unique (if provided, excluding current user)
	if user.DisplayName != "" {
		err = r.db.Model(&models.User{}).Where("display_name = ? AND id != ?", user.DisplayName, user.ID).Count(&count).Error
		if err != nil {
			return fmt.Errorf("failed to validate display name uniqueness: %w", err)
		}
		if count > 0 {
			return fmt.Errorf("display name already exists")
		}
	}

	return nil
}

// User Statistics
func (r *UserRepository) GetUserStats(userID uuid.UUID) (map[string]interface{}, error) {
	stats := make(map[string]interface{})

	// Count characters
	var characterCount int64
	err := r.db.Model(&models.Character{}).Where("user_id = ?", userID).Count(&characterCount).Error
	if err != nil {
		return nil, fmt.Errorf("failed to count characters: %w", err)
	}
	stats["character_count"] = characterCount

	// Count campaigns (as GM)
	var campaignCount int64
	err = r.db.Model(&models.Campaign{}).Where("gm_user_id = ?", userID).Count(&campaignCount).Error
	if err != nil {
		return nil, fmt.Errorf("failed to count campaigns: %w", err)
	}
	stats["campaign_count"] = campaignCount

	// Get user creation date
	var user models.User
	err = r.db.Select("created_at").First(&user, "id = ?", userID).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get user creation date: %w", err)
	}
	stats["created_at"] = user.CreatedAt

	return stats, nil
}

// User Activity
func (r *UserRepository) GetUserActivity(userID uuid.UUID, limit int) ([]map[string]interface{}, error) {
	var activities []map[string]interface{}

	// Get recent characters created
	var characters []models.Character
	err := r.db.Where("user_id = ?", userID).Order("created_at DESC").Limit(limit).Find(&characters).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get recent characters: %w", err)
	}

	for _, char := range characters {
		activities = append(activities, map[string]interface{}{
			"type":        "character_created",
			"id":          char.ID,
			"name":        char.Name,
			"created_at":  char.CreatedAt,
			"description": fmt.Sprintf("Created character '%s'", char.Name),
		})
	}

	// Get recent campaigns (as GM)
	var campaigns []models.Campaign
	err = r.db.Where("gm_user_id = ?", userID).Order("created_at DESC").Limit(limit).Find(&campaigns).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get recent campaigns: %w", err)
	}

	for _, campaign := range campaigns {
		activities = append(activities, map[string]interface{}{
			"type":        "campaign_created",
			"id":          campaign.ID,
			"name":        campaign.Name,
			"created_at":  campaign.CreatedAt,
			"description": fmt.Sprintf("Created campaign '%s'", campaign.Name),
		})
	}

	// Sort activities by created_at (most recent first)
	// Note: In a real implementation, you might want to use a dedicated activity table
	// This is a simplified version that combines character and campaign activities

	return activities, nil
}

// User Preferences (placeholder for future implementation)
func (r *UserRepository) GetUserPreferences(userID uuid.UUID) (map[string]interface{}, error) {
	// This would typically query a user_preferences table
	// For now, return default preferences
	return map[string]interface{}{
		"theme":           "default",
		"language":        "en",
		"notifications":   true,
		"email_alerts":    true,
		"default_era":     "Medieval",
		"character_limit": 10,
	}, nil
}

func (r *UserRepository) UpdateUserPreferences(userID uuid.UUID, preferences map[string]interface{}) error {
	// This would typically update a user_preferences table
	// For now, just return success
	return nil
}

// User Authentication (placeholder for future implementation)
func (r *UserRepository) AuthenticateUser(email, password string) (*models.User, error) {
	// This would typically hash the password and compare with stored hash
	// For now, just return the user if email exists
	user, err := r.GetUserByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	// TODO: Implement proper password hashing and verification
	// For now, assume authentication is successful if user exists
	return user, nil
}

func (r *UserRepository) ChangePassword(userID uuid.UUID, oldPassword, newPassword string) error {
	// This would typically verify old password and hash new password
	// For now, just return success
	return nil
}

// User Session Management (placeholder for future implementation)
func (r *UserRepository) CreateUserSession(userID uuid.UUID, sessionData map[string]interface{}) (string, error) {
	// This would typically create a session record and return session token
	// For now, return a placeholder token
	return uuid.New().String(), nil
}

func (r *UserRepository) ValidateUserSession(sessionToken string) (*models.User, error) {
	// This would typically validate session token and return associated user
	// For now, return error
	return nil, fmt.Errorf("session validation not implemented")
}

func (r *UserRepository) DeleteUserSession(sessionToken string) error {
	// This would typically delete session record
	// For now, just return success
	return nil
}
