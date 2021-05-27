package services

import (
	"silih_a3/entities"
	"silih_a3/inputs"
	"silih_a3/repositories"
)

type DonationService interface {
	GetAllDonations(userId int) ([]entities.Donation, error)
	GetDonationsById(idInput inputs.DonationIdInput) (entities.Donation, error)
}

type donationService struct {
	donationRepository repositories.DonationRepository
}

func NewDonationService(donationRepository repositories.DonationRepository) *donationService {
	return &donationService{donationRepository}
}

func (s *donationService) GetAllDonations(userId int) ([]entities.Donation, error) {
	if userId != 0 {
		donationsData, err := s.donationRepository.FindDonationsByUserId(userId)
		if err != nil {
			return donationsData, err
		}

		return donationsData, nil
	}

	donationsData, err := s.donationRepository.FindAllDonations()
	if err != nil {
		return donationsData, err
	}

	return donationsData, nil
}

func (s *donationService) GetDonationsById(input inputs.DonationIdInput) (entities.Donation, error) {
	donationData, err := s.donationRepository.FindDonationById(input.Id)
	if err != nil {
		return donationData, err
	}

	return donationData, nil
}
