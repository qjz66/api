package service

import "API/model"

func (svc *Service) UserReg(user model.User) (int64, error) {
	pwd := user.Password
	username := user.Username
	id, err := svc.d.InsertUser(username, pwd)
	if err != nil {
		return 0, err
	}
	return id, nil
}
