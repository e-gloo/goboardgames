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

func (p *Player) AliveShips() bool {
	if p.Fleet == nil {
		return false
	}
	for _, fleet := range p.Fleet.Fleet {
		if fleet.RemainingAlive > 0 {
			return true
		}
	}
	return false
}

func (p *Player) ResetFleet() {
	p.Fleet.Hits = p.Fleet.Hits[0:0]
	p.Fleet.Fleet = NewFleet()
}
