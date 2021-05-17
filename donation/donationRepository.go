package donation

import "gorm.io/gorm"

type Repository interface {
	FindAllDonations() ([]Donation, error)
	FindDonationsByUserId(userId int) ([]Donation, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAllDonations() []Donation {
	var donations []Donation
	err := r.db.Preload("DonationImages", "donation_images.is_primary = 1").Find(&donations).Error
	if err != nil {
		return donations
	}

	return donations
}

func (r *repository) FindDonationsByUserId(userId int) []Donation {
	var donations []Donation
	err := r.db.Where("user_id = ?", userId).Preload("DonationImages", "donation_images.is_primary = 1").Find(&donations).Error
	if err != nil {
		return donations
	}

	return donations
}
