package service

import (
	"car/model"
	"car/pkg/util"
	"car/serializer"
)

type CarService struct {
	CarNumber string `json:"car_number" form:"car_number"`
}

func (carService *CarService) CarRelevant(authorization string) serializer.Response {
	claims, _ := util.ParseToken(authorization)
	id := claims.Id
	carNumber := carService.CarNumber
	//绑定邮箱
	err := model.DB.Table("user").Where("id = ?", id).Update("car_n_umber", carNumber).Error
	if err != nil {
		return serializer.Response{
			Status: 400,
			Data:   nil,
			Msg:    "数据库添加车牌号错误",
			Error:  err,
		}
	}
	//获取该用户信息
	var user model.User
	if err := model.DB.First(&user).Where("id = ?", id).Error; err != nil {
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
