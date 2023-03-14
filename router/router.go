package router

import (
	"car/api"
	"car/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(sessions.Sessions("mysession", store))
	v1 := r.Group("api/v1")
	{
		//用户操作
		v1.POST("user/register", api.UserRegister) //注册
		v1.POST("user/login", api.UserLogin)       //登录
		v1.POST("picture", api.DistinguishPicture) //车牌号识别
		v1.POST("code", api.GetCode)               //发送验证码
		v1.POST("search", api.Search)              //查找车牌号对应车主
		v1.GET("show", api.GetPile)                //获取所有充电桩信息
		authed := v1.Group("/")
		authed.Use(middleware.JWT()) //中间件验证身份
		{
			authed.POST("user/email", api.BindEmail) //绑定邮箱
			authed.POST("user/phone", api.BindPhone) //绑定手机号
			authed.GET("user/show", api.UserShow)    //展示用户信息
			//authed.POST("user/picture",api.DistinguishPicture)
			authed.POST("user/carnumber", api.BindCar)   //绑定车牌号
			authed.POST("user/change", api.ChangeStatus) //更改充电桩使用状态
			authed.POST("/user/reserve", api.Reserve)    //充电桩预约
			authed.GET("/user/history", api.GetHistory)  //获取用户使用充电桩历史记录
		}
	}
	return r
}
