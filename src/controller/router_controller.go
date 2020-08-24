package controller

import (
	"github.com/labstack/echo/v4"
	echo_md "github.com/labstack/echo/v4/middleware"
	"main/src/init/middleware"
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

/*
 * web application's routers
 */

func LoadRouterController(e *echo.Echo) {
	// add default authentication
	auth = middleware.GetDefaultRequiredAuthConfig()
	e.Use(echo_md.Gzip())
	e.GET(loginUrl, login)
	e.POST(loginUrl, loginSubmit)
	e.GET(logoutUrl, logout, auth.MiddlewareRequiredAuth)
	e.GET("/", home)
	e.GET(homeUrl, home, auth.MiddlewareRequiredAuth)
	e.GET(createUserUrl, createUser, auth.MiddlewareRequiredAuth)
	e.POST(submitCreateUrl, createUserSubmit, auth.MiddlewareRequiredAuth)
	e.GET(editUserUrl, editUser, auth.MiddlewareRequiredAuth)
	e.POST(submitEditUserUrl, editUserSubmit, auth.MiddlewareRequiredAuth)




}
