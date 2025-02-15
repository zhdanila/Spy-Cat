package mission

import (
	"sca/internal/repository"
	"sca/internal/service/target"
)

type Service struct {
	missionRepo repository.Mission
	targetsRepo repository.Target
}

func (s Service) AddTargetsToMission(req *target.AddTargetsToMissionRequest) (*target.AddTargetsToMissionResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) DeleteTarget(req *target.DeleteTargetRequest) (*target.DeleteTargetResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) UpdateTargetCompletion(req *target.UpdateTargetCompletionRequest) (*target.UpdateTargetCompletionResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) UpdateTargetNotes(req *target.UpdateTargetNotesRequest) (*target.UpdateTargetNotesResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) CreateMission(req *CreateMissionRequest) (*CreateMissionResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) GetMission(req *GetMissionRequest) (*GetMissionResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) ListMissions(req *ListMissionsRequest) (*ListMissionsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) UpdateMissionCompletion(req *UpdateMissionCompletionRequest) (*UpdateMissionCompletionResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) DeleteMission(req *DeleteMissionRequest) (*DeleteMissionResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) AssignSpyCatToMission(req *AssignSpyCatToMissionRequest) (*AssignSpyCatToMissionResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewService(missionRepo repository.Mission, targetsRepo repository.Target) *Service {
	return &Service{missionRepo: missionRepo, targetsRepo: targetsRepo}
}
