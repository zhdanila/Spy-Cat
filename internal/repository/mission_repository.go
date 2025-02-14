package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"sca/internal/domain"
)

type MissionPostgres struct {
	db *sqlx.DB
}

func NewMissionPostgres(db *sqlx.DB) *MissionPostgres {
	return &MissionPostgres{db: db}
}

func (r *MissionPostgres) Create(mission domain.Mission) (int, error) {
	var missionID int

	query := fmt.Sprintf("INSERT INTO %s (cat_id, complete) VALUES ($1, $2) RETURNING id", MissionsTable)

	err := r.db.QueryRow(query, mission.CatID, mission.Complete).Scan(&missionID)
	if err != nil {
		return 0, err
	}

	query = fmt.Sprintf("INSERT INTO %s (mission_id, name, country, notes, complete) VALUES ($1, $2, $3, $4, $5)", TargetsTable)

	for _, target := range mission.Targets {
		_, err = r.db.Exec(query, missionID, target.Name, target.Country, target.Notes, target.Complete)
		if err != nil {
			return 0, err
		}
	}

	return missionID, err
}

func (r *MissionPostgres) Update(id int, mission domain.UpdatedMission) error {
	//get mission to check parameters
	var missionToCheck domain.Mission

	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", MissionsTable)

	err := r.db.Get(&missionToCheck, query, id)
	if err != nil {
		return err
	}

	//update mission complete
	if mission.Complete != missionToCheck.Complete {
		query := fmt.Sprintf("UPDATE %s SET complete=$1 WHERE id=$2", MissionsTable)
		_, err := r.db.Exec(query, mission.Complete, id)

		if err != nil {
			return err
		}
	}

	//update notes
	if mission.Notes != "" {
		query := fmt.Sprintf("UPDATE %s SET notes=$1 WHERE mission_id=$2", TargetsTable)
		_, err := r.db.Exec(query, mission.Notes, id)

		if err != nil {
			return err
		}
	}

	//update target complete state
	if mission.Complete {
		query := fmt.Sprintf("UPDATE %s SET complete=$1 WHERE mission_id=$2", TargetsTable)
		_, err := r.db.Exec(query, mission.Complete, id)

		if err != nil {
			return err
		}
	}

	//update cat id
	if mission.CatID != 0 {
		query := fmt.Sprintf("UPDATE %s SET cat_id=$1 WHERE id=$2", MissionsTable)
		_, err := r.db.Exec(query, mission.CatID, id)

		if err != nil {
			return err
		}
	}

	return nil
}

func (r *MissionPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE mission_id = $1", TargetsTable)
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	query = fmt.Sprintf("DELETE FROM %s WHERE id = $1", MissionsTable)
	_, err = r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return err
}

func (r *MissionPostgres) GetByID(id int) (domain.Mission, error) {
	var mission domain.Mission

	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", MissionsTable)
	err := r.db.Get(&mission, query, id)

	return mission, err
}

func (r *MissionPostgres) GetAll() ([]domain.Mission, error) {
	var missions []domain.Mission

	query := fmt.Sprintf("SELECT * FROM %s", MissionsTable)
	err := r.db.Select(&missions, query)

	return missions, err
}

func (r *MissionPostgres) DeleteTarget(missionId, targetId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1 AND mission_id = $2", TargetsTable)
	_, err := r.db.Exec(query, targetId, missionId)

	return err
}

func (r *MissionPostgres) CreateTarget(missionId int, target domain.Target) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (mission_id, name, country, notes, complete)"+
		" VALUES ($1, $2, $3, $4, $5) RETURNING ID", TargetsTable)
	err := r.db.QueryRow(query, missionId, target.Name, target.Country, target.Notes, target.Complete).Scan(&id)

	return id, err
}
