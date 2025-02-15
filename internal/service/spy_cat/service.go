package spy_cat

import (
	"sca/internal/domain"
	"sca/internal/repository"
)

type Service struct {
	repo repository.SpyCat
}

func NewService(repo repository.SpyCat) *Service {
	return &Service{repo: repo}
}

func (s Service) CreateSpyCat(req *CreateSpyCatRequest) (*CreateSpyCatResponse, error) {
	spyCat := &domain.SpyCat{
		Name:              req.Name,
		YearsOfExperience: req.YearsOfExperience,
		Breed:             req.Breed,
		Salary:            req.Salary,
	}

	if err := s.repo.CreateSpyCat(spyCat); err != nil {
		return nil, err
	}

	return &CreateSpyCatResponse{ID: spyCat.ID}, nil
}

func (s Service) GetSpyCat(req *GetSpyCatRequest) (*GetSpyCatResponse, error) {
	spyCat, err := s.repo.GetSpyCat(req.ID)
	if err != nil {
		return nil, err
	}

	return &GetSpyCatResponse{
		SpyCat: SpyCat{
			ID:                spyCat.ID,
			Name:              spyCat.Name,
			YearsOfExperience: spyCat.YearsOfExperience,
			Breed:             spyCat.Breed,
			Salary:            spyCat.Salary,
		},
	}, nil
}

func (s Service) ListSpyCats(req *ListSpyCatsRequest) (*ListSpyCatsResponse, error) {
	spyCats, err := s.repo.ListSpyCats()
	if err != nil {
		return nil, err
	}

	var resp ListSpyCatsResponse
	for _, spyCat := range spyCats {
		resp.SpyCats = append(resp.SpyCats, SpyCat{
			ID:                spyCat.ID,
			Name:              spyCat.Name,
			YearsOfExperience: spyCat.YearsOfExperience,
			Breed:             spyCat.Breed,
			Salary:            spyCat.Salary,
		})
	}

	return &resp, nil
}

func (s Service) UpdateSpyCatSalary(req *UpdateSpyCatSalaryRequest) (*UpdateSpyCatSalaryResponse, error) {
	if err := s.repo.UpdateSpyCatSalary(req.ID, req.Salary); err != nil {
		return nil, err
	}

	return &UpdateSpyCatSalaryResponse{}, nil
}

func (s Service) DeleteSpyCat(req *DeleteSpyCatRequest) (*DeleteSpyCatResponse, error) {
	if err := s.repo.DeleteSpyCat(req.ID); err != nil {
		return nil, err
	}

	return &DeleteSpyCatResponse{}, nil
}
