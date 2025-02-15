package target

import (
	"sca/internal/domain"
	"sca/internal/repository"
)

type Service struct {
	missionRepo repository.Mission
	targetsRepo repository.Target
}

func (s Service) AddTargetsToMission(req *AddTargetsToMissionRequest) (*AddTargetsToMissionResponse, error) {
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

func (s Service) DeleteTarget(req *DeleteTargetRequest) (*DeleteTargetResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) UpdateTargetCompletion(req *UpdateTargetCompletionRequest) (*UpdateTargetCompletionResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) UpdateTargetNotes(req *UpdateTargetNotesRequest) (*UpdateTargetNotesResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewService(missionRepo repository.Mission, targetsRepo repository.Target) *Service {
	return &Service{missionRepo: missionRepo, targetsRepo: targetsRepo}
}
