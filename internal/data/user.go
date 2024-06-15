package data

import (
	"brewnique.fdunlap.com/internal/validator"
	"errors"
)

type User struct {
	ID       int64
	Email    string
	UserName string
}

func ValidateUser(v *validator.Validator, user User) {
	v.Check(len(user.Email) > 0, "email", "email is required")
	v.Check(validator.Matches(user.Email, validator.EmailRegex), "email", "email is invalid")
	v.Check(len(user.UserName) > 0, "user_name", "user name is required")
	v.Check(len(user.UserName) < 64, "user_name", "user name is too long")
}

type UserProvider interface {
	GetUser(id int64) (User, error)
	GetUserByEmail(email string) (User, error)
	GetUserByUserName(userName string) (User, error)
	ListUsers() ([]User, error)
	PutUser(user User) (User, error)
	UpdateUser(user User) (User, error)
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

func (s *UserService) CreateNewUser(email string, userName string) (User, error) {
	user := User{
		Email:    email,
		UserName: userName,
	}

	existingUser, err := s.userProvider.GetUser(user.ID)
	if err != nil {
		return existingUser, err
	}

	if existingUser.ID > 0 {
		return existingUser, errors.New("user already exists")
	}

	return s.userProvider.PutUser(user)
}

func (s *UserService) GetUser(id int64) (User, error) {
	return s.userProvider.GetUser(id)
}

func (s *UserService) GetUserByEmail(email string) (User, error) {
	return s.userProvider.GetUserByEmail(email)
}

func (s *UserService) GetUserByUserName(userName string) (User, error) {
	return s.userProvider.GetUserByUserName(userName)
}

func (s *UserService) ListUsers() ([]User, error) {
	return s.userProvider.ListUsers()
}

func (s *UserService) UpdateUser(user User) (User, error) {
	return s.userProvider.UpdateUser(user)
}

func (s *UserService) DeleteUser(id int64) error {
	return s.userProvider.DeleteUser(id)
}
