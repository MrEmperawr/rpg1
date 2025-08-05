package repository

import (
	"github.com/mremperor-atwork/rpg1/api1/internal/database"
	"github.com/mremperor-atwork/rpg1/api1/internal/features/srd"
	"gorm.io/gorm"
)

type SRDRepository struct {
	db *gorm.DB
}

func NewSRDRepository() *SRDRepository {
	return &SRDRepository{db: database.GetDB()}
}

func (r *SRDRepository) GetAllEntries(category, search string, limit, offset int) ([]srd.SRDEntry, error) {
	var entries []srd.SRDEntry
	query := r.db.Model(&srd.SRDEntry{})

	if category != "" {
		query = query.Where("category = ?", category)
	}

	if search != "" {
		query = query.Where("title ILIKE ? OR content ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if limit > 0 {
		query = query.Limit(limit)
	}

	if offset > 0 {
		query = query.Offset(offset)
	}

	err := query.Order("category, title").Find(&entries).Error
	return entries, err
}

func (r *SRDRepository) GetEntryByID(id string) (*srd.SRDEntry, error) {
	var entry srd.SRDEntry
	err := r.db.Where("id = ?", id).First(&entry).Error
	if err != nil {
		return nil, err
	}
	return &entry, nil
}

func (r *SRDRepository) GetEntriesByCategory(category string) ([]srd.SRDEntry, error) {
	var entries []srd.SRDEntry
	err := r.db.Where("category = ?", category).Order("title").Find(&entries).Error
	return entries, err
}

func (r *SRDRepository) GetCategories() ([]string, error) {
	var categories []string
	err := r.db.Model(&srd.SRDEntry{}).Distinct().Pluck("category", &categories).Error
	return categories, err
}

func (r *SRDRepository) SearchEntries(query string, limit, offset int) ([]srd.SRDEntry, error) {
	var entries []srd.SRDEntry
	dbQuery := r.db.Model(&srd.SRDEntry{}).
		Where("title ILIKE ?", "%"+query+"%")

	if limit > 0 {
		dbQuery = dbQuery.Limit(limit)
	}

	if offset > 0 {
		dbQuery = dbQuery.Offset(offset)
	}

	err := dbQuery.Order("category, title").Find(&entries).Error
	return entries, err
}

func (r *SRDRepository) GetContentByTitle(title string) (*srd.SRDContent, error) {
	var content srd.SRDContent
	err := r.db.Joins("JOIN srd_entries ON srd_content.entry_id = srd_entries.id").
		Where("srd_entries.title = ?", title).
		First(&content).Error
	if err != nil {
		return nil, err
	}
	return &content, nil
}

func (r *SRDRepository) GetAllContent(category, search string, limit, offset int) ([]srd.SRDContent, error) {
	var content []srd.SRDContent
	query := r.db.Model(&srd.SRDContent{}).
		Joins("JOIN srd_entries ON srd_content.entry_id = srd_entries.id")

	if category != "" {
		query = query.Where("srd_entries.category = ?", category)
	}

	if search != "" {
		query = query.Where("srd_content.content ILIKE ?", "%"+search+"%")
	}

	if limit > 0 {
		query = query.Limit(limit)
	}

	if offset > 0 {
		query = query.Offset(offset)
	}

	err := query.Order("srd_entries.category, srd_entries.title").Find(&content).Error
	return content, err
}
