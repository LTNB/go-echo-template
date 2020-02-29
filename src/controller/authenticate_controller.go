package controller

import (
	"fmt"
	"github.com/LTNB/go-echo-template/src/init"
	echo_conf "github.com/LTNB/go-echo-template/src/init/echo"
	"github.com/LTNB/go-echo-template/src/init/i18n"
	"github.com/LTNB/go-echo-template/src/models/user"
	"github.com/LTNB/go-echo-template/src/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

type LoginBo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Locale   string `json:"locale"`
}

func login(c echo.Context) error {
	sess := echo_conf.GetSession(c)
	if sess != nil {
		token := sess.Values[auth.TokenKey]
		if token != nil {
			if auth.CheckSession(token.(string)) {
				return c.Redirect(http.StatusFound, "/")
			}
		}
	}
	return c.Render(http.StatusOK, "login", nil)
}

func loginSubmit(c echo.Context) error {
	var errMsg string
	login := LoginBo{}
	bo := &user.Account{}
	var err error
	accountHelper := user.GetAccountHelper()

	if err := c.Bind(&login); err != nil {
		errMsg = i18n.I18.Text("error_form_400")
		goto end
	}

	bo, err = accountHelper.FindByEmail(login.Email)

	if err == nil && bo != nil && config.IsValid(bo.Password, login.Password) {
		data := make(map[string]interface{}, 1)
		data["email"] = bo.Email

		token, err := config.GenerateToken(data)
		if err != nil {
			errMsg = i18n.I18.Text("error_form_400")
			goto end
		}

		echo_conf.SetSessionValue(c, "token", token)
		echo_conf.SetSessionValue(c, "locale", login.Locale)

		c.Set("flashMsg", fmt.Sprintf(i18n.I18.FlashMsg(login.Locale, "login_success"), bo.Email))

		return c.Redirect(http.StatusFound, homeUrl)
	} else {
		errMsg = i18n.I18.Text("error_login_failed")
		goto end
	}

end:
	return c.Render(http.StatusOK, "login", map[string]interface{}{
		"data":     utils.StructToMapAsTag(login, "json"),
		"error":    errMsg,
		"editMode": false,
	})
}

func logout(c echo.Context) error {
	echo_conf.RemoveSessionValue(c, "token")
	return c.Redirect(http.StatusFound, "login")
}
