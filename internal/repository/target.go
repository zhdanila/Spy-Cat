package repository

import (
	"github.com/jmoiron/sqlx"
	"sca/internal/domain"
)

type TargetRepository struct {
	db *sqlx.DB
}

func (t TargetRepository) AddTargetsToMission(missionID int, targets []domain.Target) error {
	//TODO implement me
	panic("implement me")
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
