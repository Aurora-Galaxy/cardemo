package api

import (
	"car/service"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

func DistinguishPicture(c *gin.Context) {
	var service service.PictureService
	//authorization := c.Request.Header.Get("Authorization")
	err := c.ShouldBind(&service)
	if err != nil {
		c.JSON(400, "DistinguishPicture出错")
		logging.Info(err)
	} else {
		res := service.PictureOCR()
		c.JSON(200, res)
	}
}
