package api

import (
	"car/service"
	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册接口
func UserRegister(c *gin.Context) {
	var userRegister service.UserService
	err := c.ShouldBind(&userRegister)
	if err != nil {
		c.JSON(400, err)
	} else {
		res := userRegister.Register()
		c.JSON(200, res)
	}
}

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	var UserLogin service.UserService
	err := c.ShouldBind(&UserLogin)
	if err != nil {
		c.JSON(400, err)
	} else {
		res := UserLogin.Login()
		c.JSON(200, res)
	}
}

// UserShow 获取用户信息接口
func UserShow(c *gin.Context) {
	var userShow service.UserInformation
	authorization := c.Request.Header.Get("Authorization")
	err := c.ShouldBind(&userShow)
	if err != nil {
		c.JSON(400, err)
	} else {
		res := userShow.UserInfo(authorization)
		c.JSON(200, res)
	}
}
