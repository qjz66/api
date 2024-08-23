package dao

import "API/model"

func (d *Dao) InsertUser(username, pwd string) (int64, error) {
	user := model.User{
		Username: username,
		Password: pwd,
	}
	err := d.Create(&user).Error
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (d *Dao) Login(username, pwd string) (int64, error) {
	user := model.User{}
	err := d.Where("username = ? and password = ?", username, pwd).First(&user).Error
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}
