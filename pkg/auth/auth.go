package auth

import (
	"time"

	"github.com/darchlabs/infra/pkg/users"
	"github.com/dgrijalva/jwt-go"
)

// Kind of users
const (
	KindUser           = "user"
	KindForgotPassword = "forgot-password"
	KindVerifyPassword = "verify-password"
)

// Token Struct defined to append JWT standar values to Token struct
type Token struct {
	UserID string `json:"user_id"`
	*jwt.StandardClaims
}

// Auth ...
type Auth struct {
	ID string `json:"id" db:"id"`

	UserID    string `json:"user_id" db:"user_id"`
	Token     string `json:"token" db:"token"`
	Blacklist bool   `json:"blacklist" db:"blacklist"`
	Kind      string `json:"kind" db:"kind"` // user, forgot-password

	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"-" db:"deleted_at"`
}

// Response ...
type Response struct {
	Data *users.User `json:"data"`
	Meta *MetaToken  `json:"meta"`
}

// MetaToken ...
type MetaToken struct {
	Token             string `json:"token"`
	VerificationToken string `json:"verification_token,omitempty"`
}

// Service ...
type Service interface {
	GetByToken(token string) (*Auth, error)
	Login(email, password string) (*Response, error)
	Signup(user *users.User) (*Response, error)
	VerifyToken(token string, kind string) (*Token, error)
	VerifyEmail(token string) error
	Logout(token string) error
	ForgotPassword(email string) (string, error)
	RecoverPassword(newPassword, token string) error
}

// Query ...
type Query struct {
	Token  string
	Email  string
	UserID string
}
