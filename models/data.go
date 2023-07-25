package models

type Report struct {
	UserId   string               `json:"user_id"`
	Projects map[string]TimeSpent `json:"update"` // map[project_id]time_spent
}

type TimeSpent struct {
	Hours   int
	Minutes int
	Seconds int
}
