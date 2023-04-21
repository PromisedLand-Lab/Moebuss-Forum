package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

type Post struct {
	Title   string
	Content string
}

var posts []Post

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	// recovery from panics and log the requests to the terminal.
	app.Use(recover.New())
	app.Use(logger.New())

	// register a template engine
	app.RegisterView(iris.HTML("./views", ".html"))

	// serve static assets (e.g. css and js files)
	app.HandleDir("/static", "./static")

	// define a route to display the form
	app.Get("/", func(ctx iris.Context) {
		ctx.View("index.html")
	})

	// define a route to handle the form submission
	app.Post("/", func(ctx iris.Context) {
		// retrieve the form values
		title := ctx.FormValue("title")
		content := ctx.FormValue("content")

		// create a new post
		post := Post{title, content}

		// add the post to the list of posts
		posts = append(posts, post)

		// render the list of posts
		ctx.ViewData("Posts", posts)
		ctx.View("posts.html")
	})

	// define a route to display the list of posts
	app.Get("/posts", func(ctx iris.Context) {
		ctx.ViewData("Posts", posts)
		ctx.View("posts.html")
	})

	// start the server
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
