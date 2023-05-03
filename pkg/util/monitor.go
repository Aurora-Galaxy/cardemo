package util

import (
	"car/model"
	logging "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

type EndTime struct {
	endtime string
}

//Monitor 检测mysql数据库，end_time的变化
func Monitor(pile_id int){
	var result model.ChargingPile
	err := model.DB.Table("charging_pile").Where("id = ?",pile_id).First(&result).Error
	if err != nil{
		logging.Info(err)
	}
	nowTime := time.Now().UnixMilli()
	endTime , _ := strconv.ParseInt(result.EndTime , 10 ,64)
	reserveTime , _ := strconv.ParseInt(result.ReserveStartTime , 10 ,64)
	if nowTime > endTime  && endTime != 0 {
		var user model.User
		err = model.DB.Table("user").Where("user_name = ?",result.UseringId).First(&user).Error
		if err != nil{
			logging.Info(err)
		}
		// var text []byte
		// text = []byte("充电完成，请及时挪走您的车辆，谢谢配合")
		text := "充电完成，请及时挪走您的车辆，谢谢配合"
		err = SendMail(text , user.Email)
		if err != nil {
			logging.Info(err)
		}
		if result.Status == 2{
			model.DB.Table("charging_pile").Where("id = ?",pile_id).Updates(map[string]interface{}{
				"start_time" : "NULL",
				"end_time" :  "NULL",
				"usering_id" : "NULL",
			})
		}else{
			model.DB.Table("charging_pile").Where("id = ?",pile_id).Updates(map[string]interface{}{
				"start_time" : "NULL",
				"end_time" :  "NULL",
				"status" : 0,
				"usering_id" : "NULL",
			})
		}
	}
	if reserveTime != 0 && nowTime >= reserveTime{
		var user model.User
		err := model.DB.Table("user").Where("user_name = ?",result.UserId).First(&user).Error
		if err!= nil{
			logging.Info(err)
		}
		// var text []byte
		// text = []byte("已到达预约时间，充电开始，钱款已自动扣除")
		text := "已到达预约时间，充电开始，钱款已自动扣除"
		err = SendMail(text , user.Email)
		if err != nil {
			logging.Info(err)
		}
		model.DB.Table("charging_pile").Where("id = ?",pile_id).Updates(map[string]interface{}{
			"start_time" : result.ReserveStartTime,
			"end_time" :  result.ReserveEndTime,
			"status" : 1,
			"usering_id" : result.UserId,
			"user_id":"NULL",
			"reserve_start_time" : "NULL",
			"reserve_end_time" : "NULL",
		})
		time1 , _ := strconv.Atoi(result.ReserveEndTime)
		time2 , _ := strconv.Atoi(result.ReserveStartTime)
		time := (time1 - time2) / 1000 / 60 / 60
		err = model.DB.Table("user").Where("id = ?",user.ID).Update("money",user.Money - time).Error
		if err!= nil{
			logging.Info(err)
		}
	}
}

func CheckPile(){
	for{
		go Monitor(1)
		time.Sleep(time.Second)
		go Monitor(2)
		time.Sleep(time.Second)
		go Monitor(3)
		time.Sleep(time.Second)
		go Monitor(4)
		time.Sleep(time.Second)
		go Monitor(5)
		time.Sleep(time.Second)
	}
}
