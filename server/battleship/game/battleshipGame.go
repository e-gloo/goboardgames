package game

import (
	"github.com/e-gloo/goboardgames/utils"
	socketio "github.com/googollee/go-socket.io"
)

type BattleshipGame struct {
	RoomCode string
	Players  map[uint8]*Player
	Turn     uint8
	Phase    BattleshipPhase
}

func NewBattleshipGame(roomCode string) *BattleshipGame {
	game := BattleshipGame{
		RoomCode: roomCode,
		Players:  make(map[uint8]*Player, 0),
		Turn:     uint8(utils.RandInt(0, 1)),
		Phase:    WAITING,
	}
	return &game
}

func (bg *BattleshipGame) GetPlayer(playerNb uint8) *Player {
	player, ok := bg.Players[playerNb]
	if !ok {
		return nil
	}
	return player
}

func (bg *BattleshipGame) GetOtherPlayers(playerNb uint8) []*Player {
	playerNbs := make([]*Player, len(bg.Players)-1)

	count := 0
	for k, v := range bg.Players {
		if k != playerNb {
			playerNbs[count] = v
			count += 1
		}
	}
	return playerNbs
}

func (bg *BattleshipGame) CheckAllPlayersReady() bool {
	for _, v := range bg.Players {
		if v.Ready == false {
			return false
		}
	}
	return true
}

func (bg *BattleshipGame) AddPlayer(socket *socketio.Conn) *Player {
	key := uint8(len(bg.Players) + 1)
	bg.Players[key] = &Player{
		Ready:  false,
		Number: key,
		Socket: socket,
		Fleet:  &PlayerFleet{Hits: make([]uint8, 0), Fleet: nil},
	}
	return bg.Players[key]
}

func (bg *BattleshipGame) RemovePlayer(playerNb uint8) {
	delete(bg.Players, playerNb)
	for _, player := range bg.Players {
		if player.Number > playerNb {
			player.Number -= 1
		}
	}
}

func (bg *BattleshipGame) NextTurn(playerNb uint8) {
	if playerNb == uint8(len(bg.Players)) {
		bg.Turn = 1
	} else {
		bg.Turn += 1
	}
}
