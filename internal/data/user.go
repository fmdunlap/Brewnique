package data

import (
	"brewnique.fdunlap.com/internal/validator"
	_ "embed"
	"time"
)

const (
	MaxUserNameLength = 64
)

type User struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
}

func (u *User) Equal(other *User) bool {
	if u == nil && other == nil {
		return true
	}
	if u == nil || other == nil {
		return false
	}
	return u.Id == other.Id &&
		u.Email == other.Email &&
		u.Username == other.Username
}

func ValidateUser(v *validator.Validator, user User) {
	v.Check(len(user.Email) > 0, "email", "email is required")
	v.Check(validator.Matches(user.Email, validator.EmailRegex), "email", "email is invalid")
	v.Check(len(user.Username) > 0, "user_name", "user name is required")
	v.Check(len(user.Username) < 64, "user_name", "user name is too long")
}
