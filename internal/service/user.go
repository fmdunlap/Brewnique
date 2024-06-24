package service

import (
	"brewnique.fdunlap.com/internal/data"
	"brewnique.fdunlap.com/internal/validator"
	"fmt"
)

type UserProvider interface {
	GetUser(id int64) (*data.User, error)
	GetUserByEmail(email string) (*data.User, error)
	GetUserByUsername(userName string) (*data.User, error)
	ListUsers() ([]*data.User, error)
	PutUser(user *data.User) (*data.User, error)
	UpdateUser(user *data.User) (*data.User, error)
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

func (s *UserService) CreateNewUser(email string, username string) (*data.User, error) {
	if len(username) > data.MaxUserNameLength {
		return nil, fmt.Errorf("username is too long")
	}
	if !validator.Matches(email, validator.EmailRegex) {
		return nil, fmt.Errorf("email is invalid")
	}

	if len(username) == 0 {
		for {
			username = data.GenerateUsername()
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

	user := data.User{
		Email:    email,
		Username: username,
	}

	return s.userProvider.PutUser(&user)
}

func (s *UserService) GetUser(id int64) (*data.User, error) {
	if id == 0 {
		return nil, fmt.Errorf("id cannot be 0")
	}

	return s.userProvider.GetUser(id)
}

func (s *UserService) GetUserByEmail(email string) (*data.User, error) {
	if len(email) == 0 {
		return nil, fmt.Errorf("email cannot be empty")
	}
	if !validator.Matches(email, validator.EmailRegex) {
		return nil, fmt.Errorf("email is invalid")
	}

	return s.userProvider.GetUserByEmail(email)
}

func (s *UserService) GetUserByUsername(userName string) (*data.User, error) {
	return s.userProvider.GetUserByUsername(userName)
}

func (s *UserService) ListUsers() ([]*data.User, error) {
	// TODO: pagination
	return s.userProvider.ListUsers()
}

func (s *UserService) UpdateUser(user *data.User) (*data.User, error) {
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
	if len(user.Username) > data.MaxUserNameLength {
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
