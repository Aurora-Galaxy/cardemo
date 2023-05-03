package service

import (
	"car/model"
	//"car/pkg/e"
	"car/pkg/util"
	"car/serializer"
	"github.com/jinzhu/gorm"
	//logging "github.com/sirupsen/logrus"
)

type AdminService struct {
	Username string `form:"user_name" json:"user_name"`
	PassWord string `form:"password" json:"password"`
}

type UserDeleteService struct{
}

// AdminRegister Register 用户注册
// func (adminService *AdminService) AdminRegister() serializer.Response {
// 	var admin model.Admin
// 	var count int
// 	model.DB.Model(&model.Admin{}).Where("user_name=?", adminService.Username).
// 		First(&admin).Count(&count) //如果count == 1 则证明数据库中存在这个人
// 	if count == 1 {
// 		return serializer.Response{
// 			Status: e.ErrorExistUser,
// 			Data:   nil,
// 			Msg:    "该用户名已经存在",
// 		}
// 	}
// 	admin.UserName = adminService.Username
// 	//密码加密
// 	err := admin.SetPassword(adminService.PassWord)
// 	if err != nil {
// 		return serializer.Response{
// 			Status: 400,
// 			Data:   nil,
// 			Msg:    "密码加密时错误",
// 			Error:  err,
// 		}
// 	}
// 	//创建用户
// 	if err = model.DB.Create(&admin).Error; err != nil {
// 		logging.Println(err)
// 		return serializer.Response{
// 			Status: 500,
// 			Data:   nil,
// 			Msg:    "创建用户时错误",
// 			Error:  err,
// 		}
// 	}
// 	return serializer.Response{
// 		Status: 200,
// 		Data:   nil,
// 		Msg:    "创建用户成功",
// 	}
// }

// AdminLogin  用户登录
func (adminService *AdminService) AdminLogin() serializer.Response {
	var admin model.Admin
	err := model.DB.Table("admin").Where("user_name=?", adminService.Username).First(&admin).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) { //没有查询到该用户
			return serializer.Response{
				Status: 400,
				Data:   nil,
				Msg:    "用户不存在，请先注册",
			}
		}
		return serializer.Response{
			Status: 500,
			Data:   nil,
			Msg:    "数据库错误",
			Error:  err,
		}
	}
	if adminService.PassWord != "admin" {
		return serializer.Response{
			Status: 400,
			Data:   nil,
			Msg:    "密码错误",
		}
	}
	//签发token
	token, err := util.GenerateToken(admin.ID, adminService.Username, adminService.PassWord)
	if err != nil {
		return serializer.Response{
			Status: 500,
			Data:   nil,
			Msg:    "token签发错误",
			Error:  err,
		}
	}
	return serializer.Response{
		Status: 200,
		Data: serializer.TokenData{
			//User:  serializer.BuildUser(user),
			Token: token,
		},
		Msg: "登录成功",
	}
}


//删除用户
func(service *UserDeleteService) DeleteUser(id string) serializer.Response {
	var user model.User
	err := model.DB.Table("user").Where("id = ?",id).First(&user).Error
	if err != nil {
		return serializer.Response{
			Status: 400,
			Data:   nil,
			Msg:    "数据库查找错误",
			Error:  err,
	   }
    } 
	err = model.DB.Table("user").Delete(&user).Error
	if err != nil{
		return serializer.Response{
			Status: 400,
			Data:   nil,
			Msg:    "数据库删除错误",
			Error:  err,
	    }
	}
	return serializer.Response{
		Status: 200,
		Data:   nil,
		Msg:    "删除用户成功",
		Error:  nil,
	}
}