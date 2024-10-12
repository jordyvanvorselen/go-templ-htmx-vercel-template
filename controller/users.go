package controller

import (
	"github.com/jordyvanvorselen/go-templ-htmx-vercel-template/view/user"
	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

func (h UserHandler) Index(c echo.Context) error {
	return render(c, user.Index())
}
