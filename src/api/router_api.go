package api

import "github.com/labstack/echo/v4"
//handle api routing
const (
	deleteUserUrl = "api/user/:id"
)
func LoadRouterApis(e *echo.Echo){
	e.GET("api/info", info)
	e.GET("api/user/:id", getUserInfo)
	e.GET("api/user", getAll)
	e.POST("/api/user", create)
	e.DELETE(deleteUserUrl, deleteUser)
}

