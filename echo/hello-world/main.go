package main

import (
	"hello-world/handlers"

	"github.com/labstack/echo"
)

func main() {
	// create a new echo instance
	e := echo.New()

	// routing
	e.GET("/hello", handlers.HelloPage())
	e.GET("/api/hello", handlers.HelloAPI())

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8080"))
}
