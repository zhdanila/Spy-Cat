package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"sca/internal/domain"
)

const TargetTable = "targets"

type TargetRepository struct {
	db *sqlx.DB
}

func (t TargetRepository) AddTargetsToMission(missionID int, targets []domain.Target) error {
	query := fmt.Sprintf("INSERT INTO %s (mission_id, name, country, notes, is_completed, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)", TargetTable)

	for _, target := range targets {
		_, err := t.db.Exec(query, missionID, target.Name, target.Country, target.Notes, target.IsCompleted, target.CreatedAt, target.UpdatedAt)
		if err != nil {
			return fmt.Errorf("could not insert target: %v", err)
		}
	}
	return nil
}

func (t TargetRepository) AddTargetsToMissionTX(missionID int, targets []domain.Target, tx *sqlx.Tx) error {
	query := fmt.Sprintf("INSERT INTO %s (mission_id, name, country, notes, is_completed, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)", TargetTable)

	for _, target := range targets {
		_, err := tx.Exec(query, missionID, target.Name, target.Country, target.Notes, target.IsCompleted, target.CreatedAt, target.UpdatedAt)
		if err != nil {
			return fmt.Errorf("could not insert target: %v", err)
		}
	}
	return nil
}

func (t TargetRepository) DeleteTarget(targetID int) error {
	//TODO implement me
	panic("implement me")
}

func (t TargetRepository) UpdateTargetCompletion(targetID int, isCompleted bool) error {
	//TODO implement me
	panic("implement me")
}

func (t TargetRepository) UpdateTargetNotes(targetID int, notes string) error {
	//TODO implement me
	panic("implement me")
}

func NewTargetRepository(db *sqlx.DB) *TargetRepository {
	return &TargetRepository{db: db}
}
