package api

import (
	"github.com/jordyvanvorselen/go-templ-htmx-vercel-template/api/controllers/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/api/users", users.Handler{}.Index)

	e.ServeHTTP(w, r)
}
