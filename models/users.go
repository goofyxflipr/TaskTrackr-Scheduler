package models

import "time"

type User struct {
	ID                string               `json:"_id"`
	FirstName         *string              `json:"first_name" validate:"required,min=1,max=100"`
	LastName          *string              `json:"last_name" validate:"required,min=1,max=100"`
	Password          *string              `json:"password" validate:"required,min=8"`
	Email             *string              `json:"email" validate:"email,required"`
	ProfileImageUrl   *string              `json:"profile_image_url"`
	Phone             *string              `json:"phone"`
	Organization      *Organization        `json:"organization"`
	Projects          map[string]Project   `json:"projects"`
	Hackathons        map[string]Hackathon `json:"hackathons"`
	Tags              []string             `json:"tags"`
	MonitorInterval   int                  `json:"monitor_interval"`
	UploadInterval    int                  `json:"upload_interval"`
	AccessToken       *string              `json:"access"`
	RefreshToken      *string              `json:"refresh"`
	CreatedAt         time.Time            `json:"created_at"`
	UpdatedAt         time.Time            `json:"updated_at"`
	LastSignInAt      time.Time            `json:"last_signin"`
	UserType          string               `json:"user_type"`
	ConfirmationToken string               `json:"confirmation_token"`
	IsVerified        bool                 `json:"is_verified"`
	Banned            bool                 `json:"banned"`
}
