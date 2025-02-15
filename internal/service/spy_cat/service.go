package spy_cat

import (
	"sca/internal/repository"
)

type Service struct {
	repo repository.SpyCat
}

func NewService(repo repository.SpyCat) *Service {
	return &Service{repo: repo}
}
