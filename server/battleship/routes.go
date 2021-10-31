package battleship

import (
	"github.com/e-gloo/goboardgames/server"
)

var Namespace = "/battleship"

func AddHttpRoutes() {
	e := server.GetServer()

	g := e.Group(Namespace)
	g.GET("playWithFriend/:code", JoinFriend)
}

func AddWSRoutes() {
	srv := server.GetSocket()

	srv.OnConnect(Namespace, OnConnectB)
	srv.OnEvent(Namespace, "playWithFriend", PlayWithFriend)
	srv.OnEvent(Namespace, "joinRoom", JoinRoom)
	srv.OnEvent(Namespace, "randomizeFleet", RandomizeFleet)
	srv.OnEvent(Namespace, "ready", Ready)
	srv.OnEvent(Namespace, "bye", Bye)
	srv.OnError(Namespace, OnError)
	srv.OnDisconnect(Namespace, OnDisconnect)
}
