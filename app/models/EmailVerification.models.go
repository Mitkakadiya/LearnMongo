package models

import (
	"crypto/rand"
	"encoding/hex"
)

type EmailVerification struct {
	Email                  string `bson:"email"`
	IsVerified             bool   `bson:"is_verified" default:"false"`
	EmailVerificationToken string `bson:"email_verification_token"`
}

type InputEmail struct {
	Email string `bson:"email" validate:"required"`
}

func GenerateToken(n int) (string, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
