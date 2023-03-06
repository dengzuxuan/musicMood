package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string
	JwtKey   string

	Db         string
	DbPort     string
	DbHost     string
	DbPassWord string
	DbName     string
	DbUser     string

	AccessKey   string
	SecretKey   string
	Bucket      string
	QiniuServer string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件错误", err)
		return
	}
	//LoadServer(file)
	LoadData(file)
	//LoadQiniu(file)
}
func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":8887")
	JwtKey = file.Section("server").Key("jwtKey").MustString("adsdwreefsdanjdq")
}
func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("123456")
	DbName = file.Section("database").Key("DbName").MustString("toysql2")
	DbUser = file.Section("database").Key("DbUser").MustString("toysql2")
}
func LoadQiniu(file *ini.File) {
	AccessKey = file.Section("qiniu").Key("AccessKey").MustString("97OSddKvoSFShgtTJ4MVpBkPROimA1l3xN4KNyFK")
	SecretKey = file.Section("qiniu").Key("SecretKey").MustString("v7xyZ8rjRaCA9SIAhf-wAk2PnCeE2qpSxqcYh3cu")
	Bucket = file.Section("Bucket").Key("Bucket").MustString("toyproject0204")
	QiniuServer = file.Section("QiniuServer").Key("QiniuServer").MustString("http://r9w6vft1j.hb-bkt.clouddn.com/")
}
