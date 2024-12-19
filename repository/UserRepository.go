package repository

import (
	"InterLibrarySystem/models"
	"InterLibrarySystem/utils"
)

type UserRepository struct {
}

// FindByUsername 根据用户名查询用户
func (r *UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	err := utils.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByUserId(userid int) (*models.User, error) {
	var user models.User
	err := utils.DB.Where("user_id = ?", userid).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
