package service

import (
	"car/model"
	"car/pkg/util"
	"car/serializer"
)

//绑定邮箱

type EmailService struct {
	//OperationType int `json:"operation_type" form:"operation_type"`
	Email string `json:"email" form:"email"`
}

func (emailService *EmailService) EmailRelevant(authorization string) serializer.Response {
	claims, _ := util.ParseToken(authorization)
	id := claims.Id
	email := emailService.Email
	//绑定邮箱
	err := model.DB.Table("user").Where("id = ?", id).Update("email", email).Error
	if err != nil {
		return serializer.Response{
			Status: 400,
			Data:   nil,
			Msg:    "数据库添加Email错误",
			Error:  err,
		}
	}
	/*		}else if emailService.OperationType == 2{  //解绑邮箱
				err := model.DB.Table("user").Where("id = ?",id).Update("email","").Error
				if err != nil{
					return serializer.Response{
						Status: 400,
						Data:   nil,
						Msg:    "数据库删除Email错误",
						Error:  err.Error(),
					}
				}
			}
	*/
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
