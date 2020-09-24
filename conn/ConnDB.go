package conn

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB
const dsn string = "root:123456@(124.70.71.78:3306)/Stitches?charset=utf8mb4&parseTime=True&loc=Local"

func init() {
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(">>>>>>>>>>>>>数据库已连接<<<<<<<<<<<<<<<<<<")
}
