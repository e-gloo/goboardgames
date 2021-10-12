package main

import (
	"github.com/e-gloo/goboardgames/battleship"
	"github.com/e-gloo/goboardgames/routes"
	"github.com/e-gloo/goboardgames/server"
)

func main() {
	e := server.GetServer()
	routes.AddRoutes()
	routes.SetupCors()
	battleship.Setup()
	e.Logger.Fatal(e.Start(":1323"))
}
