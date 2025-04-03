package models

import "gorm.io/gorm"

// profile属于permission，permission_id是外键
// foreignKey可以指定外键，默认就是"model名_id"
// 	在一对一的情况下，不过这种方式建表的话，会自动创建外键关联
//	在一对多的情况下，指向的是关联模型中的字段
// references可以指定关联外键，默认就是model的id；在GORM V1中使用的是association_foreignkey，这个在V2中被弃用了
type Profile struct {
	gorm.Model
	PermissionId       int
	Permission         Permission           `gorm:"foreignKey:PermissionId;references:ID"`
	PermissionProfiles []*PermissionProfile `gorm:"foreignKey:ProfileId`
}

func (p *Profile) TableName() string {
	return "profile"
}
