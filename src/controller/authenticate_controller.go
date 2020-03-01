package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"main/src/init"
	echo_conf "main/src/init/echo"
	"main/src/init/i18n"
	"main/src/models/user"
	"main/src/utils"
)

/*
 * login bo wih email + password
 * set default locate when login
 */
type LoginBo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Locale   string `json:"locale"`
}

/*
 * GET: ${/login}
 * render login page: GET: ${login.html}
 * redirect to ${/} if ${CheckSession} is true
 */
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

/*
 * POST: /login
 * data: ${LoginBo}
 * generate ${token} + ${locale} and add to ${session}
 * if ${locale} os nil => using default ${locale}
 */
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
		//gen token as JWT
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

/*
 * GET: /logout
 * remove ${token}
 * redirect to ${/login}
 */

func logout(c echo.Context) error {
	echo_conf.RemoveSessionValue(c, "token")
	return c.Redirect(http.StatusFound, "login")
}
