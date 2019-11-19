package v1

import (
	"github.com/dgrijalva/jwt-go"
)

// LoginRequest Scheme From Request
type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

// LoginDetailResponse scheme to detail respons
type LoginDetailResponse struct {
	IsRegistered bool             `json:"is_registered"`
	Token        string           `json:"authorization"`
	Expired      int64            `json:"expired"`
	UserData     SessionDataUsers `json:"user_data"`
}

// LoginOauthDetailResponse mapping
type LoginOauthDetailResponse struct {
	IsRegistered bool               `json:"is_registered"`
	Token        string             `json:"authorization"`
	Expired      int64              `json:"expired"`
	UserData     UserDetailResponse `json:"user_data"`
	SessionData  SessionDataUsers   `json:"session_data"`
}

// SessionDataOrganizers types mapping
type SessionDataOrganizers struct {
	ID     uint   `json:"id"`
	UserID uint   `json:"users_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

// SessionDataUsers types
type SessionDataUsers struct {
	ID           uint                  `json:"id"`
	Name         string                `json:"name"`
	Email        string                `json:"email"`
	IsOrganizers bool                  `json:"is_organizers"`
	Image        string                `json:"image"`
	Organizers   SessionDataOrganizers `json:"organizers"`
}

// Claims mapping
type Claims struct {
	SessionData string `json:"session"`
	jwt.StandardClaims
}
