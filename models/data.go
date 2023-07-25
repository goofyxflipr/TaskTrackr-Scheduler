package models

type UserReport struct {
	UserId    string             `json:"_id"`
	UserEmail string             `json:"email"`
	Name      string             `json:"name"`
	Projects  map[string]Details `json:"work_update"` // map[project_id]time_spent
}

type Details struct {
	Name      string    `json:"project_name"`
	TimeSpent TimeSpent `json:"time_spent"`
}

type TimeSpent struct {
	Hours   int
	Minutes int
	Seconds int
}
