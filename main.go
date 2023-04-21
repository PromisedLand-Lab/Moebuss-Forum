package main

import (
	"Promisedland/Moebuss-Forum/tools"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"

	"github.com/kataras/iris/v12"
)

func main() {
	fmt.Println(tools.LoginCheck("Wu110228", "batrycc"))
	run()
}

func run() {
	app := iris.New()

	app.RegisterView(iris.HTML("./views", ".html"))

	app.HandleDir("/comic", "./comic")

	authMiddleware := func(ctx iris.Context) {
		// 假设用户已经登录，这里只做示例，具体实现方式根据您的需求而定
		userLoggedIn := tools.LoginCheck("Wu110228", "123")

		// 如果用户已经登录，则继续处理请求
		if userLoggedIn {
			ctx.Next()
			return
		}

		// 如果用户未登录，则返回 401 Unauthorized 错误
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.WriteString("401 Unauthorized")
	}

	app.Get("/", authMiddleware, func(ctx iris.Context) {
		comicFolders, err := ioutil.ReadDir("./comic")
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.WriteString("Error reading comic folders")
			return
		}

		var comics []map[string]string
		for _, folder := range comicFolders {
			if folder.IsDir() {
				comic := make(map[string]string)
				comic["name"] = folder.Name()

				files, err := ioutil.ReadDir("./comic/" + folder.Name())
				if err != nil {
					ctx.StatusCode(iris.StatusInternalServerError)
					ctx.WriteString("Error reading comic folder " + folder.Name())
					return
				}

				for _, file := range files {
					if !file.IsDir() && strings.ToLower(file.Name()) == "cover.jpg" {
						comic["cover"] = "/comic/" + folder.Name() + "/cover.jpg"
						break
					}
				}

				if _, ok := comic["cover"]; !ok {
					comic["cover"] = "/static/default_cover.jpg"
				}

				comics = append(comics, comic)
			}
		}

		sort.Slice(comics, func(i, j int) bool {
			return comics[i]["name"] < comics[j]["name"]
		})

		ctx.ViewData("comics", comics)
		ctx.View("index.html")
	})

	app.Get("/comic/{folder:string}", authMiddleware, func(ctx iris.Context) {
		files, err := ioutil.ReadDir("./comic/" + ctx.Params().GetString("folder"))
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.WriteString("Error reading comic folder " + ctx.Params().GetString("folder"))
			return
		}

		var images []string
		for _, file := range files {
			if !file.IsDir() && !strings.EqualFold(file.Name(), "cover.jpg") {
				images = append(images, "/comic/"+ctx.Params().GetString("folder")+"/"+file.Name())
			}
		}

		sort.Slice(images, func(i, j int) bool {
			return strings.Compare(images[i], images[j]) < 0
		})

		ctx.ViewData("folder", ctx.Params().GetString("folder"))
		ctx.ViewData("images", images)
		ctx.View("folder.html")
	})

	app.HandleDir("/static", "./static")

	app.Run(iris.Addr(":8080"))
}
