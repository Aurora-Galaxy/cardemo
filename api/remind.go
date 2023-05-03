package api

import (
	"car/service"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

func RemindAPI(c *gin.Context) {
	var service service.RemindService
	err := c.ShouldBind(&service)
	if err != nil {
		c.JSON(400, "remind出错")
		logging.Info(err)
	} else {
		res := service.Remind()
		c.JSON(200, res)
	}
}
