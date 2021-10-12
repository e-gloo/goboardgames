package battleship

import (
	"fmt"
	"github.com/e-gloo/goboardgames/server"
	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
	"github.com/labstack/echo/v4"
	"net/http"
)

var allowOriginFunc = func(r *http.Request) bool {
	return true
}

func Setup() {
	opts := &engineio.Options{
		Transports: []transport.Transport{
			&polling.Transport{
				CheckOrigin: allowOriginFunc,
			},
			&websocket.Transport{
				CheckOrigin: allowOriginFunc,
			},
		},
	}
	srv := socketio.NewServer(opts)

	srv.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})

	srv.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		fmt.Println("notice:", msg)
		// s.Emit("reply", "a marqué: "+msg)
		srv.BroadcastToNamespace("/", "reply", "a marqué: "+msg)
	})

	srv.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		return "recv " + msg
	})

	srv.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})

	srv.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	srv.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})

	srv.OnEvent("/", "token", func(s socketio.Conn, msg string) string {
		fmt.Println("tokn:", msg)
		return msg //Sending ack with data in msg back to client, using "return statement"
	})

	go srv.Serve()
	// defer srv.Close()

	e := server.GetServer()
	e.HideBanner = true

	// e.Static("/", "../asset")
	e.Any("/socket.io/", func(context echo.Context) error {
		srv.ServeHTTP(context.Response(), context.Request())
		return nil
	})
}
