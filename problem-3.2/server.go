package main

import (
	config "ormalta/problem-3.2/config"
	routes "ormalta/problem-3.2/routes"
)

func main() {
	config.InitDb()

	e := routes.New()

	e.Logger.Fatal(e.Start(":8000"))
}

