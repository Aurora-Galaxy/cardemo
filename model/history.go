package model

import "github.com/jinzhu/gorm"

type HistoryRecord struct {
	gorm.Model
	User   User   `gorm:"ForeignKey:Uid"` //设置外键，将用户和任务进行关联
	Uid    uint   `gorm:"not null"`
	PileId uint   //所使用的充电桩id号
	Date   string //历史充电时间
	Time   string //历史充电时长
}
