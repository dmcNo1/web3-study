package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	dsn := "root:mysql@tcp(127.0.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
	DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{
		CreateBatchSize:                          1000, // 默认的插入批次大小
		DisableForeignKeyConstraintWhenMigrating: true, // 禁止自动生成外键
	})
}
