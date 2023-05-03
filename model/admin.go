package model

import (
"github.com/jinzhu/gorm"
"golang.org/x/crypto/bcrypt"
)

type Admin struct {
	gorm.Model
	UserName       string
	PasswordDigest string
}

// 设置密码
func (Admin *Admin) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	Admin.PasswordDigest = string(bytes)
	return nil
}

// 校验密码Admin
func (Admin *Admin) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(Admin.PasswordDigest), []byte(password))
	return err == nil
}


