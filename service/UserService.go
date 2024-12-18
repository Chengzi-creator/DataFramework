package service

import (
	"InterLibrarySystem/models"
	"InterLibrarySystem/repository"
	"errors"
)

type UserService struct {
	UserRepo *repository.UserRepository
}

// Login 用户登录认证
func (s *UserService) Login(username string, password string) (*models.User, error) {
	//查询用户
	user, err := s.UserRepo.FindByUsername(username)
	if err != nil {
		return nil, errors.New("该用户不存在")
	}

	//验证密码
	if user.Password != password {
		return nil, errors.New("密码错误")
	}
	return user, nil
}
