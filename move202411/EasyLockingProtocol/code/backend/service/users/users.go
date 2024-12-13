package users

import (
	"tokenVestingProtocol/global"
	"tokenVestingProtocol/model/users"
)

type UsersService struct{}

var UsersServiceApp = new(UsersService)

func (u *UsersService) GetUserInfo(user *users.Users) (err error) {
	if err := user.GetUserInfo(); err != nil {
		return err
	}
	return nil
}

func (u *UsersService) CreateUser(user *users.Users) error {
	if err := user.CreateUser(); err != nil {
		return err
	}
	return nil
}

func (u *UsersService) FindUserById(id int) (*users.Users, error) {
	var user users.Users
	if err := global.GVA_DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
