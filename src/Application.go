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
	listenAddr := config.AppConfig.Conf.GetString("http.listen_addr", "0.0.0.0")
	listenPort := config.AppConfig.Conf.GetString("http.listen_port", "9000")
	driverName :=config.AppConfig.Conf.GetString("database.driver_name", "postgres")
	dataSourceName :=config.AppConfig.Conf.GetString("database.data_source_name",
		"postgres://postgres:123456@localhost:5432/template?sslmode=disable&client_encoding=UTF-8")
	maxIdleConns := int(config.AppConfig.Conf.GetInt32("database.max_idle_conns", 5))
	maxOpenConns := int(config.AppConfig.Conf.GetInt32("database.max_open_conns", 5))
	maxLifeTime := time.Duration( config.AppConfig.Conf.GetInt32("database.max_life_time", 60)) * time.Second

	cookieSecret := config.AppConfig.Conf.GetString("secure.cookie.secret", "secret")

	db := go_dal.Config{DriverName: driverName,
		DataSourceName: dataSourceName,
		MaxIdleConns:   maxIdleConns, MaxOpenConns: maxOpenConns, MaxLifeTime: maxLifeTime}
	db.Init()

	initHelper()
	initI18n()

	expireTime := time.Duration( config.AppConfig.Conf.GetInt32("secure.jwt.expire", 600)) * time.Second
	secret := config.AppConfig.Conf.GetString("secure.jwt.secret", "secret")
	config.InitJWTConf(expireTime, secret)

	engine := echo.InitEcho()
	controller.LoadRouterController(engine)
	api.LoadRouterApis(engine)
	engine.Use(session.Middleware(sessions.NewCookieStore([]byte(cookieSecret))))

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
