package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"sca/internal/models"
)

type CatPostgres struct {
	db *sqlx.DB
}

func NewCatPostgres(db *sqlx.DB) *CatPostgres {
	return &CatPostgres{db: db}
}

func (r *CatPostgres) Create(cat models.Cat) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (name, years_of_experience, breed, salary)" +
		" VALUES ($1, $2, $3, $4) RETURNING ID", CatsTable)
	err := r.db.QueryRow(query, cat.Name, cat.YearsOfExperience, cat.Breed, cat.Salary).Scan(&id)

	return id, err
}

func (r *CatPostgres) GetById(id int) (models.Cat, error) {
	var cat models.Cat

	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", CatsTable)
	err := r.db.Get(&cat, query, id)

	return cat, err
}

func (r *CatPostgres) GetAll() ([]models.Cat, error) {
	var cats []models.Cat

	query := fmt.Sprintf("SELECT * FROM %s", CatsTable)
	err := r.db.Select(&cats, query)

	return cats, err
}

func (r *CatPostgres) Update(id int, cat models.UpdatedCat) error {
	query := fmt.Sprintf("UPDATE %s SET salary=$1 WHERE id=$2", CatsTable)
	_, err := r.db.Exec(query, cat.Salary, id)

	return err
}

func (r *CatPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1",
		CatsTable)
	_, err := r.db.Exec(query, id)

	return err
}
