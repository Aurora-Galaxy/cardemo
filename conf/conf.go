package conf

import (
	"car/model"
	"fmt"
	"gopkg.in/ini.v1"
	"strings"
)

var (
	//MongoDBClient 		*mongo.Client
	AppMode    string
	HttpPort   string
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	//MongoDBName    		string
	//MongoDBAddr  		string
	//MongoDBPwd    		string
	//MongoDBPort    		string

	AppID  string
	Secret string

	QiNiuAccessKey string
	QNSerectKey    string
	Bucket         string
	QiNiuServer    string

	ClientId    string
	BDSecretKey string

	//UniSMS

	AccessKey  string
	Signature  string
	TemplateID string
)

func Init() {
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Println("配置文件加载错误，请检查文件路径:", err)
	}
	//读取文件中各个部分的具体内容
	LoadSever(file)
	LoadMysqlData(file)
	LoadBaiDuService(file)
	LoadUniSMS(file)
	//拼接Mysql路径
	path := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8&parseTime=true"}, "")
	model.Database(path)
}

func LoadSever(file *ini.File) {
	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
}

func LoadMysqlData(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
}

func LoadQiNiuData(file *ini.File) {
	QiNiuAccessKey = file.Section("qiniu").Key("AccessKey").String()
	QNSerectKey = file.Section("qiniu").Key("SerectKey").String()
	Bucket = file.Section("qiniu").Key("Bucket").String()
	QiNiuServer = file.Section("qiniu").Key("QiniuServer").String()
}

func LoadWxChat(file *ini.File) {
	AppID = file.Section("wechat").Key("APPID").String()
	Secret = file.Section("wechat").Key("SECRET").String()
}

func LoadBaiDuService(file *ini.File) {
	ClientId = file.Section("BaiDuOCR").Key("ClientId").String()
	BDSecretKey = file.Section("BaiDuOCR").Key("SecretKey").String()
}

func LoadUniSMS(file *ini.File) {
	AccessKey = file.Section("UniSMS").Key("AccessKey").String()
	Signature = file.Section("UniSMS").Key("Signature").String()
	TemplateID = file.Section("UniSMS").Key("TemplateID").String()
}
