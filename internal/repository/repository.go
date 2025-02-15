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
	CreateMission(mission *domain.Mission, tx *sqlx.Tx) (int, error)
	DeleteMission(missionID int) error
	UpdateMissionCompletion(missionID int, isCompleted bool) error
	GetMission(missionID int) (*domain.Mission, error)
	ListMissions() ([]domain.Mission, error)
	AssignSpyCatToMission(catID, missionID int) error
	GetTX() (*sqlx.Tx, error)
}

type Target interface {
	GetTarget(targetID int) (*domain.Target, error)
	AddTargetsToMissionTX(missionID int, targets []domain.Target, tx *sqlx.Tx) error
	AddTargetsToMission(missionID int, targets []domain.Target) error
	DeleteTarget(targetID int) error
	UpdateTargetCompletion(targetID int, isCompleted bool) error
	UpdateTargetNotes(targetID int, notes string) error
}
