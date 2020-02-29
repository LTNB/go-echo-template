package config

import (
	hocon "github.com/go-akka/configuration"
	"log"
	"os"
	"path"
)

const (
	defaultConfigFile = "./configs/application.conf"
)

var (
	AppConfig *HoconConfig
)

// HoconConfig encapsulates application's configurations in HOCON format
type HoconConfig struct {
	File string        // config file
	Conf *hocon.Config // configurations
}

func LoadAppConfig(file string) *HoconConfig {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	defer os.Chdir(dir)

	config := HoconConfig{}
	log.Printf("Loading configurations from file [%s]", file)
	confDir, confFile := path.Split(file)
	os.Chdir(confDir)
	config.File = file
	config.Conf = hocon.LoadConfig(confFile)
	return &config
}

func InitAppConfig() *HoconConfig {
	configFile := os.Getenv("APP_CONFIG")
	if configFile == "" {
		log.Printf("No environment APP_CONFIG found, fallback to [%s]", defaultConfigFile)
		configFile = defaultConfigFile
	}
	return LoadAppConfig(configFile)
}
