package service

import (
	"car/model"
	logging "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

// ReserveNotice 预约后，等待时间小于10分钟时，给用户发送短信提醒
func ReserveNotice() {
	var piles []model.ChargingPile
	err := model.DB.Table("charging_pile").Find(&piles).Error
	if err != nil {
		logging.Info(err)
	}
	for _, v := range piles {
		endTime, _ := strconv.ParseInt(v.EndTime, 10, 64)
		if (endTime - time.Now().UnixMilli()) < (time.Minute * 10).Milliseconds() {
			var user model.User
			err := model.DB.Table("user").Where("user_name = ?", v.UserId).First(&user).Error
			if err != nil {
				logging.Info(err)
			}
			//phone := user.Phone
			//发送短信提醒用户
		}
	}
}
