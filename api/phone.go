package api

import (
	"car/service"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

func BindPhone(c *gin.Context) {
	var service service.PhoneService
	authorization := c.Request.Header.Get("Authorization")
	err := c.ShouldBind(&service)
	if err != nil {
		c.JSON(400, "绑定手机时出错")
		logging.Info(err)
	} else {
		res := service.PhoneRelevant(authorization)
		c.JSON(200, res)
	}
}

func GetCode(c *gin.Context) {
	var service service.CodeService
	if err := c.ShouldBind(&service); err == nil {
		res := service.SendMsg()
		c.JSON(200, res)
	} else {
		c.JSON(400, "发送验证码出错")
		logging.Info(err)
	}
}
