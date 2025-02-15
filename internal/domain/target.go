package domain

import "time"

type Target struct {
	ID          int       `db:"id"`
	MissionID   int       `db:"mission_id"`
	Name        string    `db:"name"`
	Country     string    `db:"country"`
	Notes       *string   `db:"notes"` // Nullable
	IsCompleted bool      `db:"is_completed"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
