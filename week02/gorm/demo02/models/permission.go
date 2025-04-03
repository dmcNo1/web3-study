package models

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	Name  string
	Roles []*Role `gorm:"many2many:role_permission;joinForeignKey:PermissionId;joinReferences:RoleId"`
}

func (p *Permission) TableName() string {
	return "permission"
}
