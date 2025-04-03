package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

// 定义一个结构体user，这个结构体对应的表名默认是users
type User struct {
	// gorm.Model	// Model中已经定义了一些常用的字段，可以直接引入；这里为了demo效果先不使用
	Id         uint       `gorm:"primary_key"`          // 设置这个字段为主键，默认主键就是id
	Name       string     `gorm:"column:name;size:100"` // column:指定列名，如果不指定，默认就是下划线分割形式；size:指定列长度
	IgnoreMe   int        `gorm:"-"`                    //忽略这个字段
	CreatedAt  time.Time  // 更新时间，如果这个字段为零值，则使用当前的时间
	UpdatedAt  int        `gorm:"autoUpdateTime:nano"` // 想要用时间戳的话，使用int类型即可；可以通过autoUpdateTime、autoCreateTime来制定是毫秒还是纳秒
	CreditCard CreditCard // 关联数据，只要这个对象不为空，那么在插入数据的时候，也会同步往credit_card表插入一条数据
}

// 获取表名的方法，如果不写这个方法，默认就是使用结构体的复数形式，users
func (u *User) TableName() string {
	return "user"
}

// 钩子函数，在插入前执行；可以在这里给插入的对象一些属性赋值
func (u *User) BeforeCreate(tx *gorm.DB) error {
	fmt.Println("before hook invoke")
	return nil
}
