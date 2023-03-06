package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"musicMod/utils"
)

var db *gorm.DB
var err error

func InitDb() {
	Connect()
	args := fmt.Sprintf("%s:%s@%s+tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		utils.DbUser,
		utils.DbPassWord,
		utils.Db,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	)
	fmt.Println(args)
	db, err = gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		panic("failed open mysql")
	}
	if err != nil {
		fmt.Println("连接失败", err)
		return
	}

	err = db.AutoMigrate(&User{}, &Music{}, &Usercols{}, &Userlikes{})
	if err != nil {
		return
	}

}
