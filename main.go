package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.Default()
	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "Gee test\n")
	})
	r.GET("/index", func(c *gee.Context) {
		c.String(http.StatusOK, "<h1>Index Page</h1>", "message")
	})
	r.GET("/panic", func(c *gee.Context) {
		names := []string{"panic test"}
		c.String(http.StatusOK, names[100])
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/hello", func(c *gee.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}

	r.Run(":9999")
}
