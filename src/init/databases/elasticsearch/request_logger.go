package es

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type RequestLogger struct {
	index string
	s, d                 int64
	appName, sender, msg string
	info                 map[string]interface{}
}

func InitLogger(appName string) RequestLogger {
	instance := RequestLogger{appName: appName}
	return instance
}

func (logger *RequestLogger) CreateLog(sender string) {
	now := time.Now()
	logger.index = fmt.Sprintf("go-%v", now.Unix())
	logger.s = now.Unix()
	logger.sender = sender;
}

func (logger *RequestLogger) WriteLog(msg string, info map[string]interface{}, client ClientHandler) {
	now := time.Now().Unix()
	logger.d = logger.s - now
	logger.msg = msg
	logger.info = info
	body := make(map[string]interface{})
	body["s"] = logger.s
	body["app_name"] = logger.appName
	body["sender"] = logger.sender
	body["d"] = logger.d
	body["msg"] = logger.msg
	body["info"] = logger.info
	bodyJson, _ := json.Marshal(body)
	_, err := client.write(logger.index, strconv.FormatInt(now, 10), bodyJson)
	if err != nil {
		fmt.Println(err)
	}

}
