package model

import "github.com/google/uuid"

type AuthClaims struct {
	UserId   uuid.UUID `json:"userId"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
}
