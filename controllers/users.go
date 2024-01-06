package controllers

import (
	"github.com/jordyvanvorselen/go-templ-htmx-vercel-template/views/users"
	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

func (h UserHandler) Index(c echo.Context) error {
	return render(c, users.Index())
}
