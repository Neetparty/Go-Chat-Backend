package main

import (
	auth_route "Chat/modules/auth/route"
	"fmt"

	"github.com/labstack/echo/v4"
)

var routePrefix = "/api/v1"

func main() {

	e := echo.New()

	//Set global group prefix
	global_group := e.Group(routePrefix)

	//Define routes
	auth_route.Register(global_group)

	//Start server
	e.Logger.Fatal(e.Start(":1010"))
	fmt.Print("Server started on port 1010")
}
