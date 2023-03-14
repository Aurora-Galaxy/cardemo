package api

import (
	service2 "car/service"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

// ChangeStatus 改变充电桩的使用状态
func ChangeStatus(c *gin.Context) {
	var service service2.ChangeStatusService
	authorization := c.Request.Header.Get("Authorization")
	err := c.ShouldBind(&service)
	if err != nil {
		c.JSON(400, "改变充电桩状态时出错")
		logging.Info(err)
	} else {
		res := service.Change(authorization)
		c.JSON(200, res)
	}
}

// GetPile 获取充电桩所有信息
func GetPile(c *gin.Context) {
	var service service2.ShowPileService
	err := c.ShouldBind(&service)
	if err != nil {
		c.JSON(400, "获取充电桩信息时出错")
		logging.Info(err)
	} else {
		res := service.Show_Pile()
		c.JSON(200, res)
	}
}

// Reserve 充电桩预约
func Reserve(c *gin.Context) {
	var service service2.ReserveService
	authorization := c.Request.Header.Get("Authorization")
	err := c.ShouldBind(&service)
	if err != nil {
		c.JSON(400, "预约错误")
		logging.Info(err)
	} else {
		res := service.ReservePile(authorization)
		c.JSON(200, res)
	}
}
