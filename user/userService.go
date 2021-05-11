package user

type Service struct {

}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}