package middleware

import (
	"car/pkg/util"
	"github.com/gin-gonic/gin"
	"time"
)

// JWT 中间件验证
func JWT() gin.HandlerFunc {
	return func(context *gin.Context) {
		code := 200
		token := context.GetHeader("Authorization")
		if token == "" { //token为空
			code = 404
		} else {
			claim, err := util.ParseToken(token) //对token进行解析
			if err != nil {
				code = 401 //说明token是无权限的是假的
			} else if time.Now().Unix() > claim.ExpiresAt { //token已经过期
				code = 403
			}
		}
		if code != 200 {
			context.JSON(400, gin.H{
				"status": 400,
				"msg":    "token解析错误",
			})
			context.Abort() //中止信号
			return
		}
		context.Next()
	}
}
