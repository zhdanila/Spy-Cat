package service

import (
	"sca/internal/repository"
	"sca/internal/service/mission"
	"sca/internal/service/spy_cat"
	"sca/internal/service/target"
)

type Service struct {
	SpyCat
	Mission
	Target
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		SpyCat:  spy_cat.NewService(repo.SpyCat),
		Mission: mission.NewService(repo.Mission, repo.Target),
		Target:  mission.NewService(repo.Mission, repo.Target),
	}
}

type SpyCat interface {
	CreateSpyCat(req *spy_cat.CreateSpyCatRequest) (*spy_cat.CreateSpyCatResponse, error)
	GetSpyCat(req *spy_cat.GetSpyCatRequest) (*spy_cat.GetSpyCatResponse, error)
	ListSpyCats(req *spy_cat.ListSpyCatsRequest) (*spy_cat.ListSpyCatsResponse, error)
	UpdateSpyCatSalary(req *spy_cat.UpdateSpyCatSalaryRequest) (*spy_cat.UpdateSpyCatSalaryResponse, error)
	DeleteSpyCat(req *spy_cat.DeleteSpyCatRequest) (*spy_cat.DeleteSpyCatResponse, error)
}

type Mission interface {
	CreateMission(req *mission.CreateMissionRequest) (*mission.CreateMissionResponse, error)
	GetMission(req *mission.GetMissionRequest) (*mission.GetMissionResponse, error)
	ListMissions(req *mission.ListMissionsRequest) (*mission.ListMissionsResponse, error)
	UpdateMissionCompletion(req *mission.UpdateMissionCompletionRequest) (*mission.UpdateMissionCompletionResponse, error)
	DeleteMission(req *mission.DeleteMissionRequest) (*mission.DeleteMissionResponse, error)
	AssignSpyCatToMission(req *mission.AssignSpyCatToMissionRequest) (*mission.AssignSpyCatToMissionResponse, error)
}

type Target interface {
	AddTargetsToMission(req *target.AddTargetsToMissionRequest) (*target.AddTargetsToMissionResponse, error)
	DeleteTarget(req *target.DeleteTargetRequest) (*target.DeleteTargetResponse, error)
	UpdateTargetCompletion(req *target.UpdateTargetCompletionRequest) (*target.UpdateTargetCompletionResponse, error)
	UpdateTargetNotes(req *target.UpdateTargetNotesRequest) (*target.UpdateTargetNotesResponse, error)
}
