package donation

type Service interface {
	GetAllDonations(userId int) ([]Donation, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllDonations(userId int) ([]Donation, error) {
	if userId != 0 {
		donations, err := s.repository.FindDonationsByUserId(userId)
		if err != nil {
			return donations, err
		}

		return donations, nil
	}

	donations, err := s.repository.FindDonations()
	if err != nil {
		return donations, err
	}

	return donations, nil
}
