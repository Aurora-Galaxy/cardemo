package service

import (
	"car/model"
	"car/pkg/util"
	"car/serializer"
	//"fmt"

	"github.com/jinzhu/gorm"
	logging "github.com/sirupsen/logrus"
)

type RemindService struct {
	CarNumber string `json:"car_number" form:"car_number"`
}

// Remind 输入车牌号，根据车牌号提醒用户
func (service *RemindService ) Remind() serializer.Response {
	var user model.User
	//fmt.Println(service.CarNumber)
	err := model.DB.Model(&model.User{}).Where("car_n_umber = ?", service.CarNumber).First(&user).Error
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
	text := "您好，您的车子充电已经完成。其他用户需要使用该充电桩，请您尽快挪走您的车辆，谢谢配合！！！"
	//fmt.Println(user.Email)
	err = util.SendMail(text , user.Email)
	if err != nil{
		logging.Fatal(err)
		return serializer.Response{
			Status: 400,
			Data:   nil,
			Msg:    "提醒用户邮件发送错误",
			Error:  err,
		}
	}
	return serializer.Response{
		Status: 200,
		Data:   "",
		Msg:    "success",
		Error:  err,
	}
}
