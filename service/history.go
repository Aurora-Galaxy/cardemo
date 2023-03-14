package service

import (
	"car/model"
	"car/pkg/util"
	"car/serializer"
	"fmt"
	"time"
)

type ShowHistoryService struct {
}

// AddHistoryRecord 将用户使用充电桩的信息存入历史记录表中
func AddHistoryRecord(uid uint, startTime int64, endTime int64, pileId uint) error {
	var record model.HistoryRecord
	duration := time.Duration(endTime-startTime) * time.Millisecond
	hours := fmt.Sprint(duration.Hours())
	record = model.HistoryRecord{
		Uid:    uid,
		Date:   fmt.Sprint(time.UnixMilli(startTime)),
		Time:   hours,
		PileId: pileId,
	}
	err := model.DB.Table("history_record").Create(&record).Error
	if err != nil {
		return err
	}
	return nil
}

// ShowHistory 显示用户使用充电桩的历史记录
func (service *ShowHistoryService) ShowHistory(token string) serializer.Response {
	claim, _ := util.ParseToken(token)
	var history []model.HistoryRecord
	err := model.DB.Table("history_record").Where("uid = ?", claim.Id).Find(&history).Error
	if err != nil {
		return serializer.Response{
			Status: 400,
			Data:   nil,
			Msg:    "查找历史记录时数据库错误",
			Error:  err,
		}
	}
	return serializer.Response{
		Status: 200,
		Data:   serializer.BuildHistorys(history),
		Msg:    "success",
		Error:  nil,
	}
}
