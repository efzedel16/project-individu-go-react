package user

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
	user.Password = input.Password

	newUser, err := s.repository.InsertUser(user)
	if err != nil {
		return newUser, err
	}

	return user, nil
}