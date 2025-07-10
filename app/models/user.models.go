package models

import "time"

type User struct {
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Email        string    `json:"email"`
	Mobile       string    `json:"mobile"`
	CountryCode  string    `json:"country_code"`
	OTP          string    `json:"otp"`
	ExpiredAt    int64     `json:"expired_at"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	ProfileImage string    `json:"profile_image"`
}

type InputUser struct {
	FirstName    string `json:"first_name" validate:"required"`
	LastName     string `json:"last_name" validate:"required"`
	Email        string `json:"email" validate:"required"`
	ProfileImage string `json:"profile_image" validate:"required"`
}

type LoginInput struct {
	Mobile      string `json:"mobile" validate:"required"`
	CountryCode string `json:"country_code" validate:"required"`
}
