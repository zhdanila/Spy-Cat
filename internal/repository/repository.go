package repository

import (
	"github.com/jmoiron/sqlx"
	"sca/internal/domain"
)

type Repository struct {
	SpyCat
	Mission
	Target
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		SpyCat:  NewSpyCatRepository(db),
		Mission: NewMissionRepository(db),
		Target:  NewTargetRepository(db),
	}
}

type SpyCat interface {
	CreateSpyCat(cat *domain.SpyCat) error
	DeleteSpyCat(catID int) error
	UpdateSpyCatSalary(catID int, salary float64) error
	GetSpyCat(catID int) (*domain.SpyCat, error)
	ListSpyCats() ([]domain.SpyCat, error)
}

type Mission interface {
	CreateMission(mission *domain.Mission, targets []domain.Target) (int, error)
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
