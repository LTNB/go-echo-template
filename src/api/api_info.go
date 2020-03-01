package api

import (
	"github.com/labstack/echo/v4"
	"main/src/utils"
	"net/http"
)

func info(context echo.Context) error {

	return context.JSON(http.StatusOK, utils.ResponseJson{
		Status:  http.StatusOK,
		Message: "Ok",
		Data:    "This is data",
	} )
}
