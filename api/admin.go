package api

import (
	"car/service"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

// AdminLogin 用户登录接口
func AdminLogin(c *gin.Context) {
	var AdminLogin service.AdminService
	err := c.ShouldBind(&AdminLogin)
	if err != nil {
		c.JSON(400, err)
	} else {
		res := AdminLogin.AdminLogin()
		c.JSON(200, res)
	}
}

// etUsers 获取用户所有信息
func GetUsers(c *gin.Context) {
	var service service.ShowUserService
	err := c.ShouldBind(&service)
	if err != nil {
		c.JSON(400, "获取充电桩信息时出错")
		logging.Info(err)
	} else {
		res := service.Show_User()
		c.JSON(200, res)
	}
}

//删除用户
func Delete(c *gin.Context){
	service := service.UserDeleteService{}
	res := service.DeleteUser(c.Param("id"))
	c.JSON(200, res)
}