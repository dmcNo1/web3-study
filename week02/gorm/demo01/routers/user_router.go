package routers

import (
	"demo01/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitUserRouters(router *gin.Engine) {
	db := models.DB
	userRouter := router.Group("/user")
	{
		// 根据结构体的定义，创建表
		userRouter.GET("/create_table", func(ctx *gin.Context) {
			db.AutoMigrate(&models.User{})
		})

		// 增加记录
		userRouter.GET("/create", func(ctx *gin.Context) {
			user := &models.User{CreditCard: models.CreditCard{}}
			// Table()方法可以指定表名，否则用结构体定义的或者默认的表名
			// Select()方法指定字段
			// Omit()指定忽略的字段
			// Create()，传参必须是指针，也可以是一个map
			// result := db.Table("user").Select("Id", "Name", "CreatedAt").Omit("IgnoreMe").Create(user)
			result := db.Create(user)
			ctx.JSON(http.StatusOK, gin.H{"返回的睡觉": user, "error": result.Error, "返回插入记录的条数": result.RowsAffected})
		})

		// 批量增加
		userRouter.GET("/create_in_batches", func(ctx *gin.Context) {
			users := []*models.User{{Name: "zhangsan"}, {Name: "lisi"}}
			// result := db.Create(users)	// 这样会将数据分批次插入
			// CreateInBatches()，指定插入的批次大小插入
			result := db.CreateInBatches(users, 100)
			ctx.JSON(http.StatusOK, gin.H{"返回的数据": users, "error": result.Error, "返回插入记录的条数": result.RowsAffected})
		})

		// 查询所有记录
		userRouter.GET("/find", func(ctx *gin.Context) {
			users := make([]models.User, 0)
			// Find()，查询所有结果，封装到对应的切片中；也可以指定参数，通过id in(?, ?)的方式查询
			// Where()可以指定where条件，也可以传入Map或者对象指针；
			// 	1、Where(where string)
			//  2、Where(&model.User{...}, "Id", "Name")；
			// 		如果不指定需要的字段，则会全量采用结构体中的非零值数据
			//		假如Id字段是0，如果后面不指明字段的话，Id就不会拼接到where中；
			//		也可以将结构体的属性指定为指针类型，或者sql.NullInt64的方式来让零值生效
			//  3、Where(map)
			//	如果穿了多个Where()，多个条件之间通过and拼接
			// Not()，其中的条件活用not xxx的形式体现在SQL中
			// Or()，or拼接条件
			// Group()
			// Having()
			// select * from user where id = 1 and name in ('zhangsan', 'lisi')
			// db.Where("id=?", 1).Where("name in (?)", []string{"zhangsan", "lisi"}).Find(&users)
			db.Where(&models.User{Id: 1, Name: "lisi"}, "Id").Find(&users)
			fmt.Println(users)
		})

		// 查询单条记录
		userRouter.GET("/find_first", func(ctx *gin.Context) {
			user := models.User{}
			// First()查询第一条记录，默认按照主键排序；如果没有定义主键，会按照model的第一个字段排序
			// select * from user order by id limit 1,1
			//	1、也可以指定参数，通过id = ?的方式查询
			//	2、如果传入的是对象指针，那么会查询全表并且返回第一个对象，效率非常低
			// Last()，返回最后一个，select * from user order by id desc limit 1,1
			// Take()，返回第一个，但是没有排序，select * from user limit 1,1
			// Set()，为SQL设置额外的选项；select * from user where id = 10 for update
			// FirstOrInit()，获取第一条匹配的记录，或者通过给定的条件下初始一条新的记录（仅适用与于 struct 和 map 条件）。
			//		FirstOrInit(&user, models.User{Name: "not exists user"})，当查询不到时，就会返回Name为"not exists user"的数据
			// db.Set("gorm:query_option", "FOR UPDATE").First(&user, 10)
			db.First(&user)
			ctx.JSON(http.StatusOK, user)
		})

		// 查询单条记录，通过map接收
		userRouter.GET("/find_first_map", func(ctx *gin.Context) {
			m := map[string]any{}
			// Table()指定查询的表
			// 这里只能用Take，不能用First()和Last()
			db.Table("user").Take(m)
			ctx.JSON(http.StatusOK, m)
		})

		// 查询
		userRouter.GET("/find_by_model", func(ctx *gin.Context) {
			user := models.User{}
			// 把条件封装到Model()中
			db.Model(models.User{Id: 1}).First(&user)
			ctx.JSON(http.StatusOK, user)
		})

		userRouter.GET("/find_join", func(ctx *gin.Context) {
			users := []*models.User{}
			// Joins()，可以在里头拼接join语句
			db.Model(&models.User{}).Joins("left join credit_card on user.id = credit_card.id").Find(&users)
			ctx.JSON(http.StatusOK, users)
		})

		// 更新-Save
		userRouter.GET("/save", func(ctx *gin.Context) {
			var user models.User
			db.First(&user)
			user.Name = "jackpot"
			// Save()，根据主键更新全部字段
			db.Save(&user)
		})

		// 更新
		userRouter.GET("/update", func(ctx *gin.Context) {
			user := models.User{Id: 1}
			// Update()，只更新指定的字段，
			// Updates()，可以传入map、struct，如果用struct的话，不会更新零值
			db.Model(&user).Update("name", "update")
		})

		// 删除
		userRouter.GET("/delete", func(ctx *gin.Context) {
			user := models.User{Id: 2}
			// Delete()，删除；如果model中有DeleteAt的话，会自动执行软删除
			db.Delete(&user)
		})
	}
}
