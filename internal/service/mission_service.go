package service

import (
	"errors"
	"sca/internal/domain"
	"sca/internal/repository"
)

type MissionService struct {
	missionRepo repository.Mission
	targetsRepo repository.Target
}

func NewMissionService(missionRepo repository.Mission, targetsRepo repository.Target) *MissionService {
	return &MissionService{missionRepo: missionRepo, targetsRepo: targetsRepo}
}

func (s *MissionService) Create(missionToAdd domain.Mission) (int, error) {
	//check one cat can only have one mission at a time
	missions, err := s.missionRepo.GetAll()
	if err != nil {
		return 0, err
	}

	for _, mission := range missions {
		if missionToAdd.CatID == mission.CatID {
			if !mission.Complete {
				return 0, errors.New("one cat can only have one mission at a time")
			}
		}

	}

	//a mission assumes a range of targets (minimum: 1, maximum: 3)
	if len(missionToAdd.Targets) < 1 || len(missionToAdd.Targets) > 3 {
		return 0, errors.New("a mission assumes a range of targets (minimum: 1, maximum: 3)")
	}

	return s.missionRepo.Create(missionToAdd)
}

func (s *MissionService) Update(id int, updatedMission domain.UpdatedMission) error {
	//check if either the target or the mission is completed
	if updatedMission.Notes != "" {
		mission, err := s.missionRepo.GetByID(id)
		if err != nil {
			return err
		}

		if mission.Complete {
			return errors.New("notes cannot be updated because the mission is completed\n")
		}

		targets, err := s.targetsRepo.GetAll()
		if err != nil {
			return err
		}

		for _, target := range targets {
			if target.MissionID == id && target.Complete {
				return errors.New("notes cannot be updated because the target is completed\n")
			}
		}
	}

	//update mission
	err := s.missionRepo.Update(id, updatedMission)
	if err != nil {
		return err
	}

	//check completing all the targets
	targets, err := s.targetsRepo.GetAll()
	if err != nil {
		return err
	}

	var state bool = false
	for _, target := range targets {
		if target.MissionID == id {
			if target.Complete {
				state = true
			} else {
				state = false
				break
			}
		}
	}

	if state {
		err := s.missionRepo.Update(id, domain.UpdatedMission{Complete: true})
		if err != nil {
			return err
		}
	}

	return err
}

func (s *MissionService) Delete(id int) error {
	//check if it is already assigned to a cat
	mission, err := s.missionRepo.GetByID(id)
	if err != nil {
		return err
	}

	if mission.CatID == 0 {
		return errors.New("a mission cannot be deleted because it is already assigned to a cat\n")
	}

	return s.missionRepo.Delete(id)
}

func (s *MissionService) GetByID(id int) (domain.Mission, error) {
	return s.missionRepo.GetByID(id)
}

func (s *MissionService) GetAll() ([]domain.Mission, error) {
	return s.missionRepo.GetAll()
}

func (s *MissionService) DeleteTarget(missionId, targetId int) error {
	//check if target is already completed
	target, err := s.targetsRepo.GetById(targetId)
	if err != nil {
		return err
	}

	if target.Complete {
		return errors.New("a target cannot be deleted because it is already completed")
	}

	return s.missionRepo.DeleteTarget(missionId, targetId)
}

func (s *MissionService) CreateTarget(missionId int, target domain.Target) (int, error) {
	//check if the mission is already completed
	mission, err := s.missionRepo.GetByID(missionId)
	if err != nil {
		return 0, err
	}

	if mission.Complete {
		return 0, errors.New("a target cannot be added because the mission is already completed\n")
	}

	return s.missionRepo.CreateTarget(missionId, target)
}
