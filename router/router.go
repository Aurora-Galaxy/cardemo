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
		v1.POST("remind", api.RemindAPI)           //提醒用户挪车
		authed := v1.Group("/")
		authed.Use(middleware.JWT()) //中间件验证身份
		{
			authed.POST("user/email", api.BindEmail) //绑定邮箱
			authed.POST("user/phone", api.BindPhone) //绑定手机号
			authed.GET("user/show", api.UserShow)    //展示用户信息
			authed.POST("user/carnumber", api.BindCar)   //绑定车牌号
			authed.POST("user/change", api.ChangeStatus) //更改充电桩使用状态
			authed.POST("user/reserve", api.Reserve)    //充电桩预约
			authed.GET("user/history", api.GetHistory)  //获取用户使用充电桩历史记录
			authed.POST("user/money",api.BindMoney)   //用户充值
			authed.POST("user/cancelpile",api.Cancel_pile)   //取消充电
			authed.POST("user/cancelreserve",api.Cancel_ReservePile)   //取消预约
		}
	}
	v2 := r.Group("api/v2")
	{
		v2.POST("admin/login", api.AdminLogin)       //管理员登录
		v2.GET("admin/show",api.GetUsers)            //获取用户信息
		v2.DELETE("admin/delete/:id",api.Delete)     //删除用户
		v2.GET("admin/showpile",api.GetPile)		//获取充电桩信息
	}
	return r
}
