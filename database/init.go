package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 数据库初始化
var (
	DB = DBInit() // var *gorm.DB
)

func DBInit() *gorm.DB {
	// mysql登录配置
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True&loc=Local", DB_username, DB_password, DB_ip, DB_port, DB_dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to the mysql")

	return db
}
