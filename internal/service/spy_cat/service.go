package spy_cat

import (
	"sca/internal/repository"
)

type Service struct {
	repo repository.SpyCat
}

func (s Service) CreateSpyCat(req *CreateSpyCatRequest) (*CreateSpyCatResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) GetSpyCat(req *GetSpyCatRequest) (*GetSpyCatResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) ListSpyCats(req *ListSpyCatsRequest) (*ListSpyCatsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) UpdateSpyCatSalary(req *UpdateSpyCatSalaryRequest) (*UpdateSpyCatSalaryResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) DeleteSpyCat(req *DeleteSpyCatRequest) (*DeleteSpyCatResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewService(repo repository.SpyCat) *Service {
	return &Service{repo: repo}
}
