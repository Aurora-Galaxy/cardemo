package api

import (
	"car/service"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

func BindEmail(c *gin.Context) {
	var service service.EmailService
	authorization := c.Request.Header.Get("Authorization")
	err := c.ShouldBind(&service)
	if err != nil {
		c.JSON(400, "绑定邮箱时出错")
		logging.Info(err)
	} else {
		res := service.EmailRelevant(authorization)
		c.JSON(200, res)
	}
}
