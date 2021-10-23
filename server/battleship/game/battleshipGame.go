package game

import (
	"github.com/e-gloo/goboardgames/utils"
	socketio "github.com/googollee/go-socket.io"
)

type BattleshipGame struct {
	RoomCode      string
	Player1       *PlayerFleet
	Player2       *PlayerFleet
	Player1Socket *socketio.Conn
	Player2Socket *socketio.Conn
	Turn          uint8
	Phase         BattleshipPhase
}

func NewBattleshipGame(roomCode string) *BattleshipGame {
	game := BattleshipGame{
		RoomCode:      roomCode,
		Player1Socket: nil,
		Player2Socket: nil,
		Player1:       &PlayerFleet{Fleet: nil, Hits: nil},
		Player2:       &PlayerFleet{Fleet: nil, Hits: nil},
		Turn:          uint8(utils.RandInt(0, 1)),
		Phase:         WAITING,
	}
	return &game
}
