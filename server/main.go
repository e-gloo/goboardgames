package main

import (
	"github.com/e-gloo/goboardgames/battleship"
	"github.com/e-gloo/goboardgames/routes"
	"github.com/e-gloo/goboardgames/server"
	"github.com/e-gloo/goboardgames/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	utils.InitRandom()
	e := server.GetServer()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://172.23.164.159:8080", "http://172.23.100.254:8080"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
	}))
	routes.AddAllWsRoutes()

	// ctx := context.Background()

	// rdb := redis.NewClient(&redis.Options{
	// 	Addr:	  "localhost:6379",
	// 	Password: "", // no password set
	// 	DB:		  1,  // use default DB
	// })

	// err := rdb.Set(ctx, "key", "value", 0).Err()
	// if err != nil {
	// 	panic(err)
	// }
	// val, err := rdb.Get(ctx, "key").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("key", val)

	e.HideBanner = true

	srv := server.GetSocket()
	srv.OnConnect("/", battleship.OnConnect)
	// e.Static("/", "../asset")
	go srv.Serve()
	e.Any("/socket.io/", func(context echo.Context) error {
		srv.ServeHTTP(context.Response(), context.Request())
		return nil
	})

	e.Logger.Fatal(e.Start(":1323"))
}

// "Handshake"
// step 1 : user 1 click play w/ friend
// step 2 : room code creation
// step 3 : user 1 in waiting for user 2 connection
// step 4 : user 2 connect to room
// step 5 : user 2 broadcast to room : ready
// step 6 : user 1 and user 2 start game
