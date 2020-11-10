package echo_init

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

//=========session utils======

func GetSession(c echo.Context) *sessions.Session {
	sess, _ := session.Get("session", c)
	return sess
}

func SetSessionValue(c echo.Context, key string, value interface{}) {
	sess := GetSession(c)
	if value == nil {
		delete(sess.Values, key)
	} else {
		sess.Values[key] = value
	}
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		panic(err)
	}
}

func GetLocale(c echo.Context) string {
	sess := GetSession(c)
	locale := sess.Values["locale"]
	if locale != nil {
		return locale.(string)
	}
	return ""
}

func RemoveSessionValue(c echo.Context, key string) {
	sess := GetSession(c)
	delete(sess.Values, key)
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		panic(err)
	}
}

func AddFlashMsg(c echo.Context, msg string) {
	sess := GetSession(c)
	sess.AddFlash(msg)
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		panic(err)
	}
}
