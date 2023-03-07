package service

import (
	"car/model"
	"car/pkg/e"
	"car/pkg/util"
	"car/serializer"
	"github.com/jinzhu/gorm"
	logging "github.com/sirupsen/logrus"
)

type UserService struct {
	Username string `form:"user_name" json:"user_name"`
	PassWord string `form:"password" json:"password"`
}

type UserInformation struct {
}

// Register 用户注册
func (userService *UserService) Register() serializer.Response {
	var user model.User
	var count int
	model.DB.Model(&model.User{}).Where("user_name=?", userService.Username).
		First(&user).Count(&count) //如果count == 1 则证明数据库中存在这个人
	if count == 1 {
		return serializer.Response{
			Status: e.ErrorExistUser,
			Data:   nil,
			Msg:    "该用户名已经存在",
		}
	}
	user.UserName = userService.Username
	//密码加密
	err := user.SetPassWord(userService.PassWord)
	if err != nil {
		return serializer.Response{
			Status: 400,
			Data:   nil,
			Msg:    "密码加密时错误",
			Error:  err,
		}
	}
	//创建用户
	if err = model.DB.Create(&user).Error; err != nil {
		logging.Println(err)
		return serializer.Response{
			Status: 500,
			Data:   nil,
			Msg:    "创建用户时错误",
			Error:  err,
		}
	}
	return serializer.Response{
		Status: 200,
		Data:   nil,
		Msg:    "创建用户成功",
	}
}

// Login 用户登录
func (userService *UserService) Login() serializer.Response {
	var user model.User
	err := model.DB.Model(&model.User{}).Where("user_name=?", userService.Username).First(&user).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) { //没有查询到该用户
			return serializer.Response{
				Status: 400,
				Data:   nil,
				Msg:    "用户不存在，请先注册",
			}
		}
		return serializer.Response{
			Status: 500,
			Data:   nil,
			Msg:    "数据库错误",
			Error:  err,
		}
	}
	if !user.CheckPassWord(userService.PassWord) {
		return serializer.Response{
			Status: 400,
			Data:   nil,
			Msg:    "密码错误",
		}
	}
	//签发token
	token, err := util.GenerateToken(user.ID, userService.Username, userService.PassWord)
	if err != nil {
		return serializer.Response{
			Status: 500,
			Data:   nil,
			Msg:    "token签发错误",
			Error:  err,
		}
	}
	return serializer.Response{
		Status: 200,
		Data: serializer.TokenData{
			User:  serializer.BuildUser(user),
			Token: token,
		},
		Msg: "登录成功",
	}
}

func (userInformation *UserInformation) UserInfo(token string) serializer.Response {
	claim, _ := util.ParseToken(token)
	//获取该用户信息
	var user model.User
	if err := model.DB.Where("id = ?", claim.Id).First(&user).Error; err != nil {
		return serializer.Response{
			Status: 400,
			Data:   nil,
			Msg:    "数据库错误",
			Error:  err,
		}
	}
	//返回用户信息
	return serializer.Response{
		Status: 200,
		Data:   serializer.BuildUser(user),
		Msg:    "success",
	}
}
