package service

import (
	"car/model"
	"car/serializer"
	"github.com/jinzhu/gorm"
)

type SearchService struct {
	CarNumber string `json:"car_number" form:"car_number"`
}

func (service *SearchService) SearchUser() serializer.Response {
	var user model.User
	err := model.DB.Model(&model.User{}).Where("car_n_umber=?", service.CarNumber).First(&user).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) { //没有查询到该车牌对应用户
			return serializer.Response{
				Status: 400,
				Data:   nil,
				Msg:    "绑定该车牌的用户不存在，请确保车牌输入正确！",
				Error:  err,
			}
		}
		return serializer.Response{
			Status: 500,
			Data:   nil,
			Msg:    "数据库错误",
			Error:  err,
		}
	}
	return serializer.Response{
		Status: 200,
		Data:   serializer.BuildUser(user),
		Msg:    "success",
		Error:  err,
	}
}
