package repositories

import (
	"gorm.io/gorm"
	"silih_a3/entities"
)

type UserRepository interface {
	InsertUser(userData entities.User) (entities.User, error)
	UpdateUser(userData entities.User) (entities.User, error)
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

func (r *userRepository) FindAllUsers() ([]entities.User, error) {
	var users []entities.User
	if err := r.userDb.Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

func (r *userRepository) FindUserById(id int) (entities.User, error) {
	var user entities.User
	if err := r.userDb.Where("id = ?", id).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) FindUserByEmail(userEmail string) (entities.User, error) {
	var user entities.User
	if err := r.userDb.Where("email = ?", userEmail).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
