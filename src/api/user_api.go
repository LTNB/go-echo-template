package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	config "main/src/init"
	echo_conf "main/src/init/echo"
	"main/src/init/i18n"
	"main/src/models/user"
	"main/src/utils"
)

func getUserInfo(c echo.Context) error {
	id := c.Param("id")
	account := user.Account{}
	account.ID = id
	accountHelper := user.GetAccountHelper()
	accountHelper.GetOne(&account)
	return c.JSON(http.StatusOK, utils.ResponseJson{
		Status:  http.StatusOK,
		Message: "Ok",
		Data:    account,
	})
}

func getAll(c echo.Context) error {
	accountHelper := user.GetAccountHelper()
	data, _ := accountHelper.GetAll()
	return c.JSON(http.StatusOK, utils.ResponseJson{
		Status:  http.StatusOK,
		Message: "Ok",
		Data:    data,
	})
}

func create(c echo.Context) error {
	accountHelper := user.GetAccountHelper()
	bo := user.Account{}
	c.Bind(&bo)
	node, _ := utils.NewNode(1)
	bo.ID = node.Generate().String()
	bo.Password, _ = config.HashString(bo.Password)
	accountHelper.Create(bo)
	return c.JSON(http.StatusOK, utils.ResponseJson{
		Status:  http.StatusOK,
		Message: "Ok",
		Data:    utils.StructToMapAsTag(bo, "json"),
	})
}

func deleteUser(c echo.Context) error {
	id := c.Param("id")
	var errMsg string
	var affected int64
	var err error
	locale := echo_conf.GetLocale(c)

	accountHelper := user.GetAccountHelper()
	if id =="1" {
		errMsg = i18n.I18.FlashMsg(locale, "error_delete_demo_account")
		goto end
	}
	if id == "" {
		errMsg = i18n.I18.FlashMsg(locale, "error_user_not_found")
		goto end
	}

	affected, err = accountHelper.Delete(user.Account{ID: id,})

	if affected == 0 || err != nil {
		goto end
	}
	return c.JSON(http.StatusOK, utils.ResponseJson{
		Status:  http.StatusOK,
		Message: i18n.I18.FlashMsg(locale, "success_api_delete", ),
		Data:    true,
	})
end:
	return c.JSON(http.StatusOK,  utils.ResponseJson{
		Status:  http.StatusBadRequest,
		Message: errMsg,
		Data:    false,
	})
}
