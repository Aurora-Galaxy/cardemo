package api

import (
	"car/service"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

func BindCar(c *gin.Context) {
	var service service.CarService
	authorization := c.Request.Header.Get("Authorization")
	err := c.ShouldBind(&service)
	if err != nil {
		c.JSON(400, "绑定车牌时出错")
		logging.Info(err)
	} else {
		res := service.CarRelevant(authorization)
		c.JSON(200, res)
	}
}
