package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

//User 用户模型
type User struct {
	gorm.Model
	UserName  string `gorm:"unique"`
	PassWord  string //存储密文
	Email     string
	Phone     string
	CarNUmber string
	Money     int
	//Bantime   int
	//OpenId    string `gorm:"unique"`
	//Relations []User `gorm:"many2many:relation; association_jointable_foreignkey:relation_id"`
}

const PassWordCost = 12 //密码加密等级

// SetPassWord 设置密码
func (user *User) SetPassWord(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost) //密码加密
	if err != nil {
		return err
	}
	user.PassWord = string(bytes)
	return nil
}

// CheckPassWord 校验密码
func (user *User) CheckPassWord(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PassWord), []byte(password))
	if err != nil {
		return false
	}
	return true
}
