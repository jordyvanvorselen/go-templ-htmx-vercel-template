package users

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct{}

func (h Handler) Index(c echo.Context) error {
	return c.String(http.StatusOK, "Hello User Handler!")
}
