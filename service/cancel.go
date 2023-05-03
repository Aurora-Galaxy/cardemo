package service

import (
	"car/model"
	"car/pkg/util"
	"car/serializer"

	logging "github.com/sirupsen/logrus"
)

type CancelService struct {
}

// CancelPile 取消正在使用的充电桩
func(service *CancelService) CancelPile(token string) serializer.Response{
	claim, _ := util.ParseToken(token)
	var count int
	err := model.DB.Table("charging_pile").Where("usering_id = ?", claim.UserName).Count(&count).Error
	if count == 0 {
		return serializer.Response{
			Status: 400,
			Data:   nil,
			Msg:    "该用户未使用充电桩，无法取消",
			Error:  nil,
		}
	}
	if count > 0{
		var pile model.ChargingPile
		err = model.DB.Table("charging_pile").Where("usering_id = ?", claim.UserName).First(&pile).Error
		if err != nil{
			logging.Info(err)
		}
		if pile.Status == 1{
			err = model.DB.Table("charging_pile").Where("usering_id = ?", claim.UserName).
				Updates(map[string]interface{}{
					"start_time" : "NULL",
					"end_time" : "NULL",
					"status" : 0,
					"usering_id" : "NULL",
			}).Error
			if err != nil{
				return serializer.Response{
					Status: 400,
					Data:   nil,
					Msg:    "取消充电出错",
					Error:  err,
				}
			}
		}
		if pile.Status == 2{
			err = model.DB.Table("charging_pile").Where("usering_id = ?", claim.UserName).
				Updates(map[string]interface{}{
					"start_time" : "NULL",
					"end_time" : "NULL",
				    "usering_id" : "NULL",
				}).Error
			if err != nil{
				return serializer.Response{
					Status: 400,
					Data:   nil,
					Msg:    "取消充电出错",
					Error:  err,
				}
			}
		}

	}
	return serializer.Response{
		Status: 200,
		Data:   "",
		Msg:    "success",
		Error:  nil,
	}
}

// CancelReserve 取消正在使用的充电桩
func(service *CancelService) CancelReserve(token string) serializer.Response{
	claim, _ := util.ParseToken(token)
	var count int
	err := model.DB.Table("charging_pile").Where("user_id = ?", claim.UserName).Count(&count).Error
	if count == 0 {
		return serializer.Response{
			Status: 400,
			Data:   nil,
			Msg:    "该用户未预约，无法取消",
			Error:  nil,
		}
	}
	if count > 0{
		var pile model.ChargingPile
		err = model.DB.Table("charging_pile").Where("user_id = ?", claim.UserName).First(&pile).Error
		if err != nil{
			logging.Info(err)
		}
		if pile.EndTime != "NULL"{
			err = model.DB.Table("charging_pile").Where("user_id = ?", claim.UserName).
				Updates(map[string]interface{}{
					"reserve_start_time" : "NULL",
					"reserve_end_time" : "NULL",
					"status" : 1,
					"user_id" : "NULL",
				}).Error
			if err != nil{
				return serializer.Response{
					Status: 400,
					Data:   nil,
					Msg:    "取消预约出错",
					Error:  err,
				}
			}
		}else{
			err = model.DB.Table("charging_pile").Where("user_id = ?", claim.UserName).
				Updates(map[string]interface{}{
					"reserve_start_time" : "NULL",
					"reserve_end_time" : "NULL",
					"status" : 0,
					"user_id" : "NULL",
				}).Error
			if err != nil{
				return serializer.Response{
					Status: 400,
					Data:   nil,
					Msg:    "取消预约出错",
					Error:  err,
				}
			}
		}

	}
	return serializer.Response{
		Status: 200,
		Data:   "",
		Msg:    "success",
		Error:  nil,
	}
}


