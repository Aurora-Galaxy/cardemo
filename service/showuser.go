package service

import (
	"car/model"
	"car/serializer"
)

type ShowUserService struct {
}

func (service *ShowUserService) Show_User() serializer.Response {
	var users []model.User
	err := model.DB.Table("user").Find(&users).Error
	if err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    "获取所有用户信息出错",
		}
	}
	return serializer.Response{
		Status: 200,
		Data:   serializer.BuildUsers(users),
		Msg:    "success",
	}
}
