package repository

import (
	"github.com/jmoiron/sqlx"
	"sca/internal/domain"
)

type SpyCatRepository struct {
	db *sqlx.DB
}

func (s SpyCatRepository) CreateSpyCat(cat *domain.SpyCat) error {
	//TODO implement me
	panic("implement me")
}

func (s SpyCatRepository) DeleteSpyCat(catID int) error {
	//TODO implement me
	panic("implement me")
}

func (s SpyCatRepository) UpdateSpyCatSalary(catID int, salary float64) error {
	//TODO implement me
	panic("implement me")
}

func (s SpyCatRepository) GetSpyCat(catID int) (*domain.SpyCat, error) {
	//TODO implement me
	panic("implement me")
}

func (s SpyCatRepository) ListSpyCats() ([]domain.SpyCat, error) {
	//TODO implement me
	panic("implement me")
}

func NewSpyCatRepository(db *sqlx.DB) *SpyCatRepository {
	return &SpyCatRepository{db: db}
}
