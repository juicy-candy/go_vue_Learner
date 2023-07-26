package database

import (
	"fmt"
	"ginvue/pkg/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB = InitDB()

func GetDB() *gorm.DB {
	return DB
}

func InitDB() *gorm.DB {
	host := "localhost"
	port := "3306"
	database := "gogin2"
	username := "goGin"
	passwd := "123456"
	charset := "utf8mb4"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username,
		passwd,
		host,
		port,
		database,
		charset,
	)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic("Failed to connect database, err:" + err.Error())
	} else {
		fmt.Printf("数据库启动成功!")
	}
	db.AutoMigrate(&model.User{})

	return db
}
