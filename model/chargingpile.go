package model

import (
	"github.com/jinzhu/gorm"
)

// ChargingPile 充电桩信息
type ChargingPile struct {
	gorm.Model
	//Id        int
	StartTime        string
	EndTime          string
	Status           int    //状态（0 空闲  1 使用中 2 预约）
	UseringId        string //正在使用用户的用户名
	UserId           string //预约充电桩的用户名
	ReserveStartTime string //预约开始时间
	ReserveEndTime   string //预约结束时间

}
