package chart

import (
	"go-gorm-hierarchy/db"
	"go-gorm-hierarchy/models"

	"github.com/sirupsen/logrus"
)

// Service service interface
type Service interface {
	GetAll() ([]*models.Chart, error)
}

type service struct {
	repository Repository
}

// NewService new service
func NewService() Service {
	return &service{
		repository: NewRepository(),
	}
}

// Get all
func (s *service) GetAll() ([]*models.Chart, error) {
	entities, err := s.repository.FindRoot(db.Database)
	if err != nil {
		logrus.Errorf("find all error: %s", err)
		return nil, err
	}

	for i := range entities {
		object := &models.Chart{}
		condition := map[string]interface{}{}
		found := s.repository.FindFullHierarchy(db.Database, entities[i].ID, object, condition)
		if found {
			entities[i] = object
		}
	}

	return entities, nil

}
