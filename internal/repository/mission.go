package repository

import (
	"github.com/jmoiron/sqlx"
	"sca/internal/domain"
)

type MissionRepository struct {
	db *sqlx.DB
}

func (m MissionRepository) CreateMission(mission *domain.Mission, targets []domain.Target) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (m MissionRepository) DeleteMission(missionID int) error {
	//TODO implement me
	panic("implement me")
}

func (m MissionRepository) UpdateMissionCompletion(missionID int, isCompleted bool) error {
	//TODO implement me
	panic("implement me")
}

func (m MissionRepository) GetMission(missionID int) (*domain.Mission, error) {
	//TODO implement me
	panic("implement me")
}

func (m MissionRepository) ListMissions() ([]domain.Mission, error) {
	//TODO implement me
	panic("implement me")
}

func (m MissionRepository) AssignSpyCatToMission(catID, missionID int) error {
	//TODO implement me
	panic("implement me")
}

func NewMissionRepository(db *sqlx.DB) *MissionRepository {
	return &MissionRepository{db: db}
}
