package api

import (
	"car/service"
	"github.com/gin-gonic/gin"
	//logging "github.com/sirupsen/logrus"
)

// Cancel_pile 取消充电
func Cancel_pile(c *gin.Context) {
	var service service.CancelService
	authorization := c.Request.Header.Get("Authorization")
	//err := c.ShouldBind(&service)
	// if err != nil {
	// 	c.JSON(400, "cancel_pile出错")
	// 	logging.Info(err)
	// } else {
	res := service.CancelPile(authorization)
	c.JSON(200, res)
	
}


// Cancel_ReservePile 取消充电
func Cancel_ReservePile(c *gin.Context) {
	var service service.CancelService
	authorization := c.Request.Header.Get("Authorization")
	// err := c.ShouldBind(&service)
	// if err != nil {
	// 	c.JSON(400, "Cancel_ReservePile出错")
	// 	logging.Info(err)
	// } else {
	res := service.CancelReserve(authorization)
	c.JSON(200, res)
	
}
