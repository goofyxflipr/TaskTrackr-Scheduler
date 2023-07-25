package models

import "time"

type Project struct {
	ID              string    `json:"_id"`
	ProjectName     *string   `json:"name"`
	RepositoryURI   *string   `json:"uri"`
	Tasks           []Task    `json:"tasks"`
	Tags            []string  `json:"tags"`
	MonitorInterval int       `json:"monitor_interval"` // Change to random amount of time
	UploadInterval  int       `json:"upload_interval"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type TaskDetails struct {
	Topic       *string `json:"topic"`
	Description *string `json:"description"`
}

type Task struct {
	ID          string      `json:"_id"`
	TaskDetails TaskDetails `json:"details"`
	Status      *string     `json:"status"`
	Priority    int         `json:"priority"`
	Deadline    time.Time   `json:"deadline"`
}
