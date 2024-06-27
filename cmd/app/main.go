package main

import (
	"Chat/configs"
	auth_route "Chat/modules/auth/route"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

var routePrefix = "/api/v1"

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("DOCKER_PORT")

	e := echo.New()

	//Set global group prefix
	global_group := e.Group(routePrefix)

	//Define routes
	auth_route.Register(global_group)

	//Init Database
	mongo_err := configs.InitMongo()

	if mongo_err != nil {
		log.Fatal(mongo_err)
	} else {
		fmt.Print("MongoDB connected")
	}

	//Start server
	e.Logger.Fatal(e.Start(":" + PORT))
	fmt.Print("Server started on port " + PORT)
}
