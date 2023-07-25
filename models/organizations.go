package models

import "time"

type Organization struct {
	ID              string             `json:"_id"`
	Name            *string            `json:"name"`
	Projects        map[string]Project `json:"projects"` // this field is different for all users.
	MonitorInterval int                `json:"monitor_interval"`
	UploadInterval  int                `json:"upload_interval"`
	CreatedAt       time.Time          `json:"created_at"`
	UpdatedAt       time.Time          `json:"updated_at"`
}
