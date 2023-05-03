package util

import (
	//"github.com/jordan-wright/email"
	//"net/smtp"
	gomail "gopkg.in/gomail.v2"
	
)

//发送通知
func SendMail(text string , mail string) error{
	// e := gomail.NewMessage()
	// //设置发送方的邮箱
	// e.From = "1302997173@qq.com"
	// // 设置接收方的邮箱
	// e.To = []string{mail}
	// //设置主题
	// e.Subject = "充电桩预约系统"
	// //设置文件发送的内容
	// e.Text = text
	// //设置服务器相关的配置
	// err := e.Send("smtp.qq.com:25", smtp.PlainAuth("",
	// 	"1302997173@qq.com", "kmibjpfvgamxfggh", "smtp.qq.com"))
	m := gomail.NewMessage()
	//发送人
	m.SetHeader("From", "1302997173@qq.com")
	//接收人
	m.SetHeader("To", mail)
	//抄送人
	//m.SetAddressHeader("Cc", "xxx@qq.com", "xiaozhujiao")
	//主题
	m.SetHeader("Subject", "充电桩预约系统")
	//内容
	m.SetBody("text/html",text)
	//附件
	//m.Attach("./myIpPic.png")
 
	//拿到token，并进行连接,第4个参数是填授权码
	d := gomail.NewDialer("smtp.qq.com", 587, "1302997173@qq.com", "kmibjpfvgamxfggh")
 
	// 发送邮件
	if err := d.DialAndSend(m); err != nil {
		return err
	}
    return nil
	
	//kmibjpfvgamxfggh
}
