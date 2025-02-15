package target

import "time"

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

type AddTargetsToMissionRequest struct {
	MissionID int      `json:"mission_id" validate:"required"`
	Targets   []Target `json:"targets" validate:"required,dive"`
}

type AddTargetsToMissionResponse struct{}

type DeleteTargetRequest struct {
	ID int `json:"id"`
}

type DeleteTargetResponse struct{}

type UpdateTargetCompletionRequest struct {
	ID          int  `json:"id"`
	IsCompleted bool `json:"is_completed"`
}

type UpdateTargetCompletionResponse struct{}

type UpdateTargetNotesRequest struct {
	ID    int    `json:"id"`
	Notes string `json:"notes"`
}

type UpdateTargetNotesResponse struct{}
