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
	s.Emit("gameOptions", game.GameConstants)
	return nil
}

func OnDisconnect(s socketio.Conn, reason string) {
	fmt.Println("closed ", s.ID(), reason)
}

func PlayWithFriend(s socketio.Conn) {
	randCode := "test"
	// randCode := generateCode()
	s.Join(randCode)
	s.Emit("newCode", randCode)
	g := game.NewBattleshipGame(randCode)
	TempMap[randCode] = g
	player := g.AddPlayer(&s)
	s.SetContext(&SocketContext{PlayerNb: player.Number, RoomCode: randCode})
	s.Emit("PlayerNb", player.Number)
}

func JoinRoom(s socketio.Conn, code string) string {
	srv := server.GetSocket()
	g, ok := TempMap[code]
	if ok {
		_, ok := s.Context().(*SocketContext)
		if !ok {
			player := g.AddPlayer(&s)
			s.SetContext(&SocketContext{PlayerNb: player.Number, RoomCode: code})
			s.Emit("PlayerNb", player.Number)
		}
		s.Join(code)
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

func RandomizeFleet(s socketio.Conn) string {
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
	player := g.GetPlayer(ctx.PlayerNb)
	player.Fleet.Fleet = fleet
	s.Emit("NewFleet", fleet)
	return "randomizeFleetOk"
}

func Ready(s socketio.Conn) string {
	ctx := s.Context().(*SocketContext)
	g, ok := TempMap[ctx.RoomCode]
	if !ok {
		return "ko"
	}

	srv := server.GetSocket()
	// otherPlayers := g.GetOtherPlayers(ctx.PlayerNb)
	player := g.GetPlayer(ctx.PlayerNb)
	player.Ready = true

	// check si tout le monde est prêt
	if g.CheckAllPlayersReady() {
		if g.Phase == game.PREPARATION {
			whosTurn := uint8(utils.RandInt(1, len(g.Players)+1))
			g.Phase = game.PLAYING
			g.Turn = whosTurn
			srv.BroadcastToRoom(Namespace, ctx.RoomCode, "gamePhaseUpdated", g.Phase)
			srv.BroadcastToRoom(Namespace, ctx.RoomCode, "PlayersTurn", whosTurn)
		} else if g.Phase == game.OVER {
			g.Phase = game.PREPARATION
			srv.BroadcastToRoom(Namespace, ctx.RoomCode, "gamePhaseUpdated", g.Phase)
		}
	}

	srv.BroadcastToRoom(Namespace, ctx.RoomCode, "PlayerReady", ctx.PlayerNb)

	return "ok"
}

func Attack(s socketio.Conn, cell uint8, playerNb uint8) string {
	srv := server.GetSocket()
	ctx := s.Context().(*SocketContext)
	g, ok := TempMap[ctx.RoomCode]
	if !ok || ctx.PlayerNb != g.Turn {
		return "attackError"
	}

	player := g.GetPlayer(playerNb)
	if player == nil || player.Fleet.IsAlreadyHit(cell) {
		return "attackError"
	}

	ship := player.Fleet.Attack(cell)
	var result AttackResultEnum = MISS
	if ship == nil {
		result = MISS
	} else if ship.RemainingAlive > 0 {
		result = HIT
	} else {
		result = SUNK
	}

	srv.BroadcastToRoom(Namespace, ctx.RoomCode, "AttackResult", &AttackResult{Position: cell, Result: result, PlayerNb: playerNb})
	if result == SUNK {
		srv.BroadcastToRoom(Namespace, ctx.RoomCode, "BoatSunk", ship.Cells, playerNb)
	}

	if result == SUNK {
		winner := g.GetWinner()
		if winner != nil {
			g.Phase = game.OVER
			srv.BroadcastToRoom(Namespace, ctx.RoomCode, "gamePhaseUpdated", g.Phase, winner.Number)
			g.ResetReady()
		}
	} else if result != HIT {
		g.NextTurn(ctx.PlayerNb)
		srv.BroadcastToRoom(Namespace, ctx.RoomCode, "PlayersTurn", g.Turn)
	}

	return "attackOk"
}
