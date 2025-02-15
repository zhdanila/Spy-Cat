package spy_cat

import "time"

type SpyCat struct {
	ID                int       `json:"id"`
	Name              string    `json:"name"`
	YearsOfExperience int       `json:"years_of_experience"`
	Breed             string    `json:"breed"`
	Salary            float64   `json:"salary"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type CreateSpyCatRequest struct {
	Name              string  `json:"name" validate:"required,min=1,max=100"`
	YearsOfExperience int     `json:"years_of_experience" validate:"required,min=0"`
	Breed             string  `json:"breed" validate:"required,min=1,max=50"`
	Salary            float64 `json:"salary" validate:"required,min=0"`
}

type CreateSpyCatResponse struct {
	ID int `json:"id"`
}

type GetSpyCatRequest struct {
	ID int `json:"id"`
}

type GetSpyCatResponse struct {
	SpyCat SpyCat `json:"spy_cat"`
}

type UpdateSpyCatSalaryRequest struct {
	ID     int     `json:"id"`
	Salary float64 `json:"salary" validate:"required,min=0"`
}

type UpdateSpyCatSalaryResponse struct{}

type DeleteSpyCatRequest struct {
	ID int `json:"id"`
}

type DeleteSpyCatResponse struct{}

type ListSpyCatsRequest struct{}

type ListSpyCatsResponse struct {
	SpyCats []SpyCat `json:"spy_cats"`
}
