package middleware

import (
	"github.com/labstack/echo"
	es "main/src/init/databases/elasticsearch"
)
//TODO add request/response to map content base on bz
/*
 * init ES performance logger
 */
func ESPerLogger(next echo.HandlerFunc) echo.HandlerFunc {
	client := es.ClientSingleNode{
		Address:             "http://localhost:9200",
		MaxIdleConnsPerHost: 10,
	}
	if err := client.Init(); err != nil {
		panic(err)
	}
	requestLogger := es.InitLogger("go-template")
	return func(c echo.Context) error {
		//content := make(map[string] interface{})
		requestLogger.CreateLog("LamTNB") //add current user login
		//requestLogger.WriteLog("first send", map[string]interface{}{"name": "LamTNB"}, es.GetSingleNodeClient())
		return next(c)
	}
}
