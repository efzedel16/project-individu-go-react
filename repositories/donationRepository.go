package repositories

import (
	"gorm.io/gorm"
	"silih_a3/entities"
)

type DonationRepository interface {
	FindAllDonations() ([]entities.Donation, error)
	FindDonationsByUserId(userId int) ([]entities.Donation, error)
	FindDonationById(donationId int) (entities.Donation, error)
}

type donationRepository struct {
	donationDb *gorm.DB
}

func NewDonationRepository(donationDb *gorm.DB) *donationRepository {
	return &donationRepository{donationDb}
}

func (r *donationRepository) FindAllDonations() ([]entities.Donation, error) {
	var donationsData []entities.Donation
	err := r.donationDb.Preload("DonationImages", "donation_images.is_primary = 1").Find(&donationsData).Error
	if err != nil {
		return donationsData, err
	}

	return donationsData, nil
}

func (r *donationRepository) FindDonationsByUserId(userId int) ([]entities.Donation, error) {
	var donationsData []entities.Donation
	err := r.donationDb.Where("user_id = ?", userId).Preload("DonationImages", "donation_images.is_primary = 1").Find(&donationsData).Error
	if err != nil {
		return donationsData, err
	}

	return donationsData, nil
}

func (r *donationRepository) FindDonationById(donationId int) (entities.Donation, error) {
	var donationsData entities.Donation
	err := r.donationDb.Preload("User").Preload("DonationImages").Where("donation_id = ?", donationId).Find(&donationsData).Error
	if err != nil {
		return donationsData, err
	}

	return donationsData, nil
}
