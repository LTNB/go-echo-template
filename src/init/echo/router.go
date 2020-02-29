package echo

import (
	"github.com/LTNB/go-echo-template/src/init/i18n"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	staticPath = "/static"
)

func InitEcho() *echo.Echo {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))
	// register static route
	e.Static(staticPath, "public")

	// register template renderer
	e.Renderer = newTemplateRenderer("src/templates", ".html", i18n.I18)

	return e
}
