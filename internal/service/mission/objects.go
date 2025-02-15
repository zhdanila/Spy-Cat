package mission

import "time"

type Mission struct {
	ID          int       `json:"id"`
	CatID       int       `json:"cat_id"`
	Targets     []Target  `json:"targets"`
	IsCompleted bool      `json:"is_completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Target struct {
	ID          int       `json:"id"`
	MissionID   int       `json:"mission_id"`
	Name        string    `json:"name"`
	Country     string    `json:"country"`
	Notes       *string   `json:"notes"`
	IsCompleted bool      `json:"is_completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateMissionRequest struct {
	CatID   int      `json:"cat_id" validate:"required"`
	Targets []Target `json:"targets" validate:"required,dive"`
}

type CreateMissionResponse struct {
	ID int `json:"id"`
}

type GetMissionRequest struct {
	ID int `json:"id"`
}

type GetMissionResponse struct {
	Mission Mission `json:"mission"`
}

type UpdateMissionCompletionRequest struct {
	ID          int  `json:"id"`
	IsCompleted bool `json:"is_completed"`
}

type UpdateMissionCompletionResponse struct{}

type DeleteMissionRequest struct {
	ID int `json:"id"`
}

type DeleteMissionResponse struct{}

type ListMissionsRequest struct{}

type ListMissionsResponse struct {
	Missions []Mission `json:"missions"`
}

type AssignSpyCatToMissionRequest struct {
	CatID     int `json:"cat_id"`
	MissionID int `json:"mission_id"`
}

type AssignSpyCatToMissionResponse struct{}
