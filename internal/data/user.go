package data

import (
	"brewnique.fdunlap.com/internal/validator"
	_ "embed"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

//go:embed embed/adjective_list.txt
var adjectiveList string

//go:embed embed/noun_list.txt
var nounList string

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

type UserProvider interface {
	GetUser(id int64) (*User, error)
	GetUserByEmail(email string) (*User, error)
	GetUserByUsername(userName string) (*User, error)
	ListUsers() ([]*User, error)
	PutUser(user *User) (*User, error)
	UpdateUser(user *User) (*User, error)
	DeleteUser(id int64) error
}

type UserService struct {
	userProvider UserProvider
}

func NewUserService(userProvider UserProvider) *UserService {
	return &UserService{
		userProvider: userProvider,
	}
}

func (s *UserService) CreateNewUser(email string, username string) (*User, error) {
	if len(username) > MaxUserNameLength {
		return nil, fmt.Errorf("username is too long")
	}
	if !validator.Matches(email, validator.EmailRegex) {
		return nil, fmt.Errorf("email is invalid")
	}

	if len(username) == 0 {
		for {
			username = s.GenerateUserName()
			if _, err := s.userProvider.GetUserByUsername(username); err != nil {
				break
			}
		}
	}

	if _, err := s.userProvider.GetUserByEmail(email); err == nil {
		return nil, fmt.Errorf("user with email %s already exists", email)
	}
	if _, err := s.userProvider.GetUserByUsername(username); err == nil {
		return nil, fmt.Errorf("user with username %s already exists", username)
	}

	//TODO should probably check for profanity?

	user := User{
		Email:    email,
		Username: username,
	}

	return s.userProvider.PutUser(&user)
}

func (s *UserService) GetUser(id int64) (*User, error) {
	if id == 0 {
		return nil, fmt.Errorf("id cannot be 0")
	}

	return s.userProvider.GetUser(id)
}

func (s *UserService) GetUserByEmail(email string) (*User, error) {
	if len(email) == 0 {
		return nil, fmt.Errorf("email cannot be empty")
	}
	if !validator.Matches(email, validator.EmailRegex) {
		return nil, fmt.Errorf("email is invalid")
	}

	return s.userProvider.GetUserByEmail(email)
}

func (s *UserService) GetUserByUsername(userName string) (*User, error) {
	return s.userProvider.GetUserByUsername(userName)
}

func (s *UserService) ListUsers() ([]*User, error) {
	// TODO: pagination
	return s.userProvider.ListUsers()
}

func (s *UserService) UpdateUser(user *User) (*User, error) {
	if user.Id == 0 {
		return nil, fmt.Errorf("id cannot be 0")
	}
	if len(user.Email) == 0 {
		return nil, fmt.Errorf("email cannot be empty")
	}
	if !validator.Matches(user.Email, validator.EmailRegex) {
		return nil, fmt.Errorf("email is invalid")
	}
	if len(user.Username) == 0 {
		return nil, fmt.Errorf("username cannot be empty")
	}
	if len(user.Username) > MaxUserNameLength {
		return nil, fmt.Errorf("username is too long")
	}

	return s.userProvider.UpdateUser(user)
}

func (s *UserService) DeleteUser(id int64) error {
	if id == 0 {
		return fmt.Errorf("id cannot be 0")
	}
	return s.userProvider.DeleteUser(id)
}

func (s *UserService) GenerateUserName() string {
	adjectives := strings.Split(adjectiveList, "\n")
	nouns := strings.Split(nounList, "\n")

	randomAdjective := adjectives[rand.Intn(len(adjectives))]
	randomNoun := nouns[rand.Intn(len(nouns))]

	return strings.Title(randomAdjective) + strings.Title(randomNoun)
}
