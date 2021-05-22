package donation

import "gorm.io/gorm"

type Repository interface {
	FindAllDonations() ([]Donation, error)
	FindDonationsByUserId(userId int) ([]Donation, error)
	FindDonationById(id int) (Donation, error)
}

type repository struct {
	db *gorm.DB
}

func NewDonationRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAllDonations() ([]Donation, error) {
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

func (r *repository) FindDonationById(id int) (Donation, error) {
	var donation Donation
	err := r.db.Preload("User").Preload("DonationImages").Where("id = ?", id).Find(&donation).Error
	if err != nil {
		return donation, err
	}

	return donation, nil
}
