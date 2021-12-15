package main

/*
$ curl "http://localhost:9999"
Hello JeffZuo
$ curl "http://localhost:9999/panic"
{"message":"Internal Server Error"}
$ curl "http://localhost:9999"
Hello JeffZuo
*/

import (
	"jin"
	"net/http"
)

func main() {
	r := jin.Default()
	r.GET("/", func(c *jin.Context) {
		c.String(http.StatusOK, "Hello JeffZuo\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *jin.Context) {
		names := []string{"jeffZuo"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
