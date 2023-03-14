package api

import (
	service2 "car/service"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

func GetHistory(c *gin.Context) {
	var service service2.ShowHistoryService
	authorization := c.Request.Header.Get("Authorization")
	err := c.ShouldBind(&service)
	if err != nil {
		c.JSON(400, "获取用户历史使用记录错误")
		logging.Info(err)
	} else {
		res := service.ShowHistory(authorization)
		c.JSON(200, res)
	}
}
