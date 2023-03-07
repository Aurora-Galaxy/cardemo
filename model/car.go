package model

import "github.com/jinzhu/gorm"

type Car struct {
	gorm.Model
	User        User   `gorm:"ForeignKey:CarBossId"` //设置外键，将用户和车辆进行关联
	CarBossId   uint   `gorm:"not null"`             //车主ID
	CarBossName string //车主用户名
	CarName     string //车牌名称
	CarNum      string //车牌号
	CarImages   string //车照片
}
