package database

import (
	"fmt"
	"ginvue/pkg/model"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetDB() *gorm.DB {
	return DB
}

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		viper.Get("database.username"),
		viper.Get("database.passwd"),
		viper.Get("database.host"),
		viper.Get("database.port"),
		viper.Get("database.db"),
		viper.Get("database.charset"),
	)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Print("连接至数据库失败\n")
		panic("Failed to connect database, err :" + err.Error())
	} else {
		log.Print("数据库启动成功")
	}
	DB = db
	db.AutoMigrate(&model.User{})

	DB = db
}
