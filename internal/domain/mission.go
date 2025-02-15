package domain

import "time"

type Mission struct {
	ID          int       `db:"id"`
	CatID       *int      `db:"cat_id"`
	IsCompleted bool      `db:"is_completed"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
