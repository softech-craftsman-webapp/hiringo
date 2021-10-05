package view

import (
	"time"
)

type UserEmptyView struct {
	ID string `json:"id"`
}

type UserPublicView struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserAuthView struct {
	ID              string     `json:"id"`
	Name            string     `json:"name"`
	Email           string     `json:"email"`
	EmailVerifiedAt *time.Time `json:"email_verified_at"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	Auth            AuthView   `json:"auth,omitempty"`
}

type AuthView struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}
