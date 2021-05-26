package services

import (
	"silih_a3/entities"
	"silih_a3/inputs"
	"silih_a3/repositories"
)

type Service interface {
	GetAllDonations(userId int) ([]entities.Donation, error)
	GetDonationsById(input inputs.DonationIdInput) (entities.Donation, error)
}

type service struct {
	repository repositories.Repository
}

func NewDonationService(repository repositories.Repository) *service {
	return &service{repository}
}

func (s *service) GetAllDonations(userId int) ([]entities.Donation, error) {
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

func (s *service) GetDonationsById(input inputs.DonationIdInput) (entities.Donation, error) {
	donation, err := s.repository.FindDonationById(input.Id)
	if err != nil {
		return donation, err
	}

	return donation, nil
}
