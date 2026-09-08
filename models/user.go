package models

import "time"

type User struct {
	ID            int       `json:"id"`
	Email         string    `json:"email"`
	PasswordHash  string    `json:"-"`
	EmailVerified bool      `json:"email_verification"`
	CreatedAt     time.Time `json:"created_at"`
}
