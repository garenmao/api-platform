package service

import (
	"api-platform/global"
	"api-platform/model/domain"
	utils "api-platform/util"
)

type UserService struct{}

func (us *UserService) CreateUser(user *domain.User) {
	// 对密码加密
	user.PassWord = utils.BcryptHash(user.PassWord)
	global.DB.Create(user)
}
