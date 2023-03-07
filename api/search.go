package api

import (
	"car/service"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

func Search(c *gin.Context) {
	var service service.SearchService
	err := c.ShouldBind(&service)
	if err != nil {
		c.JSON(400, "search出错")
		logging.Info(err)
	} else {
		res := service.SearchUser()
		c.JSON(200, res)
	}
}
