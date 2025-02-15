package mission

import (
	"sca/internal/service/target"
	"time"
)

type Mission struct {
	ID          int             `json:"id"`
	CatID       int             `json:"cat_id"`
	Targets     []target.Target `json:"targets"`
	IsCompleted bool            `json:"is_completed"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

type CreateMissionRequest struct {
	CatID   int             `json:"cat_id" validate:"required"`
	Targets []target.Target `json:"targets" validate:"required,dive"`
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
