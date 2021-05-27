package repositories

import (
	"gorm.io/gorm"
	"silih_a3/entities"
)

type UserRepository interface {
	InsertUser(userData entities.User) (entities.User, error)
	UpdateUser(userData entities.User) (entities.User, error)
	UpdateUserById(userId int, updateData map[string]interface{}) (entities.User, error)
	DeleteUserById(userId int) (string, error)
	FindAllUsers() ([]entities.User, error)
	FindUserById(userId int) (entities.User, error)
	FindUserByEmail(userEmail string) (entities.User, error)
}

type userRepository struct {
	userDb *gorm.DB
}

func NewUserRepository(userDb *gorm.DB) *userRepository {
	return &userRepository{userDb}
}

func (r *userRepository) InsertUser(userData entities.User) (entities.User, error) {
	if err := r.userDb.Create(&userData).Error; err != nil {
		return userData, err
	}

	return userData, nil
}

func (r *userRepository) UpdateUser(userData entities.User) (entities.User, error) {
	if err := r.userDb.Save(&userData).Error; err != nil {
		return userData, err
	}

	return userData, nil
}

func (r *userRepository) UpdateUserById(userId int, updateData map[string]interface{}) (entities.User, error) {
	var userData entities.User
	if err := r.userDb.Model(&userData).Where("id = ?", userId).Updates(updateData).Error; err != nil {
		return userData, err
	}

	if err := r.userDb.Where("id = ?", userId).Find(&userData).Error; err != nil {
		return userData, err
	}

	return userData, nil
}

func (r *userRepository) DeleteUserById(userId int) (string, error) {
	if err := r.userDb.Where("id = ?", userId).Delete(&entities.User{}).Error; err != nil {
		return "error", err
	}

	return "success", nil
}

func (r *userRepository) FindAllUsers() ([]entities.User, error) {
	var usersData []entities.User
	if err := r.userDb.Find(&usersData).Error; err != nil {
		return usersData, err
	}

	return usersData, nil
}

func (r *userRepository) FindUserById(userId int) (entities.User, error) {
	var userData entities.User
	if err := r.userDb.Where("id = ?", userId).Find(&userData).Error; err != nil {
		return userData, err
	}

	return userData, nil
}

func (r *userRepository) FindUserByEmail(userEmail string) (entities.User, error) {
	var userData entities.User
	if err := r.userDb.Where("email = ?", userEmail).Find(&userData).Error; err != nil {
		return userData, err
	}

	return userData, nil
}
