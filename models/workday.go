package models

type Workday struct {
	ID         string            `json:"_id"`
	UserId     string            `json:"user_id"`
	Date       string            `json:"workdate"`
	JoinedAt   string            `json:"joined_at"`
	WorkedOn   []TimeFrame       `json:"worked_on"`
	LastUpdate map[string]string `json:"update"` // map[project_id]time_spent
}

type TimeFrame struct {
	ProjectId        *string `json:"project_id"`
	ScreenRefUrl     string  `json:"screen_ref"` // aws presigned uri
	TimeStamp        string  `json:"timestamp"`
	MouseActivity    int     `json:"mouse_activity"`
	KeyboardActivity int     `json:"keyboard_activity"`
}
