package chart

import (
	"go-gorm-hierarchy/models"
	"go-gorm-hierarchy/repositories"

	"gorm.io/gorm"
)

// Repository repository interface
type Repository interface {
	FindRoot(db *gorm.DB) ([]*models.Chart, error)
	FindFullHierarchy(db *gorm.DB, id interface{}, object interface{}, condition map[string]interface{}) bool
}

type repository struct {
	repositories.Repository
}

// NewRepository new repository
func NewRepository() Repository {
	return &repository{
		repositories.NewRepository(),
	}
}

// FindRoot find root
func (r *repository) FindRoot(db *gorm.DB) ([]*models.Chart, error) {
	entities := []*models.Chart{}
	// if there are many other associations, you could use clause.Associations, it will preload all first level associations
	// to preload the next level, use it like "Children." + clause.Associations, for example:
	// Preload(clause.Associations).Preload("Children." + clause.Associations)
	err := db.Where("group_id IS NULL").Find(&entities).Error
	if err != nil {
		return nil, err
	}

	return entities, nil
}
