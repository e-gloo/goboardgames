package server

import (
	"github.com/labstack/echo/v4"
)

var (
	server *echo.Echo = nil
)

func GetServer() *echo.Echo {
	if server == nil {
		server = echo.New()
	}
	return server
}
