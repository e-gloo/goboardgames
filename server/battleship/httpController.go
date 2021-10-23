package battleship

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func JoinFriend(c echo.Context) error {
  code := c.Param("code")
  return c.String(http.StatusOK, code)
}