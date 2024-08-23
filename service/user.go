package service

import (
	"API/model"
	"API/util"
)

func (svc *Service) UserReg(user model.User) (int64, error) {
	pwd := user.Password
	username := user.Username
	id, err := svc.d.InsertUser(username, pwd)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (svc *Service) Login(user model.User) (int64, string, error) {
	username := user.Username
	pwd := user.Password
	id, err := svc.d.Login(username, pwd)
	if err != nil {
		return 0, "", err
	}
	token, err := util.GenerateToken(id)
	if err != nil {
		return 0, "", err
	}
	return id, token, nil
}
