package users

import (
	"github.com/jordyvanvorselen/go-templ-htmx-vercel-template/controllers"
	"github.com/jordyvanvorselen/go-templ-htmx-vercel-template/views/users"
	"github.com/labstack/echo/v4"
)

type Handler struct{}

func (h Handler) Index(c echo.Context) error {
	return controllers.Render(c, users.Index())
}
