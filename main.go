package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"message": "hello",
		})
	})
	r.GET("/hello", func(c *gee.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})
	r.GET("/hello/:name", func(c *gee.Context) {

		c.JSON(http.StatusOK, gee.H{
			"name": c.Param("name"),
		})
	})
	v1 := r.Group("/group")
	{
		v1.GET("/11", func(c *gee.Context) {
			c.JSON(http.StatusOK, gee.H{
				"message": "v1/11",
			})
		})
		v1.GET("/22", func(c *gee.Context) {
			c.JSON(http.StatusOK, gee.H{
				"message": "v1/22",
			})
		})
	}

	err := r.Run(":9999")
	if err != nil {
		panic(err)
	}
}
