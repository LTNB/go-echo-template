package middleware

import (
	es "github.com/LTNB/go-echo-template/src/init/databases/elasticsearch"
	"github.com/labstack/echo"
)
//TODO add request/response to map content base on bz
func MiddlewareESPerLogger(next echo.HandlerFunc) echo.HandlerFunc {
	client := es.ClientSingleNode{
		Address:             "http://localhost:9200",
		MaxIdleConnsPerHost: 10,
	}
	client.Init()
	requestLogger := es.InitLogger("go-template")
	return func(c echo.Context) error {
		//content := make(map[string] interface{})
		requestLogger.CreateLog("LamTNB") //add current user login
		next(c)
		//requestLogger.WriteLog("first send", map[string]interface{}{"name": "LamTNB"}, es.GetSingleNodeClient())
		return nil
	}
}
