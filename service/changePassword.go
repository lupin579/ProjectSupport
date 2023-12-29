package service

import (
	"eee/cache"
	"eee/dao/mysql"
)

type ChangePassword struct {
	Uname    string `json:"uname"`
	Code     string `json:"code"`
	Password string `json:"password"`
}

func (changePassword *ChangePassword) ChangePasswordService() error {
	email, err := mysql.GetEmailByUname(changePassword.Uname)
	if err != nil {
		return err
	}
	if err = cache.SendEmailCode(changePassword.Uname, email); err != nil {
		return err
	}
	return nil
}

func (changePassword *ChangePassword) EmailCodeValidator() error {
	err := cache.ValidateEmailCode(changePassword.Uname, changePassword.Code)
	if err != nil {
		return err
	}
	err = mysql.ChangePassword(changePassword.Uname, changePassword.Password)
	return err
}
