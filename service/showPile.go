package service

import (
	"car/model"
	"car/serializer"
)

type ShowPileService struct {
}

func (service *ShowPileService) Show_Pile() serializer.Response {
	var piles []model.ChargingPile
	err := model.DB.Table("charging_pile").Find(&piles).Error
	if err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    "获取所有充电桩信息出错",
		}
	}
	return serializer.Response{
		Status: 200,
		Data:   serializer.BuildPiles(piles),
		Msg:    "success",
	}
}
