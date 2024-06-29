package service

import (
	"sca/internal/models"
	"sca/internal/repository"
)

type Service struct {
	Cat
	Mission
	Target
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Cat:     NewCatService(repo.Cat),
		Mission: NewMissionService(repo.Mission, repo.Target),
	}
}

type Cat interface {
	Create(cat models.Cat) (int, error)
	GetById(id int) (models.Cat, error)
	GetAll() ([]models.Cat, error)
	Update(id int, cat models.UpdatedCat) error
	Delete(id int) error
}

type Mission interface {
	Create(mission models.Mission) (int, error)
	CreateTarget(missionId int, target models.Target) (int, error)
	Update(id int, mission models.UpdatedMission) error
	Delete(id int) error
	DeleteTarget(missionId, targetId int) error
	GetByID(id int) (models.Mission, error)
	GetAll() ([]models.Mission, error)
}

type Target interface {
	GetAll() ([]models.Target, error)
	GetById(id int) (models.Target, error)
}
