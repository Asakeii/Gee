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
		c.JSON(http.StatusOK, gee.H{
			"message": "hello world!",
		})
	})

	err := r.Run(":9999")
	if err != nil {
		panic(err)
	}
}
