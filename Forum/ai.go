package main

/*
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"golang.org/x/crypto/bcrypt" // 密码哈希加密库
*/

// 用户结构体，包含用户名和密码
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 用户列表，初始有三个用户，可以通过注册添加新用户
var users = []User{{"user1", "pass1"}, {"user2", "pass2"}, {"user3", "pass3"}}

/*
func main() {
	app := iris.New() // 创建一个新的Iris应用

	// 添加中间件
	app.Use(recover.New()) // 恢复中间件，用于处理panic
	app.Use(logger.New())  // 日志中间件，记录请求和响应

	// 注册路由
	app.Post("/register", func(ctx iris.Context) {
		var newUser User
		if err := ctx.ReadJSON(&newUser); err != nil { // 从请求中读取JSON格式的用户数据
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.WriteString(err.Error())
			return
		}

		if len(newUser.Password) < 8 { // 检查密码是否过于简单
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.WriteString("Password is too weak")
			return
		}

		// 对密码进行哈希加密
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.WriteString(err.Error())
			return
		}

		newUser.Password = string(hashedPassword) // 将加密后的密码存入用户结构体
		users = append(users, newUser)            // 将新用户添加到用户列表中
		ctx.StatusCode(iris.StatusCreated)        // 返回状态码201
		ctx.JSON(newUser)                         // 返回JSON格式的用户数据
	})

	app.Post("/login", func(ctx iris.Context) {
		var loginUser User
		if err := ctx.ReadJSON(&loginUser); err != nil { // 从请求中读取JSON格式的用户数据
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.WriteString(err.Error())
			return
		}

		for _, user := range users { // 遍历用户列表
			if user.Username == loginUser.Username { // 如果找到该用户
				err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password)) // 对密码进行哈希比较
				if err != nil {                                                                         // 如果密码不匹配
					ctx.StatusCode(iris.StatusUnauthorized)
					ctx.WriteString("Incorrect password")
					return
				}

				ctx.StatusCode(iris.StatusOK)
				ctx.WriteString("Logged in successfully")
				return
			}
		}

		ctx.StatusCode(iris.StatusNotFound)
		ctx.WriteString("User not found")
	})

	app.Post("/logout", func(ctx iris.Context) {
		ctx.StatusCode(iris.StatusOK)
		ctx.WriteString("Logged out successfully")
	})

	app.Post("/forgot-password", func(ctx iris.Context) {
		var forgotUser User
		if err := ctx.ReadJSON(&forgotUser); err != nil { // 从请求中读取JSON格式的用户数据
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.WriteString(err.Error())
			return
		}

		for _, user := range users { // 遍历用户列表
			if user.Username == forgotUser.Username { // 如果找到该用户
				hashedPassword, err := bcrypt.GenerateFromPassword([]byte(forgotUser.Password), bcrypt.DefaultCost) // 对新密码进行哈希加密
				if err != nil {
					ctx.StatusCode(iris.StatusInternalServerError)
					ctx.WriteString(err.Error())
					return
				}

				user.Password = string(hashedPassword) // 将加密后的新密码存入用户结构体
				ctx.StatusCode(iris.StatusOK)
				ctx.WriteString("Password reset successfully")
				return
			}
		}

		ctx.StatusCode(iris.StatusNotFound)
		ctx.WriteString("User not found")
	})

	// 要求访问该页面的用户必须先登录，否则返回401状态码
	app.Get("/secured", func(ctx iris.Context) {
		if ctx.GetCookie("auth") == "" { // 检查是否有授权cookie
			ctx.StatusCode(iris.StatusUnauthorized)
			ctx.WriteString("Not authorized")
			return
		}

		ctx.StatusCode(iris.StatusOK)
		ctx.WriteString("Secure content")
	})

	// 登录路由
	app.Post("/login", func(ctx iris.Context) {
		var loginUser User
		if err := ctx.ReadJSON(&loginUser); err != nil { // 从请求中读取JSON格式的用户数据
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.WriteString(err.Error())
			return
		}

		for _, user := range users { // 遍历用户列表
			if user.Username == loginUser.Username { // 如果找到该用户
				err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password)) // 对密码进行哈希比较
				if err != nil {                                                                         // 如果密码不匹配
					ctx.StatusCode(iris.StatusUnauthorized)
					ctx.WriteString("Incorrect password")
					return
				}

				ctx.SetCookieKV("auth", "true", iris.CookieExpires(time.Hour*24)) // 添加授权cookie，有效期为24小时
				ctx.StatusCode(iris.StatusOK)
				ctx.WriteString("Logged in successfully")
				return
			}
		}

		ctx.StatusCode(iris.StatusNotFound)
		ctx.WriteString("User not found")
	})

	// 登出路由，删除授权cookie
	app.Post("/logout", func(ctx iris.Context) {
		ctx.RemoveCookie("auth")
		ctx.StatusCode(iris.StatusOK)
		ctx.WriteString("Logged out successfully")
	})

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed)) // 启动应用，监听8080端口
}
*/
