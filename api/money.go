package api

import (
	"car/service"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

func BindMoney(c *gin.Context) {
	var service service.MoneyService
	authorization := c.Request.Header.Get("Authorization")
	err := c.ShouldBind(&service)
	if err != nil {
		c.JSON(400, "充值时出错")
		logging.Info(err)
	} else {
		res := service.MoneyRelevant(authorization)
		c.JSON(200, res)
	}
}