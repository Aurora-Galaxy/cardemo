package service

import (
	"car/model"
	"car/pkg/util"
	"car/serializer"
)
// MoneyService 实现用户充值功能
type MoneyService struct {
	Money int `json:"money" form:"money"`
}

func (moneyService *MoneyService) MoneyRelevant(authorization string) serializer.Response {
	claims, _ := util.ParseToken(authorization)
	id := claims.Id
	money := moneyService.Money
	//充值
	var user1 model.User
	model.DB.Table("user").Where("id = ?", id).First(&user1)
	money = money + user1.Money
	err := model.DB.Table("user").Where("id = ?", id).Update("money", money).Error
	if err != nil {
		return serializer.Response{
			Status: 400,
			Data:   nil,
			Msg:    "数据库添加Money错误",
			Error:  err,
		}
	}
	//获取该用户信息
	var user model.User
	if err := model.DB.Where("id = ?", id).First(&user).Error; err != nil {
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