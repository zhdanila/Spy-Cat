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

func (t TargetRepository) GetTarget(targetID int) (*domain.Target, error) {
	var target domain.Target
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", TargetTable)

	err := t.db.Get(&target, query, targetID)
	if err != nil {
		return nil, fmt.Errorf("could not get target: %v", err)
	}

	return &target, nil
}

func (t TargetRepository) DeleteTarget(targetID int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", TargetTable)

	_, err := t.db.Exec(query, targetID)
	if err != nil {
		return fmt.Errorf("could not delete target: %v", err)
	}

	return nil
}

func (t TargetRepository) UpdateTargetCompletion(targetID int, isCompleted bool) error {
	query := fmt.Sprintf("UPDATE %s SET is_completed = $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2", TargetTable)

	_, err := t.db.Exec(query, isCompleted, targetID)
	if err != nil {
		return fmt.Errorf("could not update target completion: %v", err)
	}

	return nil
}

func (t TargetRepository) UpdateTargetNotes(targetID int, notes string) error {
	query := fmt.Sprintf("UPDATE %s SET notes = $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2", TargetTable)

	_, err := t.db.Exec(query, notes, targetID)
	if err != nil {
		return fmt.Errorf("could not update target notes: %v", err)
	}

	return nil
}

func NewTargetRepository(db *sqlx.DB) *TargetRepository {
	return &TargetRepository{db: db}
}
