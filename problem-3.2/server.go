package main

import (
	"fmt"
	config "ormalta/problem-3.2/config"
	routes "ormalta/problem-3.2/routes"

	"github.com/labstack/echo/v4/middleware"

)

func main() {
	config.InitDb()
	fmt.Println("databases has been started!!")

	e := routes.New()

	e.Pre(middleware.RemoveTrailingSlash())

	e.Logger.Fatal(e.Start(":8000"))
}

