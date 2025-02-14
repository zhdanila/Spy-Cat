package domain

type Target struct {
	ID        int    `json:"ID"`
	MissionID int    `json:"missionID" db:"mission_id"`
	Name      string `json:"name"`
	Country   string `json:"country"`
	Notes     string `json:"notes"`
	Complete  bool   `json:"complete"`
}
