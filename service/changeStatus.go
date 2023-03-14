package service

import (
	"car/model"
	"car/pkg/util"
	"car/serializer"
	"strconv"
	"time"
)

type ChangeStatusService struct {
	Id   int `json:"id" form:"id"`
	Time int `json:"time" form:"time"`
}

type ReserveService struct {
	Id   int `json:"id" form:"id"`
	Time int `json:"time" form:"time"`
}

// Change 改变充电桩的状态
func (service ChangeStatusService) Change(token string) serializer.Response {
	claims, _ := util.ParseToken(token)
	StartTime := time.Now().UnixMilli()
	var Endtime int64
	if service.Time == 1 {
		Endtime = StartTime + time.Hour.Milliseconds()
	}
	if service.Time == 2 {
		Endtime = StartTime + (time.Hour * 2).Milliseconds()
	}
	if service.Time == 3 {
		Endtime = StartTime + (time.Hour * 3).Milliseconds()
	}
	if service.Time == 4 {
		Endtime = StartTime + (time.Hour * 4).Milliseconds()
	}
	if service.Time == 5 {
		Endtime = StartTime + (time.Hour * 5).Milliseconds()
	}
	err := model.DB.Table("charging_pile").Where("id = ?", service.Id).
		Update("start_time", strconv.FormatInt(StartTime, 10)).Update("end_time", strconv.FormatInt(Endtime, 10)).Error
	if err != nil {
		return serializer.Response{
			Status: 400,
			Data:   nil,
			Msg:    "添加充电桩开始和结束时间出错",
			Error:  err,
		}
	}
	//添加使用记录
	err = AddHistoryRecord(claims.Id, StartTime, Endtime, uint(service.Id))
	if err != nil {
		return serializer.Response{
			Status: 400,
			Data:   nil,
			Msg:    "历史记录中添加充电桩开始和结束时间出错",
			Error:  err,
		}
	}
	err = model.DB.Table("charging_pile").Where("id = ?", service.Id).
		Update("status", 1).Error
	if err != nil {
		return serializer.Response{
			Status: 400,
			Data:   nil,
			Msg:    "更改充电桩状态出错",
			Error:  err,
		}
	}
	err = model.DB.Table("charging_pile").Where("id = ?", service.Id).
		Update("usering_id", claims.UserName).Error
	if err != nil {
		return serializer.Response{
			Status: 400,
			Data:   nil,
			Msg:    "添加充电桩正在使用用户时出错",
			Error:  err,
		}
	}
	//获取充电桩信息
	var pile model.ChargingPile
	err = model.DB.Table("charging_pile").Where("id = ?", service.Id).First(&pile).Error
	if err != nil {
		return serializer.Response{
			Status: 400,
			Data:   nil,
			Msg:    "返回充电桩信息出错",
			Error:  err,
		}
	}
	return serializer.Response{
		Status: 200,
		Data:   serializer.BuildPile(pile),
		Msg:    "success",
	}
}

//ReservePile 预约充电桩
func (service *ReserveService) ReservePile(token string) serializer.Response {
	claim, _ := util.ParseToken(token)
	//查找该用户是否已经预约
	var count int
	model.DB.Table("charging_pile").Where("user_id = ?", claim.UserName).Count(&count)
	if count > 0 {
		return serializer.Response{
			Status: 400,
			Data:   nil,
			Msg:    "该用户已经预约，请勿重复预约",
			Error:  nil,
		}
	}
	err := model.DB.Table("charging_pile").Where("id = ?", service.Id).Update("status", 2).Error
	if err != nil {
		return serializer.Response{
			Status: 400,
			Data:   nil,
			Msg:    "更改充电桩预约状态出错",
			Error:  err,
		}
	}
	err = model.DB.Table("charging_pile").Where("id = ?", service.Id).
		Update("user_id", claim.UserName).Error
	if err != nil {
		return serializer.Response{
			Status: 400,
			Data:   nil,
			Msg:    "添加预约用户时出错",
			Error:  err,
		}
	}
	StartTime := time.Now().UnixMilli()
	var Endtime int64
	if service.Time == 1 {
		Endtime = StartTime + time.Hour.Milliseconds()
	}
	if service.Time == 2 {
		Endtime = StartTime + (time.Hour * 2).Milliseconds()
	}
	if service.Time == 3 {
		Endtime = StartTime + (time.Hour * 3).Milliseconds()
	}
	if service.Time == 4 {
		Endtime = StartTime + (time.Hour * 4).Milliseconds()
	}
	if service.Time == 5 {
		Endtime = StartTime + (time.Hour * 5).Milliseconds()
	}
	err = model.DB.Table("charging_pile").Where("id = ?", service.Id).
		Update("reserve_start_time", strconv.FormatInt(StartTime, 10)).
		Update("reserve_end_time", strconv.FormatInt(Endtime, 10)).Error
	if err != nil {
		return serializer.Response{
			Status: 400,
			Data:   nil,
			Msg:    "添加充电桩预约开始和结束时间出错",
			Error:  err,
		}
	}
	//获取充电桩信息
	var pile model.ChargingPile
	err = model.DB.Table("charging_pile").Where("id = ?", service.Id).First(&pile).Error
	if err != nil {
		return serializer.Response{
			Status: 400,
			Data:   nil,
			Msg:    "返回充电桩信息出错",
			Error:  err,
		}
	}
	return serializer.Response{
		Status: 200,
		Data:   serializer.BuildPile(pile),
		Msg:    "success",
	}
}
