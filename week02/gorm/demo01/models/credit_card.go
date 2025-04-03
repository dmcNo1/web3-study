package models

import "gorm.io/gorm"

type CreditCard struct {
	gorm.Model
	Number string `gorm:"default:00000"` // default设置默认值
	UserId uint
}

func (c *CreditCard) TableName() string {
	return "credit_card"
}
