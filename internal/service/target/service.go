package target

import (
	"fmt"
	"sca/internal/domain"
	"sca/internal/repository"
)

type Service struct {
	missionRepo repository.Mission
	targetsRepo repository.Target
}

func NewService(missionRepo repository.Mission, targetsRepo repository.Target) *Service {
	return &Service{missionRepo: missionRepo, targetsRepo: targetsRepo}
}

func (s Service) AddTargetsToMission(req *AddTargetsToMissionRequest) (*AddTargetsToMissionResponse, error) {
	mission, err := s.missionRepo.GetMission(req.MissionID)
	if err != nil {
		return nil, err
	}

	if mission.IsCompleted {
		return nil, fmt.Errorf("cannot update notes because the target or mission is completed")
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

	if err := s.targetsRepo.AddTargetsToMission(req.MissionID, targets); err != nil {
		return nil, err
	}

	return &AddTargetsToMissionResponse{}, nil
}

func (s Service) UpdateTargetNotes(req *UpdateTargetNotesRequest) (*UpdateTargetNotesResponse, error) {
	target, err := s.targetsRepo.GetTarget(req.ID)
	if err != nil {
		return nil, err
	}

	mission, err := s.missionRepo.GetMission(target.MissionID)
	if err != nil {
		return nil, err
	}

	if mission.IsCompleted {
		return nil, fmt.Errorf("cannot update notes because the target or mission is completed")
	}

	if err := s.targetsRepo.UpdateTargetNotes(req.ID, req.Notes); err != nil {
		return nil, err
	}

	return &UpdateTargetNotesResponse{}, nil
}

func (s Service) DeleteTarget(req *DeleteTargetRequest) (*DeleteTargetResponse, error) {
	target, err := s.targetsRepo.GetTarget(req.ID)
	if err != nil {
		return nil, err
	}

	mission, err := s.missionRepo.GetMission(target.MissionID)
	if err != nil {
		return nil, err
	}

	if mission.IsCompleted {
		return nil, fmt.Errorf("cannot delete target because the mission is completed")
	}

	if err := s.targetsRepo.DeleteTarget(req.ID); err != nil {
		return nil, err
	}

	return &DeleteTargetResponse{}, nil
}

func (s Service) UpdateTargetCompletion(req *UpdateTargetCompletionRequest) (*UpdateTargetCompletionResponse, error) {
	if err := s.targetsRepo.UpdateTargetCompletion(req.ID, req.IsCompleted); err != nil {
		return nil, err
	}

	return &UpdateTargetCompletionResponse{}, nil
}
