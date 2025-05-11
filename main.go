package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprintf(writer, "URL.PATH = %q\n", request.URL.Path)
	})

	r.GET("/hello", func(writer http.ResponseWriter, request *http.Request) {
		for k, v := range request.Header {
			_, _ = fmt.Fprintf(writer, "Header[%q] = %q\n", k, v)
		}
	})

	err := r.Run(":9999")
	if err != nil {
		panic(err)
	}
}
