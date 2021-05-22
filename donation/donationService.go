package donation

type Service interface {
	GetAllDonations(userId int) ([]Donation, error)
	GetDonationsById(input DonationIdInput) (Donation, error)
}

type service struct {
	repository Repository
}

func NewDonationService(repository Repository) *service {
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

	donations, err := s.repository.FindAllDonations()
	if err != nil {
		return donations, err
	}

	return donations, nil
}

func (s *service) GetDonationsById(input DonationIdInput) (Donation, error) {
	donation, err := s.repository.FindDonationById(input.Id)
	if err != nil {
		return donation, err
	}

	return donation, nil
}
