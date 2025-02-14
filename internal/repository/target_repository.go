package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"sca/internal/domain"
)

type TargetPostgres struct {
	db *sqlx.DB
}

func NewTargetPostgres(db *sqlx.DB) *TargetPostgres {
	return &TargetPostgres{db: db}
}

func (r *TargetPostgres) GetAll() ([]domain.Target, error) {
	var targets []domain.Target

	query := fmt.Sprintf("SELECT * FROM %s", TargetsTable)
	err := r.db.Select(&targets, query)

	return targets, err
}

func (r *TargetPostgres) GetById(id int) (domain.Target, error) {
	var target domain.Target

	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", TargetsTable)
	err := r.db.Get(&target, query, id)

	return target, err
}
