package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"sca/internal/domain"
)

const MissionTable = "missions"

type MissionRepository struct {
	db *sqlx.DB
}

func (m MissionRepository) GetTX() (*sqlx.Tx, error) {
	return m.db.Beginx()
}

func (m MissionRepository) CreateMission(mission *domain.Mission, tx *sqlx.Tx) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (cat_id, is_completed, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at", MissionTable)

	err := tx.QueryRow(query, mission.CatID, mission.IsCompleted, mission.CreatedAt, mission.UpdatedAt).
		Scan(&mission.ID, &mission.CreatedAt, &mission.UpdatedAt)
	if err != nil {
		return 0, fmt.Errorf("could not create mission: %v", err)
	}

	return mission.ID, nil
}

func (m MissionRepository) DeleteMission(missionID int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", MissionTable)

	_, err := m.db.Exec(query, missionID)
	if err != nil {
		return fmt.Errorf("could not delete mission: %v", err)
	}

	return nil
}

func (m MissionRepository) UpdateMissionCompletion(missionID int, isCompleted bool) error {
	query := fmt.Sprintf("UPDATE %s SET is_completed = $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2", MissionTable)

	_, err := m.db.Exec(query, isCompleted, missionID)
	if err != nil {
		return fmt.Errorf("could not update mission completed status: %v", err)
	}

	return nil
}

func (m MissionRepository) GetMission(missionID int) (*domain.Mission, error) {
	query := fmt.Sprintf(`
		SELECT id, cat_id, is_completed, created_at, updated_at
		FROM %s
		WHERE id = $1`, MissionTable)

	var mission domain.Mission

	err := m.db.Get(&mission, query, missionID)
	if err != nil {
		return nil, fmt.Errorf("could not get mission: %v", err)
	}

	queryTargets := fmt.Sprintf(`
		SELECT id, mission_id, name, country, notes, is_completed, created_at, updated_at
		FROM %s
		WHERE mission_id = $1`, TargetTable)

	err = m.db.Select(&mission.Targets, queryTargets, missionID)
	if err != nil {
		return nil, fmt.Errorf("could not get targets for mission: %v", err)
	}

	return &mission, nil
}

func (m MissionRepository) ListMissions() ([]domain.Mission, error) {
	query := fmt.Sprintf(`
		SELECT id, cat_id, is_completed, created_at, updated_at
		FROM %s`, MissionTable)

	var missions []domain.Mission

	err := m.db.Select(&missions, query)
	if err != nil {
		return nil, fmt.Errorf("could not get missions: %v", err)
	}

	for i := range missions {
		queryTargets := fmt.Sprintf(`
        SELECT id, mission_id, name, country, notes, is_completed, created_at, updated_at
        FROM %s WHERE mission_id = $1`, TargetTable)

		err := m.db.Select(&missions[i].Targets, queryTargets, missions[i].ID)
		if err != nil {
			return nil, fmt.Errorf("could not get targets for mission %d: %v", missions[i].ID, err)
		}
	}

	return missions, nil
}

func (m MissionRepository) AssignSpyCatToMission(catID, missionID int) error {
	query := fmt.Sprintf("UPDATE %s SET cat_id = $1 WHERE id = $2", MissionTable)

	_, err := m.db.Exec(query, catID, missionID)
	if err != nil {
		return fmt.Errorf("could not assign spy cat to mission: %v", err)
	}

	return nil
}

func NewMissionRepository(db *sqlx.DB) *MissionRepository {
	return &MissionRepository{db: db}
}
