package repository

import (
	"github.com/jmoiron/sqlx"
	"sca/internal/models"
)

const (
	CatsTable     string = "cats"
	MissionsTable string = "missions"
	TargetsTable  string = "targets"
)

type Repository struct {
	Cat
	Mission
	Target
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Cat:     NewCatPostgres(db),
		Mission: NewMissionPostgres(db),
		Target:  NewTargetPostgres(db),
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
	Update(id int, mission models.UpdatedMission) error
	Delete(id int) error
	GetByID(id int) (models.Mission, error)
	GetAll() ([]models.Mission, error)
	DeleteTarget(missionId, targetId int) error
	CreateTarget(missionId int, target models.Target) (int, error)
}

type Target interface {
	GetAll() ([]models.Target, error)
	GetById(id int) (models.Target, error)
}
