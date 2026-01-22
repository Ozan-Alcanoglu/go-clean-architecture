package repositories

import (
	"gobackend/domain/interfaces"
	"gobackend/domain/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepository {

	return &userRepository{db: db}

}

func (ur *userRepository) CreateUser(user *models.User) error {

	return ur.db.Create(user).Error

}

func (ur *userRepository) GetById(id uuid.UUID) (*models.User, error) {

	var user models.User

	err := ur.db.Where("id= ?", id).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil

}

func (ur *userRepository) GetByEmail(email string) (*models.User, error) {

	var user models.User

	err := ur.db.Where("email= ?", email).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil

}

func (ur *userRepository) DeleteUserById(id uuid.UUID) error {
	err := ur.db.Where("id = ?", id).Delete(&models.User{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) UpdateUser(user *models.User) (*models.User, error) {
	if err := ur.db.Save(user).Error; err != nil {
		return nil, err
	}

	var updated models.User
	if err := ur.db.Where("id = ?", user.ID).First(&updated).Error; err != nil {
		return nil, err
	}

	return &updated, nil
}
