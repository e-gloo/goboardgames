package battleship

import (
	"fmt"

	"github.com/e-gloo/goboardgames/battleship/game"
	"github.com/e-gloo/goboardgames/server"
	"github.com/e-gloo/goboardgames/utils"
	socketio "github.com/googollee/go-socket.io"
)

var TempMap = make(map[string]*game.BattleshipGame)

type SocketContext struct {
	PlayerNb uint8
	RoomCode string
}

func generateCode() string {
	return utils.RandSeq(10)
}

func OnConnect(s socketio.Conn) error {
	// s.SetContext(SocketContext{})
	fmt.Println("connected /:", s.ID())
	return nil
}

func OnConnectB(s socketio.Conn) error {
	fmt.Println("connected /bat:", s.ID())
	return nil
}

func OnDisconnect(s socketio.Conn, reason string) {
	fmt.Println("closed ", s.ID(), reason)
}

func Bye(s socketio.Conn) string {
	last := s.Context().(string)
	s.Emit("bye", last)
	s.Close()
	return last
}

func PlayWithFriend(s socketio.Conn) {
	randCode := generateCode()
	s.Join(randCode)
	s.SetContext(&SocketContext{PlayerNb: 1, RoomCode: randCode})
	s.Emit("newCode", randCode)
	fmt.Printf("Setting player N°1 for room %s\n", randCode)
	TempMap[randCode] = game.NewBattleshipGame(randCode)
	TempMap[randCode].Player1Socket = &s
}

func JoinRoom(s socketio.Conn, code string) string {
	srv := server.GetSocket()
	_, ok := s.Context().(*SocketContext)
	if !ok {
		fmt.Printf("Setting player N°2 for room %s\n", code)
		s.SetContext(&SocketContext{PlayerNb: 2, RoomCode: code})
	}
	if g, ok := TempMap[code]; ok {
		s.Join(code)
		g.Player2Socket = &s
		g.Phase = game.PREPARATION
		srv.BroadcastToRoom(Namespace, code, "gamePhaseUpdated", g.Phase)
	} else {
		return "joinRoomError"
	}
	return "joinRoomOk"
}

func OnError(s socketio.Conn, e error) {
	fmt.Println("meet error:", e)
}

func RandomizeFleet(s socketio.Conn, e error) string {
	ctx := s.Context().(*SocketContext)

	g, ok := TempMap[ctx.RoomCode]
	if !ok {
		return "randomizeFleetError"
	}

	if g.Phase != game.PREPARATION {
		return "randomizeFleetError"
	}

	fleet := game.NewFleet()
	fmt.Println(fleet)
	if ctx.PlayerNb == 1 {
		g.Player1.Fleet = fleet
	} else {
		g.Player2.Fleet = fleet
	}
	s.Emit("NewFleet", fleet)
	return "randomizeFleetOk"
}
