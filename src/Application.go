package src

import (
	"fmt"
	"github.com/LTNB/go-dal"
	_ "github.com/LTNB/go-dal/postgres"
	"github.com/LTNB/go-echo-template/src/api"
	"github.com/LTNB/go-echo-template/src/controller"
	config "github.com/LTNB/go-echo-template/src/init"
	"github.com/LTNB/go-echo-template/src/init/echo"
	"github.com/LTNB/go-echo-template/src/init/i18n"
	"github.com/LTNB/go-echo-template/src/models/user"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"time"
)

func initI18n() {
	i18n.NewI18n("./configs/i18n")
}

func initHelper() {
	accountHelper := user.AccountHelper{}
	accountHelper.Init()
}
func Start() {
	config.AppConfig = config.InitAppConfig()
	db := go_dal.Config{DriverName: "postgres",
		DataSourceName: "postgres://postgres:123456@localhost:5432/template?sslmode=disable&client_encoding=UTF-8",
		MaxIdleConns:   5, MaxOpenConns: 5, MaxLifeTime: 1 * time.Minute}
	db.Init()
	initHelper()
	listenAddr := config.AppConfig.Conf.GetString("http.listen_addr", "0.0.0.0")
	listenPort := config.AppConfig.Conf.GetString("http.listen_port", "9000")
	initI18n()

	engine := echo.InitEcho()
	controller.LoadRouterController(engine)
	api.LoadRouterApis(engine)
	engine.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	// sample using es logger
	//client := es.ClientSingleNode{
	//	Address:             "http://localhost:9200",
	//	MaxIdleConnsPerHost: 10,
	//}
	//client.Init()
	//requestLogger := es.InitLogger("go-template")
	//requestLogger.CreateLog("LamTNB")
	//requestLogger.WriteLog("first send", map[string]interface{}{"name": "LamTNB"}, es.GetSingleNodeClient())

	// sample using statsd logger
	//engine.Use(profiler.ProfilerWithConfig(profiler.ProfilerConfig{Address: "localhost:8125", Service: "app-name"}))

	engine.Logger.Fatal(engine.Start(listenAddr + ":" + listenPort))
	fmt.Println("Call Start success")

}
