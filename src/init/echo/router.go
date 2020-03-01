package echo_init

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"main/src/init/i18n"
)

const (
	staticPath = "/static"
)
/*
 * init default logger and render page
 */
func InitEcho() *echo.Echo {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))
	// register static route
	e.Static(staticPath, "public")

	// register template renderer
	e.Renderer = newTemplateRenderer("views", ".html", i18n.I18)

	return e
}
