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

// FindByUserID 根据用户ID查询用户
func (r *UserRepository) FindByUserID(userid int) (*models.User, error) {
	var user models.User
	err := utils.DB.Where("id = ?", userid).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateByUserID 根据用户ID更改用户信息
func (r *UserRepository) UpdateByUserID(user *models.User) error {
	err := utils.DB.Save(user).Error
	if err != nil {
		return err
	}
	return nil
}

// RegisterNewUser 注册新用户
func (r *UserRepository) RegisterNewUser(user *models.User) error {
	err := utils.DB.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}
