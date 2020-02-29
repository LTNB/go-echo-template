package middleware

import (
	config "github.com/LTNB/go-echo-template/src/init"
	echo2 "github.com/LTNB/go-echo-template/src/init/echo"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/url"
	"time"
)

var lastSessionCheck = time.Now()

type ILoginHandler interface {
	CheckSession(token string) bool
}

type RequiredAuthConf struct {
	LoginUri string
	TokenKey string
	ILoginHandler
}

func NewRequiredAuth(loginUri string, tokenKey string, loginHandler ILoginHandler) RequiredAuthConf {
	instance := RequiredAuthConf{
		LoginUri:      loginUri,
		TokenKey:      tokenKey,
		ILoginHandler: loginHandler,
	}
	return instance
}

func GetDefaultRequiredAuthConfig() *RequiredAuthConf {
	instance := RequiredAuthConf{
		LoginUri:      "/login",
		TokenKey:      "token",
		ILoginHandler: DefaultLoginHandler{},
	}
	return &instance
}

func (requiredAuth RequiredAuthConf) MiddlewareRequiredAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess := echo2.GetSession(c)
		returnUrl := url.QueryEscape(c.Request().RequestURI)
		token := sess.Values[requiredAuth.TokenKey]
		if token == nil || token == "" {
			return c.Redirect(http.StatusFound, requiredAuth.LoginUri+"?returnUrl="+returnUrl)
		}
		if requiredAuth.checkSession(token.(string)) {
			return next(c)
		}
		return c.Redirect(http.StatusFound, requiredAuth.LoginUri+"?returnUrl="+returnUrl)
	}
}

func (requiredAuth RequiredAuthConf) checkSession(token string) bool {
	return requiredAuth.ILoginHandler.CheckSession(token)
}

type DefaultLoginHandler struct {
}

func (defaultLoginHandler DefaultLoginHandler) CheckSession(token string) bool {
	data := config.ParseToken(token)
	if data != nil {
		return true
	}
	return false
}