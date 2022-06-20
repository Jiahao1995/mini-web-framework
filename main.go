package main

import (
	"dun"
	"net/http"
)

func main() {
	r := dun.New()
	r.GET("/", func(c *dun.Context) {
		c.HTML(http.StatusOK, "<h1>Hello dun</h1>")
	})

	r.GET("/hello", func(c *dun.Context) {
		// expect /hello?name=dunbarb
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *dun.Context) {
		// expect /hello/dunbarb
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *dun.Context) {
		c.JSON(http.StatusOK, dun.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}
