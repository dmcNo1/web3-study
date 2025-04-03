package models

import "gorm.io/gorm"

// 角色
// @many2many
// 用于指定中间关联表，默认是通过一个模型_id和另一个模型_id关联
// 	select *
// 	from role
// 	join role_permission on role.id = role_permission.role_id
// 	join permission on role_permission.permission_id = permission.id
// 如果想要一对多，可以只设置一个模型的many2many
// 如果想要多对多，两个模型都需要设置many2many
//
// @joinForeignKey、joinReferences
// 指定关联查询的字段，默认就是role_id和permission_id，可以根据需求重写
type Role struct {
	gorm.Model
	Name        string
	Permissions []*Permission `gorm:"many2many:role_permission;joinForeignKey:RoleId;joinReferences:PermissionId"`
}

func (r *Role) TableName() string {
	return "role"
}
