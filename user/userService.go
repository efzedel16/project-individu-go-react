package user

import "golang.org/x/crypto/bcrypt"

type Service interface {
	SignUpUser(input SignUpUserInput) (User, error)
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