package game

import (
	socketio "github.com/googollee/go-socket.io"
)

type Player struct {
	Ready  bool
	Number uint8
	Socket *socketio.Conn
	Fleet  *PlayerFleet
}
