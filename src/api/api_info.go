package api

import (
	"github.com/LTNB/go-echo-template/src/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

func info(context echo.Context) error {

	return context.JSON(http.StatusOK, utils.ResponseJson{
		Status:  http.StatusOK,
		Message: "Ok",
		Data:    "This is data",
	} )
}
