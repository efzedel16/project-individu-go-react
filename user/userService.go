package user

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	SignUpUser(input SignUpUserInput) (User, error)
	SignInUser(input SigInInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	SaveAvatar(id int, FileLocation string) (User, error)
	GetUserById(id int) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) SignUpUser(input SignUpUserInput) (User, error) {
	user := User{}
	user.FirstName = input.FirstName
	user.LastName = input.LastName
	user.Email = input.Email
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(passwordHash)
	user.Role = "admin"

	newUser, err := s.repository.InsertUser(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) SignInUser(input SigInInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.Id == 0 {
		return user, errors.New("no user found on that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	email := input.Email
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}
	if user.Id == 0 {
		return true, nil
	}
	return false, nil
}

func (s *service) SaveAvatar(id int, FileLocation string) (User, error) {
	user, err := s.repository.FindById(id)
	if err != nil {
		return user, err
	}
	user.Avatar = FileLocation

	updatedUser, err := s.repository.UpdateUser(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (s *service) GetUserById(id int) (User, error) {
	user, err := s.repository.FindById(id)
	if err != nil {
		return user, err
	}

	if user.Id == 0 {
		return user, errors.New("no user found on with email")
	}

	return user, nil
}