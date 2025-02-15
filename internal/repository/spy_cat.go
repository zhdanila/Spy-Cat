package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"sca/internal/domain"
)

const SpyCatTable = "spy_cats"

type SpyCatRepository struct {
	db *sqlx.DB
}

func NewSpyCatRepository(db *sqlx.DB) *SpyCatRepository {
	return &SpyCatRepository{db: db}
}

func (s *SpyCatRepository) CreateSpyCat(cat *domain.SpyCat) error {
	query := fmt.Sprintf("INSERT INTO %s (name, years_of_experience, breed, salary) VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at", SpyCatTable)

	err := s.db.QueryRow(query, cat.Name, cat.YearsOfExperience, cat.Breed, cat.Salary).Scan(&cat.ID, &cat.CreatedAt, &cat.UpdatedAt)
	if err != nil {
		return fmt.Errorf("could not create spy cat: %v", err)
	}
	return nil
}

func (s *SpyCatRepository) DeleteSpyCat(catID int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", SpyCatTable)

	_, err := s.db.Exec(query, catID)
	if err != nil {
		return fmt.Errorf("could not delete spy cat: %v", err)
	}

	return nil
}

func (s *SpyCatRepository) UpdateSpyCatSalary(catID int, salary float64) error {
	query := fmt.Sprintf("UPDATE %s SET salary = $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2", SpyCatTable)

	_, err := s.db.Exec(query, salary, catID)
	if err != nil {
		return fmt.Errorf("could not update spy cat salary: %v", err)
	}

	return nil
}

func (s *SpyCatRepository) GetSpyCat(catID int) (*domain.SpyCat, error) {
	var cat domain.SpyCat
	query := fmt.Sprintf("SELECT id, name, years_of_experience, breed, salary, created_at, updated_at FROM %s WHERE id = $1", SpyCatTable)

	err := s.db.Get(&cat, query, catID)
	if err != nil {
		return nil, fmt.Errorf("could not get spy cat: %v", err)
	}

	return &cat, nil
}

func (s *SpyCatRepository) ListSpyCats() ([]domain.SpyCat, error) {
	var cats []domain.SpyCat
	query := fmt.Sprintf("SELECT id, name, years_of_experience, breed, salary, created_at, updated_at FROM %s", SpyCatTable)

	err := s.db.Select(&cats, query)
	if err != nil {
		return nil, fmt.Errorf("could not get list of spy cats: %v", err)
	}

	return cats, nil
}
