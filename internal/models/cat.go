package models

type Cat struct {
	ID                int     `json:"id"`
	Name              string  `json:"name"`
	YearsOfExperience int     `json:"years_of_experience" db:"years_of_experience"`
	Breed             string  `json:"breed"`
	Salary            float64 `json:"salary"`
}

type UpdatedCat struct {
	Salary float64 `json:"salary"`
}
