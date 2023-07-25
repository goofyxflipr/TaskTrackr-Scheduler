package models

import "time"

type Hackathon struct {
	ID                     string       `json:"_id"`
	Name                   *string      `json:"name"`
	Description            *string      `json:"description"`
	BannerUrl              *string      `json:"banner_url"`
	ConductingOrganization Organization `json:"organization"`
	Template               Project      `json:"project"`
	StartTime              time.Time    `json:"start_time"`
	EndTime                time.Time    `json:"end_time"`
	CreatedAt              time.Time    `json:"created_at"`
	UpdatedAt              time.Time    `json:"updated_at"`
}
