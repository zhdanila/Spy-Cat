package service

import (
	"errors"
	"sca/internal/models"
	"sca/internal/repository"
	handler "sca/internal/util"
)

type CatService struct {
	repo repository.Cat
}

func NewCatService(repo repository.Cat) *CatService {
	return &CatService{repo: repo}
}

func (s *CatService) Create(cat models.Cat) (int, error) {
	//validate breed
	valid, err := handler.ValidateBreedName(cat.Breed)
	if err != nil {
		return 0, err
	}

	if !valid {
		return 0, errors.New("invalid breed name")
	}

	return s.repo.Create(cat)
}

func (s *CatService) GetById(id int) (models.Cat, error) {
	return s.repo.GetById(id)
}

func (s *CatService) GetAll() ([]models.Cat, error) {
	return s.repo.GetAll()
}

func (s *CatService) Update(id int, cat models.UpdatedCat) error {
	return s.repo.Update(id, cat)
}

func (s *CatService) Delete(id int) error {
	return s.repo.Delete(id)
}
