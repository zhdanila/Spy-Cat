package repository

import (
	"github.com/jmoiron/sqlx"
	"sca/internal/domain"
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
	Create(cat domain.Cat) (int, error)
	GetById(id int) (domain.Cat, error)
	GetAll() ([]domain.Cat, error)
	Update(id int, cat domain.UpdatedCat) error
	Delete(id int) error
}

type Mission interface {
	Create(mission domain.Mission) (int, error)
	Update(id int, mission domain.UpdatedMission) error
	Delete(id int) error
	GetByID(id int) (domain.Mission, error)
	GetAll() ([]domain.Mission, error)
	DeleteTarget(missionId, targetId int) error
	CreateTarget(missionId int, target domain.Target) (int, error)
}

type Target interface {
	GetAll() ([]domain.Target, error)
	GetById(id int) (domain.Target, error)
}
