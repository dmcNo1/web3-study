# GORM

参考的笔记文档东西比较啰嗦，这里也没有写的很细，尤其是多表查询这块，没有怎么写，还有事务控制这块也没有看过，最好自己看官方的文档。https://gorm.io/docs/

GORM是GoLang中提供的一个ORM框架，类似于Java的JDBC+MyBatis。GORM官方支持的数据库类型有：MySQL、PostgreSQL、SQlite、SQL Server，这里通过MySql来演示。

## 安装GORM

```sh
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

## 建立连接

创建`models/core.go`

```go
package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 为了方便全局使用，声明成全局变量
var DB *gorm.DB
var Err error

func Init() {
    // 连接信息，格式：user:password@/dbname?charset=utf8&parseTime=True&loc=Local
	dsn := "root:mysql@tcp(127.0.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
    // 开启数据库连接
	DB, Err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
```

## 定义操作数据库的模型

也就是创建Java中的entity/bean实体类，GoLang中通过定义对应的结构体来实现。也可以通过gorm来定义字段相关信息，详细可以参考：

* [topgoer](https://topgoer.cn/docs/gorm/gormmoxingdingyi)
* [GORM官方文档](https://gorm.io/zh_CN/docs/models.html)

profile belongs to permission

# 自己看文档和代码吧，笔记写不清楚