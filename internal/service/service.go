package service

import (
	"sca/internal/domain"
	"sca/internal/repository"
	"sca/internal/service/mission"
	"sca/internal/service/spy_cat"
)

type Service struct {
	SpyCat
	Mission
	Target
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		SpyCat:  spy_cat.NewService(repo.SpyCat),
		Mission: mission.NewService(repo.Mission, repo.Target),
		Target:  mission.NewService(repo.Mission, repo.Target),
	}
}

type SpyCat interface {
	CreateSpyCat(name string, yearsOfExperience int, breed string, salary float64) (int, error)
	DeleteSpyCat(catID int) error
	UpdateSpyCatSalary(catID int, salary float64) error
	GetSpyCat(catID int) (*domain.SpyCat, error)
	ListSpyCats() ([]domain.SpyCat, error)
}

type Mission interface {
	CreateMission(catID int, targets []domain.Target) (int, error)
	DeleteMission(missionID int) error
	UpdateMissionCompletion(missionID int, isCompleted bool) error
	GetMission(missionID int) (*domain.Mission, error)
	ListMissions() ([]domain.Mission, error)
	AssignSpyCatToMission(catID, missionID int) error
}

type Target interface {
	AddTargetsToMission(missionID int, targets []domain.Target) error
	DeleteTarget(targetID int) error
	UpdateTargetCompletion(targetID int, isCompleted bool) error
	UpdateTargetNotes(targetID int, notes string) error
}
