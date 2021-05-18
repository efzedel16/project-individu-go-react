package donation

type Service interface {
	GetDonations(userId int) ([]Donation, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetDonations(userId int) ([]Donation, error) {
	if userId != 0 {
		donations, err := s.repository.FindDonationsByUserId(userId)
		if err != nil {
			return donations, err
		}

		return donations, nil
	}

	donations, err := s.repository.FindAllDonations()
	if err != nil {
		return donations, err
	}

	return donations, nil
}
