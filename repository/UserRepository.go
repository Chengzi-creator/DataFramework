package repository

import (
	"InterLibrarySystem/models"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

// FindByUsername 根据用户名查询用户
func (r UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
