package src

import (
	"fmt"
	go_dal "github.com/LTNB/go-dal"
	_ "github.com/LTNB/go-dal/postgres"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"main/src/api"
	"main/src/controller"
	config "main/src/init"
	es "main/src/init/databases/elasticsearch"
	"main/src/init/echo"
	"main/src/init/i18n"
	"main/src/models/user"
	"time"
)

func initI18n() {
	i18n.NewI18n("./configs/i18n")
}

/*
 * load helper support for data access layer
 */
func initHelper() {
	accountHelper := user.AccountHelper{}
	accountHelper.Init()
}

/*
 * init data source, default connect to postgresSQL
 */
func initDataSource() {
	driverName := config.AppConfig.Conf.GetString("database.driver_name", "postgres")
	dataSourceName := config.AppConfig.Conf.GetString("database.data_source_name",
		"postgres://postgres:123456@localhost:5432/template?sslmode=disable&client_encoding=UTF-8")
	maxIdleConns := int(config.AppConfig.Conf.GetInt32("database.max_idle_conns", 5))
	maxOpenConns := int(config.AppConfig.Conf.GetInt32("database.max_open_conns", 5))
	maxLifeTime := time.Duration(config.AppConfig.Conf.GetInt32("database.max_life_time", 60)) * time.Second

	db := go_dal.Config{DriverName: driverName,
		DataSourceName: dataSourceName,
		MaxIdleConns:   maxIdleConns, MaxOpenConns: maxOpenConns, MaxLifeTime: maxLifeTime}
	db.Init()
}

/*
 * init echo server
 */
func initEcho() {
	listenAddr := config.AppConfig.Conf.GetString("http.listen_addr", "0.0.0.0")
	listenPort := config.AppConfig.Conf.GetString("http.listen_port", "9000")
	cookieSecret := config.AppConfig.Conf.GetString("secure.cookie.secret", "secret")

	engine := echo_init.InitEcho()
	controller.LoadRouterController(engine)
	api.LoadRouterApis(engine)

	engine.Use(session.Middleware(sessions.NewCookieStore([]byte(cookieSecret))))

	engine.Logger.Fatal(engine.Start(listenAddr + ":" + listenPort))
	fmt.Println("Call Start success")

}

/*
 * init elastic search log
 */
func initESLogger() es.RequestLogger {
	esAddress := config.AppConfig.Conf.GetString("es.address", "http://localhost:9200")
	esMaxIdleConnsPerHost := int(config.AppConfig.Conf.GetInt32("es.max_idle_conns_per_host", 10))
	esAppName := config.AppConfig.Conf.GetString("es.app_name", "go-template")
	// sample using es logger
	client := es.ClientSingleNode{
		Address:             esAddress,
		MaxIdleConnsPerHost: esMaxIdleConnsPerHost,
	}
	client.Init()
	requestLogger := es.InitLogger(esAppName)

	//sample write log
	//requestLogger.CreateLog("LamTNB")
	requestLogger.WriteLog("first send", map[string]interface{}{"name": "LamTNB"}, es.GetSingleNodeClient())

	return requestLogger
}

/*
 * init jwt config
 */
func initJWTConfig() {
	expireTime := time.Duration(config.AppConfig.Conf.GetInt32("secure.jwt.expire", 600)) * time.Second
	secret := config.AppConfig.Conf.GetString("secure.jwt.secret", "secret")
	config.InitJWTConf(expireTime, secret)
}

func Start() {
	//load config
	config.AppConfig = config.InitAppConfig()

	initDataSource()
	initHelper()
	initI18n()
	initJWTConfig()
	//initESLogger()
	initEcho()

}
