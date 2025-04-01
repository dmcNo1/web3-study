# Gin

https://gin-gonic.com/

https://github.com/gin-gonic/gin

https://gin-gonic.com/zh-cn/docs/examples/upload-file/single-file/

## 环境搭建

在VisualCode中，在工作区完成了`go mod init mode_name`之后，通过**`go get github.com/gin-gonic/gin`**即可，如果是在cmd窗口，则需要使用`go install`。这样就可以使用gin框架了，可以配合`net/http`包使用已经封装好了的常用的httpStatus。

demo如下：

```go
func main() {
	// 配置一个默认的路由
	router := gin.Default()
	// 绑定路由
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello World!")
	})

	// 启动一个Web服务，可以通过router.Run([address]:port)指定端口，router.Run("127.0.0.1:8081")
	router.Run(":8081")
}
```

## 热加载

https://github.com/gravityblast/fresh

```shell
# 在项目的路径下，用这个执行fresh
go get github.com/pilu/fresh
go run github.com/pilu/fresh
```

## 路由

最简单的路由配置，以及数据处理方式如下

```go
// 配置路由信息，/news会调用指定的方法
router.GET("/news", func(ctx *gin.Context) {
	// 获取参数
	name := ctx.Query("name")
	page := ctx.DefaultQuery("page", "1")
	fmt.Printf("name = %v, page = %v\n", name, page)
    // 返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "success",
	})
})
```

## 路由分组

像上面那样配置路由的话，代码的可读性很低，可以会像Java一样，把具体的业务处理，抽取到xxxController中。比如，先定义一个`controller/user/user_controller.go`文件，然后在文件里定义一个`UserController`结构体，给这个结构体实现具体的业务逻辑。然后将路由信息配置到一起，并且可以根据不同的路由路径，进行分组。最后在`main.go`中，调用对应的路由初始化方法，初始化路由信息。不过GoLang的项目结构不全都像Spring那样，很多项目都不是`controller->service->mapper`的结构（MVC），根据这里只是举一个例子，方便学习。代码如下：

`main.go`

```go
func main() {
	// 配置一个默认的路由
	router := gin.Default()
	// 绑定路由
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello World!")
	})

	// 初始化路由
	routers.InitRouter(router)

	// 启动一个Web服务，可以通过router.Run([address]:port)指定端口，router.Run("127.0.0.1:8081")
	router.Run(":8081")
}
```

`router_init.go`

```go
// 初始化路由
func InitRouter(router *gin.Engine) {
	initUserRouters(router)
}
```

`user_router.go`

```go
func initUserRouters(router *gin.Engine) {
	// 创建一个路由分组，所有/user的路由都会匹配到这里头
	userRouters := router.Group("/user")
	{
		// 绑定路由到对应的controller，不同的项目不一样，很多项目都不是mvc架构的
		userController := user.UserController{}
		userRouters.GET("/user/:name/*action", userController.Action)
		userRouters.POST("/form", userController.Form)
	}
}
```

`user_controller.go`

```go
type UserController struct {
}

// 获取参数，localhost:8081/user/zhagnsan/coding?level=slowly
func (u *UserController) Action(ctx *gin.Context) {
	// Restful风格可以通过func (c *Context) Param(key string) string获取
	name := ctx.Param("name")
	action := strings.Trim(ctx.Param("action"), "/")
	// 获取Get中的传参，可以用Query()或者DefaultQuery()
	level := ctx.DefaultQuery("level", "fast")
	ctx.String(http.StatusOK, fmt.Sprintf("%s is %s %s", name, action, level))
}

// 获取表单中的参数
func (u *UserController) Form(ctx *gin.Context) {
	// DefaultPostForm()或者PostForm()获取表单中的数据
	method := ctx.DefaultPostForm("method", "post")
	username := ctx.PostForm("username")
	password := ctx.PostForm("userpassword")
	ctx.String(http.StatusOK, fmt.Sprintf("method:%s --- username:%s password:%s", method, username, password))
}
```

## 参数绑定

GoLang会根据请求头中content-type自动推断传参的类型，然后根据结构体的配置绑定参数

```go
// 传参对应的结构体，对应Param
// 可以在"“"之间，定义不同格式下该字段对应的参数，也可以配置是否必填
type Login struct {
	Username string `form:"un" json:"username" uri:"username" xml:"username" binding:"required"`
	Pwd      string `form:"pwd" json:"password" uri:"password" xml:"password" binding:"required"`
}
```

可以通过不同的方法，将类型以不同的方式绑定到结构体对象上

```go
type AuthController struct{}

func (a *AuthController) LoginJson(ctx *gin.Context) {
	var login Login
	// 按照JSON格式绑定到对象
	if err := ctx.ShouldBindJSON(&login); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, login)
}

func (a *AuthController) LoginForm(ctx *gin.Context) {
	var login Login
	// 按照JSON格式绑定到对象
	if err := ctx.ShouldBind(&login); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, login)
}

// 测试用uri：localhost:8081/auth/loginUri/admin/888888
func (a *AuthController) LoginUri(ctx *gin.Context) {
	var login Login
	// 按照JSON格式绑定到对象
	if err := ctx.ShouldBindUri(&login); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, login)
}
```

## 中间件

gin的中间件，必须是`gin.HandlerFunc`类型，这是这个类型的定义`type HandlerFunc func(*Context)`。配置路由的时候可以传递多个func回调函数，最后一个func回调函数前面触发的方法都可以称为中间件。可以这样配置中间件：

```go
userRouters.GET("/index", middlewares.LogMiddleware, middlewares.GoRoutineMiddleware, userController.Index)
```

在创建路由的时候，通常是使用`gin.Default()`，这样创建的路由，默认使用了Logger和Recovery中间件。

实现一个中间件：

```go
// 创建一个中间件，模拟记录方法执行实现的日志中间件
// 以“/user/index”为例，实际执行的顺序就是
// 1、计算startTime
// 2、执行具体的业务逻辑（也就是调用userController.Index()）/调用下一个中间件，
// 3、计算endTime...
// 如果配置了多个中间件的话，把ctx.Next()方法前的视为before，后面的视为after，执行顺序就是a.before -> b.before -> 业务逻辑 -> b.after -> a.after
func LogMiddleware(ctx *gin.Context) {
	startTime := time.Now().UnixNano()

	// 获取其他中间件设置的数据
	username, existsFlag := ctx.Get("username")
	if existsFlag {
		fmt.Printf("logging -- username = %v\n", username)
	}

	// 这个方法表示可以执行剩下的业务逻辑了，也就是具体的业务逻辑执行方法
	ctx.Next()

	endTime := time.Now().UnixNano()

	fmt.Printf("logging -- startTime = %v, endTime = %v, coastNanos = %v\n", startTime, endTime, endTime-startTime)
}
```

如果一个个的请求都这样配置的话，那太麻烦了，可以在全局路由，或者路由分组下，定义全局的中间件，通过`router.Use()`来设置：

```go
// 配置一个全局中间件，一样可以配置多个；func (engine *Engine) Use(middleware ...HandlerFunc) IRoutes {}
defaultRouter.Use(globalMiddleware)

// 也可以在这里为路由分组指定中间件UserRouters := router.Group("/user", userGroupMiddleware)
userRouters := router.Group("/user")
userRouters.Use(middlewares.UserGroupMiddleware)
```

### 中间件的数据共享

不同的中间件之间，可以实现数据的共享，包括controller也可以获取到数据，只需要对`gin.Context`进行操作：

```go
// 设置共享数据，其他中间件也能访问到
ctx.Set("username", "username")
// 获取其他中间件设置的数据
username, existsFlag := ctx.Get("username")
```

### 在中间件中开启协程

当需要在中间件中使用goroutine的时候，**不能使用原始的上下文（\*gin.Context）**，必须使用其只读副本，`ctx.Copy()`可以创建一个副本。

```go
// 在中间件中，开启一个协程
func GoRoutineMiddleware(ctx *gin.Context) {
	// 必须使用副本，不然会出问题。
	cCp := ctx.Copy()
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Done! In path", cCp.Request.URL.Path)
	}()
}
```

## Cookie

https://gin-gonic.com/zh-cn/docs/examples/cookie/

可以通过`ctx.SetCookie()`来存放Cookie，`ctx.Cookie()`来获取Cookie，将Cookie过期时间设为-1来删除

```go
// 设置Cookie
func (u UserController) SetCookie(ctx *gin.Context) {
	// func (c *Context) SetCookie(name, value string, maxAge int, path, domain string, secure, httpOnly bool)
	// name：key
	// value：value
	// maxAge：过期时间，单位为秒
	// path：有效路径，只有访问路径为/user才能访问这个Cookie
	// domain：域名，主要用于实现多个域名共享；比如两个网址，一个在192.168.100.1，一个在192.168.100.2，这两个都想使用就能这样
	// secure：当这个值为true时，只有在https中才能生效
	// httpOnly，微软对Cookie的扩展，防止XSS攻击
	ctx.SetCookie("name", "dick", 1000, "/user", "localhost", false, true)
	u.Success(ctx)
}

// 获取Cookie
func (u UserController) GetCookie(ctx *gin.Context) {
	if name, err := ctx.Cookie("name"); err == nil {
		ctx.String(http.StatusOK, "name = "+name)
	} else {
		u.Error(ctx)
	}
}

func (u UserController) DeleteCookie(ctx *gin.Context) {
	// 删除Cookie，或者把值设为""，也是可以实现删除
	ctx.SetCookie("name", "", -1, "/", "localhost", false, true)
}
```

## Session

Gin官方没有提供Session相关的文档，可以通过这个第三方插件来实现：https://github.com/gin-contrib/sessions。`gin-contrib/sessions`中间件支持的存储引擎：

* cookie 
* memstore
* redis
* memcached
* mongodb

这里先通过根据Cookie作为存储引擎来介绍，先引入依赖：

```sh
go get github.com/gin-contrib/sessions
```

在`main.go`中配置session中间件：

```go
func main() {
	// 创建一个默认的路由
	defaultRouter := gin.Default()

	// 配置session中间件
	// 创建基于Cookie的存储引擎，security888888是用于加密的秘钥
	store := cookie.NewStore([]byte("security888888"))
	// 配置session的中间件，store是之前创建的存储引擎；当创建session的时候，会调用这个中间件
	defaultRouter.Use(sessions.Sessions("mySession", store))

	routers.UserRoutersInit(defaultRouter)

	defaultRouter.Run("127.0.0.1:8081")
}
```

这里是`controller`层实现设置和获取Session数据的业务逻辑：

```go
// 设置Session
func (u UserController) Index(ctx *gin.Context) {
	// 初始化session
	session := sessions.Default(ctx)
	// 设置过期时间
	session.Options(sessions.Options{
		MaxAge: 3600,
	})
	// 设置session
	session.Set("username", "Penis Hair")
	// 设置完之后要保存
	session.Save()
	u.Success(ctx)
}

// 获取Session数据
func (u UserController) GetSession(ctx *gin.Context) {
	session := sessions.Default(ctx)
	username := session.Get("username")
	ctx.String(http.StatusOK, "username = "+username.(string))
}
```

工作中常用的还有基于`redis`存储，只需要将存储引擎按照redis来配置即可，具体参考文档或者官网。



## 模板渲染、静态文件

直接看笔记吧，没学

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
    // 连接信息
	dsn := "root:mysql@tcp(127.0.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
    // 开启数据库连接
	DB, Err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
```

## 定义操作数据库的模型

也就是创建Java中的entity/bean实体类，GoLang中通过定义对应的结构体来实现；对应的结构体需要实现`TableName`方法，来获取对应的库表名称，如果不实现这个方法，默认是对应的结构体的复数形式，这里也就是`users`表，建议自己实现比较好。

```go
// 对应的实体
type User struct {
	Id       int
	Username string
	Age      int
	Email    string
	AddTime  int	// 对应字段add_time
}

// 获取表明的方法，一定要这样命名
func (u User) TableName() string {
	return "user"
}
```

## CRUD

### 查找

```go
// 将查找结果赋值到userList中
func (u UserController) Find(ctx *gin.Context) {
    userList := []models.User{}
    // 如果没有where条件，可以忽略Where()这个方法
    models.DB.Where(whereSql).Find(&userList)
}
```

### 插入

```go
func (u UserController) Add(ctx *gin.Context) {
	user := &models.User{}
	if err := ctx.ShouldBind(user); err == nil {
		// insert
		models.DB.Create(user)
	}
}
```

### 修改

```go
func (u UserController) Patch(ctx *gin.Context) {
	// 根据patch传来的参数，生成对应的entity，用于查询
	user := &models.User{}
	if id := ctx.Query("id"); id != "" {
		user.Id, _ = strconv.Atoi(id)
	}

	models.DB.Find(&user)
	if user != nil {
		if err := ctx.ShouldBind(user); err == nil {
			// update
			// 这样会把所有的字段都更新：models.DB.Save(user)，这样即使是0或者nil也会更新，这样用的最多
			// 更新单列：models.DB.Model(&user).Where("id = ?", user.Id).Update("email", user.Email)
			// 更新多列
			models.DB.Save(user)
		}
	}

}
```

### 删除

```go
func (u UserController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		u.Error(ctx)
		return
	}

	// delete，最常用的写法
    // 也可以这样：models.DB.Where("username = ?", user.Username).Delete(user)
	user := &models.User{Id: id}
	models.DB.Delete(user)
	
}
```

### 关联查询

这里比较麻烦，可以多看看官方文档。

首先，模型需要指定对应的业务外键：

```go
// 对应的实体
type User struct {
	Id       int
	Username string
	Age      int
	Email    string
	AddTime  int
	// 假如user表和role表是一对一关联，那么role默认会以role_id作为业务外键，也可以自定义对应的外键
	RoleId int
	Role   Role `gorm:"foreignKey:RoleId"`
}

// 获取表明的方法，一定要这样命名
func (u User) TableName() string {
	return "user"
}
```

然后对应的查询逻辑可以这样处理：

```go
// select * from user u join role r on u.role_id = r.id
func (u UserController) GetUserRole(ctx *gin.Context) {
	userList := []models.User{}
	// 通过Preload制定需要关联查询的对象，这里会关联查询Role对象对应的结构体所对应的表
	models.DB.Preload("Role").Find(&userList)
	ctx.JSON(http.StatusOK, gin.H{
		"result": userList,
	})
}
```

### 使用原生的SQL

也就是自己手写SQL

```go
// 通过原生SQL更新数据
db.Exec("update nav set url = ? where id = ?", "xxx.xxx.com", 1)
```

# 微服务

自己找资料

# 遇到的问题

## 工具安装失败，网络连接问题

网络问题，即使是挂了梯子也连不上。手动改下镜像源即可。

```shell
# 错误提示
2025-02-09 12:50:54.267 [info] goimports: failed to install goimports(golang.org/x/tools/cmd/goimports@latest): Error: Command failed: D:\Go\bin\go.exe install -v golang.org/x/tools/cmd/goimports@latest
go: golang.org/x/tools/cmd/goimports@latest: module golang.org/x/tools/cmd/goimports: Get "https://proxy.golang.org/golang.org/x/tools/cmd/goimports/@v/list": dial tcp 142.250.69.209:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.

# 解决方案：
# 这个命令可以查看当前的镜像配置
go env GOPROXY
# 结果：https://proxy.golang.org,direct
# 这个命令可以修改当前的镜像到阿里云
go env -w GOPROXY=https://mirrors.aliyun.com/goproxy,direct

# 或者用这个试试
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```

然后再执行安装命令即可

```shell
go install -v golang.org/x/tools/cmd/goimports@latest
```

不过，当VSCode自动安装的时候，会安装很多的插件，这些插件都会安装失败，要一次性安装，可以在配置了镜像之后，可以按下键盘上的`ctrl+shift+P` ，然后输入`go install/update tools`，把所有工具都选上，点击ok即可

## 热部署命令不生效
