package main

import (
	"dun"
	"net/http"
)

func main() {
	engine := dun.New()
	engine.GET("/", func(ctx *dun.Context) {
		ctx.HTML(http.StatusOK, "<h1>Hello Dun</h1>")
	})
	engine.GET("/hello", func(ctx *dun.Context) {
		ctx.String(http.StatusOK, "hello %s, you are at %s\n", ctx.Query("name"), ctx.Path)
	})

	engine.POST("/login", func(ctx *dun.Context) {
		ctx.JSON(http.StatusOK, dun.H{
			"username": ctx.PostForm("username"),
			"password": ctx.PostForm("password"),
		})
	})

	engine.Run(":9999")
}
