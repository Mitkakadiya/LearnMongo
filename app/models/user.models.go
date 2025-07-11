package models

import "time"

type User struct {
	FirstName    string    `bson:"first_name"`
	LastName     string    `bson:"last_name"`
	Email        string    `bson:"email"`
	Mobile       string    `bson:"mobile"`
	CountryCode  string    `bson:"country_code"`
	OTP          string    `bson:"otp"`
	ExpiredAt    int64     `bson:"expired_at"`
	CreatedAt    time.Time `bson:"created_at"`
	UpdatedAt    time.Time `bson:"updated_at"`
	ProfileImage string    `bson:"profile_image"`
}

type InputUser struct {
	FirstName    string `bson:"first_name" validate:"required"`
	LastName     string `bson:"last_name" validate:"required"`
	Email        string `bson:"email" validate:"required"`
	ProfileImage string `bson:"profile_image" validate:"required"`
}

type LoginInput struct {
	Mobile      string `bson:"mobile" validate:"required"`
	CountryCode string `bson:"country_code" validate:"required"`
}
