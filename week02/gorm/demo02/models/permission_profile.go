package models

import "gorm.io/gorm"

type PermissionProfile struct {
	gorm.Model
	Name      string
	ProfileId uint
}

func (p *PermissionProfile) TableName() string {
	return "permission_profile"
}
