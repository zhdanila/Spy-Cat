package service

import (
	"sca/internal/domain"
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
	Create(cat domain.Cat) (int, error)
	GetById(id int) (domain.Cat, error)
	GetAll() ([]domain.Cat, error)
	Update(id int, cat domain.UpdatedCat) error
	Delete(id int) error
}

type Mission interface {
	Create(mission domain.Mission) (int, error)
	CreateTarget(missionId int, target domain.Target) (int, error)
	Update(id int, mission domain.UpdatedMission) error
	Delete(id int) error
	DeleteTarget(missionId, targetId int) error
	GetByID(id int) (domain.Mission, error)
	GetAll() ([]domain.Mission, error)
}

type Target interface {
	GetAll() ([]domain.Target, error)
	GetById(id int) (domain.Target, error)
}
