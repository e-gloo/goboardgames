package routes

import (
	// "github.com/e-gloo/goboardgames/server"
	"github.com/e-gloo/goboardgames/battleship"
)

// func AddAllHttpRoutes() {
// 	e := server.GetServer()
// }

func AddAllWsRoutes() {
	battleship.AddWSRoutes()
}
