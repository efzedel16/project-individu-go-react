package repositories

import (
	"gorm.io/gorm"
	"silih_a3/entities"
)

type Repository interface {
	FindAllDonations() ([]entities.Donation, error)
	FindDonationsByUserId(userId int) ([]entities.Donation, error)
	FindDonationById(id int) (entities.Donation, error)
}

type repository struct {
	db *gorm.DB
}

func NewDonationRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAllDonations() ([]entities.Donation, error) {
	var donations []entities.Donation
	err := r.db.Preload("DonationImages", "donation_images.is_primary = 1").Find(&donations).Error
	if err != nil {
		return donations, err
	}

	return donations, nil
}

func (r *repository) FindDonationsByUserId(userId int) ([]entities.Donation, error) {
	var donations []entities.Donation
	err := r.db.Where("user_id = ?", userId).Preload("DonationImages", "donation_images.is_primary = 1").Find(&donations).Error
	if err != nil {
		return donations, err
	}

	return donations, nil
}

func (r *repository) FindDonationById(id int) (entities.Donation, error) {
	var donation entities.Donation
	err := r.db.Preload("User").Preload("DonationImages").Where("id = ?", id).Find(&donation).Error
	if err != nil {
		return donation, err
	}

	return donation, nil
}
