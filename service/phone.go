package service

import (
	"car/cache"
	"car/conf"
	"car/model"
	"car/pkg/util"
	"car/serializer"
	"fmt"
	unims "github.com/apistd/uni-go-sdk/sms"
	logging "github.com/sirupsen/logrus"
	"math/rand"
	"strconv"
	"time"
)

type PhoneService struct {
	Phone string `json:"phone" form:"phone"`
	Code  string `json:"code" form:"code"`
}

type CodeService struct {
	Phone string `json:"phone" form:"phone"`
}

func (phoneService *PhoneService) PhoneRelevant(authorization string) serializer.Response {
	claims, _ := util.ParseToken(authorization)
	id := claims.Id
	codeString := phoneService.Phone + "code"
	//验证验证码是否准确
	if err := cache.RedisClient.Get(codeString).Err(); err != nil {
		fmt.Println(err)
	}
	RedisCode := fmt.Sprintf("%s", cache.RedisClient.Get(codeString))[21:] //取出验证码进行比较
	if phoneService.Code != RedisCode {
		return serializer.Response{
			Status: 400,
			Data:   nil,
			Msg:    "验证码错误",
		}
	}
	phone := phoneService.Phone
	//绑定手机
	err := model.DB.Table("user").Where("id = ?", id).Update("phone", phone).Error
	if err != nil {
		return serializer.Response{
			Status: 400,
			Data:   nil,
			Msg:    "数据库添加phone错误",
			Error:  err,
		}
	}
	//获取该用户信息
	var user model.User
	if err := model.DB.First(&user).Where("id = ?", id).Error; err != nil {
		return serializer.Response{
			Status: 400,
			Data:   nil,
			Msg:    "数据库错误",
			Error:  err,
		}
	}
	//返回用户信息
	return serializer.Response{
		Status: 200,
		Data:   serializer.BuildUser(user),
		Msg:    "success",
	}
}

// SendMsg 发送短信验证码
func (service *CodeService) SendMsg() serializer.Response {
	rand.Seed(time.Now().UnixNano())    //设置随机种子
	codeInt := rand.Intn(6000)          //随机生成验证码
	codeString := strconv.Itoa(codeInt) //将code转成string
	temp := service.Phone + "code"
	err := cache.RedisClient.Set(temp, codeString, 0) //保存到redis，有效期10分钟
	if err != nil {
		logging.Info(err)
	}
	//初始化 简易验证
	client := unims.NewClient(conf.AccessKey)
	//构建信息
	message := unims.BuildMessage()
	message.SetTo(service.Phone)
	message.SetSignature(conf.Signature)
	message.SetTemplateId(conf.TemplateID)
	message.SetTemplateData(map[string]string{
		"code": codeString,
		"ttl":  "10",
	}) // 设置自定义参数 (变量短信)
	//发送短信
	_, err2 := client.Send(message)
	if err2 != nil {
		return serializer.Response{
			Status: 400,
			Data:   nil,
			Msg:    "发送验证码时出错",
			Error:  err2,
		}
	}
	return serializer.Response{
		Status: 200,
		Data:   nil,
		Msg:    "验证码发送成功",
	}
}
