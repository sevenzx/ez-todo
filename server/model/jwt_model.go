package model

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Claims struct {
	CustomClaims
	BufferTime int64 `json:"buffer_time"`
	jwt.RegisteredClaims
}

type CustomClaims struct {
	UUID     uuid.UUID `json:"uuid"`
	Username string    `json:"username"`
	Nickname string    `json:"nickname"`
}
