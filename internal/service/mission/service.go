package mission

import (
	"fmt"
	"sca/internal/domain"
	"sca/internal/repository"
	"sca/internal/service/target"
	"time"
)

type Service struct {
	missionRepo repository.Mission
	targetsRepo repository.Target
}

func NewService(missionRepo repository.Mission, targetsRepo repository.Target) *Service {
	return &Service{missionRepo: missionRepo, targetsRepo: targetsRepo}
}

func (s Service) CreateMission(req *CreateMissionRequest) (*CreateMissionResponse, error) {
	mission := domain.Mission{
		CatID:       &req.CatID,
		IsCompleted: false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	targets := make([]domain.Target, 0)
	for _, target := range req.Targets {
		targets = append(targets, domain.Target{
			Name:        target.Name,
			Country:     target.Country,
			Notes:       target.Notes,
			IsCompleted: target.IsCompleted,
		})
	}

	tx, err := s.missionRepo.GetTX()
	if err != nil {
		return nil, fmt.Errorf("could not start transaction: %v", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	missionID, err := s.missionRepo.CreateMission(&mission, tx)
	if err != nil {
		return nil, err
	}

	if err = s.targetsRepo.AddTargetsToMissionTX(mission.ID, targets, tx); err != nil {
		return nil, err
	}

	return &CreateMissionResponse{ID: missionID}, nil
}

func (s Service) GetMission(req *GetMissionRequest) (*GetMissionResponse, error) {
	mission, err := s.missionRepo.GetMission(req.ID)
	if err != nil {
		return nil, err
	}

	targets := make([]target.Target, 0)
	for _, t := range mission.Targets {
		targets = append(targets, target.Target{
			ID:          t.ID,
			MissionID:   t.MissionID,
			Name:        t.Name,
			Country:     t.Country,
			Notes:       t.Notes,
			IsCompleted: t.IsCompleted,
			CreatedAt:   t.CreatedAt,
			UpdatedAt:   t.UpdatedAt,
		})
	}

	return &GetMissionResponse{
		Mission: Mission{
			ID:          mission.ID,
			CatID:       *mission.CatID,
			Targets:     targets,
			IsCompleted: mission.IsCompleted,
			CreatedAt:   mission.CreatedAt,
			UpdatedAt:   mission.UpdatedAt,
		},
	}, nil
}

func (s Service) ListMissions(req *ListMissionsRequest) (*ListMissionsResponse, error) {
	missions, err := s.missionRepo.ListMissions()
	if err != nil {
		return nil, err
	}

	resp := make([]Mission, 0)
	for _, m := range missions {
		targets := make([]target.Target, 0)
		for _, t := range m.Targets {
			targets = append(targets, target.Target{
				ID:          t.ID,
				MissionID:   t.MissionID,
				Name:        t.Name,
				Country:     t.Country,
				Notes:       t.Notes,
				IsCompleted: t.IsCompleted,
				CreatedAt:   t.CreatedAt,
				UpdatedAt:   t.UpdatedAt,
			})
		}

		resp = append(resp, Mission{
			ID:          m.ID,
			CatID:       *m.CatID,
			Targets:     targets,
			IsCompleted: m.IsCompleted,
			CreatedAt:   m.CreatedAt,
			UpdatedAt:   m.UpdatedAt,
		})
	}

	return &ListMissionsResponse{Missions: resp}, nil
}

func (s Service) UpdateMissionCompletion(req *UpdateMissionCompletionRequest) (*UpdateMissionCompletionResponse, error) {
	if err := s.missionRepo.UpdateMissionCompletion(req.ID, req.IsCompleted); err != nil {
		return nil, err
	}

	return &UpdateMissionCompletionResponse{}, nil
}

func (s Service) DeleteMission(req *DeleteMissionRequest) (*DeleteMissionResponse, error) {
	if err := s.missionRepo.DeleteMission(req.ID); err != nil {
		return nil, err
	}

	return &DeleteMissionResponse{}, nil
}

func (s Service) AssignSpyCatToMission(req *AssignSpyCatToMissionRequest) (*AssignSpyCatToMissionResponse, error) {
	if err := s.missionRepo.AssignSpyCatToMission(req.CatID, req.MissionID); err != nil {
		return nil, err
	}

	return &AssignSpyCatToMissionResponse{}, nil
}
