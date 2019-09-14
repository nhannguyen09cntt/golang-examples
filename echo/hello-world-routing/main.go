package main

import (
	"hello-world-routing/routers"

	"github.com/labstack/echo"
)

func main() {
	// create a new echo instance
	e := echo.New()

	// routing
	routers.InitRoutes(e)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8080"))
}
