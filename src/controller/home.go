package controller

import (
	"github.com/labstack/echo/v4"
	config "main/src/init"
	echoConf "main/src/init/echo"
	"main/src/init/i18n"
	"main/src/models/user"
	"main/src/utils"
	"net/http"
)

/*
 * GET: ${/} or ${/cp/user}
 * render: ${home.html}
 */
func home(c echo.Context) error {
	accountHelper := user.GetAccountHelper()
	k, _ := accountHelper.GetAllAsMap()
	return c.Render(http.StatusOK, "layout:home", map[string]interface{}{
		"active":  "home",
		"bo_list": k,
	})
}

/*
 * GET: ${/cp/user/create}
 * render: ${create_update_form.html}
 */
func createUser(c echo.Context) error {
	return c.Render(http.StatusOK, "layout:create_update_form", map[string]interface{}{
		"active": "home",
	})
}

/*
 * POST: ${/cp/user/create}
 * data: ${Account}
 * redirect ${/cp/user} if success
 */
func createUserSubmit(c echo.Context) error {
	var errMsg string
	locale := echoConf.GetLocale(c)
	bo := user.Account{}

	accountHelper := user.GetAccountHelper()

	node, _ := utils.NewNode(1)
	if err := c.Bind(&bo); err != nil {
		errMsg = i18n.I18.FlashMsg(locale, "error_form_400")
	}
	if ok, err := accountHelper.EmailIsExisted(bo.Email); ok || err != nil {
		errMsg = i18n.I18.FlashMsg(locale, "error_form_email_existed")
		goto end
	}
	bo.ID = node.Generate().String()
	bo.Password, _ = config.HashString(bo.Password)
	if _, err := accountHelper.Create(bo); err != nil {
		errMsg = i18n.I18.FlashMsg(locale, "error_form_400")
		goto end
	}
	echoConf.AddFlashMsg(c, i18n.I18.FlashMsg(locale,"success_form_create", bo.Email))
	return c.Redirect(http.StatusFound, homeUrl)
end:
	return c.Render(http.StatusOK, "layout:create_update_form", map[string]interface{}{
		"active":   "home",
		"data":     utils.StructToMapAsTag(bo, "form"),
		"error":    errMsg,
		"editMode": false,
	})
}

/*
 * GET: ${/cp/user/edit/:id}
 */
func editUser(c echo.Context) error {
	id := c.Param("id")

	var errMsg string
	locale := echoConf.GetLocale(c)

	accountHelper := user.GetAccountHelper()
	bo := user.Account{}
	var data map[string]interface{}
	if id == "" {
		errMsg = i18n.I18.FlashMsg(locale,"error_user_not_found")
		goto end
	}
	bo.ID = id
	data, _ = accountHelper.GetOneAsMap(&bo)
	if bo.ID == "" {
		errMsg = i18n.I18.FlashMsg(locale,"error_user_not_found")
		goto end
	}
	delete(data, "password")
	return c.Render(http.StatusOK, "layout:create_update_form", map[string]interface{}{
		"active":   "home",
		"data":     data,
		"error":    errMsg,
		"editMode": true,
	})
end:
	return c.Render(http.StatusOK, "layout:create_update_form", map[string]interface{}{
		"active": "home",
		"error":  errMsg,
	})
}

/*
 * GET: ${/cp/user/edit/:id}
 * data: ${Account}
 * redirect ${/cp/user} if success
 */
func editUserSubmit(c echo.Context) error {
	bo := user.Account{}
	var err error
	var affected int64

	accountHelper := user.GetAccountHelper()
	userDb := user.Account{}
	var errMsg, warnMsg string
	locale := echoConf.GetLocale(c)

	if err := c.Bind(&bo); err != nil {
		errMsg = i18n.I18.FlashMsg(locale,"error_form_400")
		goto end
	}
	userDb.ID = bo.ID
	err = accountHelper.GetOne(&userDb)
	if err != nil {
		errMsg = i18n.I18.FlashMsg(locale,"error_form_400")
		goto end
	}

	if userDb.Email == "" {
		errMsg = i18n.I18.FlashMsg(locale,"error_user_not_found")
		goto end
	}

	if bo.Email != userDb.Email {
		if ac, err := accountHelper.FindByEmail(bo.Email); ac != nil || err != nil {
			errMsg = i18n.I18.FlashMsg(locale,"error_form_email_existed")
			goto end
		}
	}

	//reuse password if not change
	if bo.Password == "" {
		bo.Password = userDb.Password
	} else {
		bo.Password, _ = config.HashString(bo.Password)
	}

	affected, err = accountHelper.Update(bo)
	if err != nil {
		errMsg = i18n.I18.FlashMsg(locale,"error_user_update")
		goto end
	}

	if affected == 0 {
		//update error
		warnMsg = i18n.I18.FlashMsg(locale,"warn_user_not_change")
		goto end
	}
	echoConf.AddFlashMsg(c, i18n.I18.FlashMsg(locale,"success_form_update", bo.Email))
	return c.Redirect(http.StatusFound, homeUrl)

end:
	return c.Render(http.StatusOK, "layout:create_update_form", map[string]interface{}{
		"active":   "home",
		"error":    errMsg,
		"warn":     warnMsg,
		"editMode": false,
		"data":     utils.StructToMapAsTag(bo, "form"),
	})
}
