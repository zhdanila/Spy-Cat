package models

type Mission struct {
	ID       int      `json:"id" db:"id"`
	CatID    int      `json:"catID" db:"cat_id"`
	Complete bool     `json:"complete" db:"complete"`
	Targets  []Target `json:"targets" db:"-"`
}

type UpdatedMission struct {
	Complete bool     `json:"complete"`
	Notes    string   `json:"notes"`
	CatID    int      `json:"catID"`
	Targets  []Target `json:"targets" db:"-"`
}
