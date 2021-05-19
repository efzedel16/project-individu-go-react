package donation

import "gorm.io/gorm"

type Repository interface {
	FindDonations() ([]Donation, error)
	FindDonationsByUserId(userId int) ([]Donation, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindDonations() ([]Donation, error) {
	var donations []Donation
	err := r.db.Preload("DonationImages", "donation_images.is_primary = 1").Find(&donations).Error
	if err != nil {
		return donations, err
	}

	return donations, nil
}

func (r *repository) FindDonationsByUserId(userId int) ([]Donation, error) {
	var donations []Donation
	err := r.db.Where("user_id = ?", userId).Preload("DonationImages", "donation_images.is_primary = 1").Find(&donations).Error
	if err != nil {
		return donations, err
	}

	return donations, nil
}
