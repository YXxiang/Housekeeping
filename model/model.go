/**
 * @Author: LOFTER
 * @Description:
 * @File:  model
 * @Date: 2021/1/5 4:08 下午
 */
package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"os"
)

var DB *gorm.DB

//Initialized  连接mysql
func Initialized() {
	// 从本地读取环境变量
	DSN := os.Getenv("MYSQL_DSN")
	var err error
	DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		log.Fatalf("database connection error: %v", err)
		return
	}
	_ = DB.AutoMigrate(&Banner{}, &Goods{}, &King{}, &Review{})

}
