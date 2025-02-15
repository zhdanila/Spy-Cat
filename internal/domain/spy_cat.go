package domain

import "time"

type SpyCat struct {
	ID                int       `db:"id"`
	Name              string    `db:"name"`
	YearsOfExperience int       `db:"years_of_experience"`
	Breed             string    `db:"breed"`
	Salary            float64   `db:"salary"`
	CreatedAt         time.Time `db:"created_at"`
	UpdatedAt         time.Time `db:"updated_at"`
}
