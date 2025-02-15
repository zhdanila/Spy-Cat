package target

import (
	"sca/internal/repository"
)

type Service struct {
	missionRepo repository.Mission
	targetsRepo repository.Target
}

func NewService(missionRepo repository.Mission, targetsRepo repository.Target) *Service {
	return &Service{missionRepo: missionRepo, targetsRepo: targetsRepo}
}
