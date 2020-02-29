package controller

import (
	"github.com/LTNB/go-echo-template/src/init/middleware"
	"github.com/labstack/echo/v4"
)

const (
	homeUrl           = "/cp/user"
	submitCreateUrl   = "/cp/user/create"
	createUserUrl     = "/cp/user/create"
	editUserUrl       = "/cp/user/edit/:id"
	submitEditUserUrl = "/cp/user/edit/:id"

	loginUrl        = "/login"
	logoutUrl		= "/logout"
)
var auth *middleware.RequiredAuthConf

func LoadRouterController(e *echo.Echo) {
	auth = middleware.GetDefaultRequiredAuthConfig()

	e.GET(loginUrl, login)
	e.POST(loginUrl, loginSubmit)
	e.GET(logoutUrl, logout, auth.MiddlewareRequiredAuth)

	e.GET("/", home, auth.MiddlewareRequiredAuth)
	e.GET(homeUrl, home, auth.MiddlewareRequiredAuth)
	e.GET(createUserUrl, createUser, auth.MiddlewareRequiredAuth)
	e.POST(submitCreateUrl, createUserSubmit, auth.MiddlewareRequiredAuth)
	e.GET(editUserUrl, editUser, auth.MiddlewareRequiredAuth)
	e.POST(submitEditUserUrl, editUserSubmit, auth.MiddlewareRequiredAuth)




}
