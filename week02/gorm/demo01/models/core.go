package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Err error

func init() {
	// 连接信息，格式：user:password@/dbname?charset=utf8&parseTime=True&loc=Local
	dsn := "root:mysql@tcp(127.0.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
	// 开启数据库连接
	DB, Err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		CreateBatchSize: 1000, // 默认的插入批次大小
	})
}
