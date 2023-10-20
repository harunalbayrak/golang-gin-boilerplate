package config

import (
	"path/filepath"

	"github.com/harunalbayrak/golang-gin-boilerplate/log"
	"github.com/harunalbayrak/golang-gin-boilerplate/pkg/e"
	"github.com/spf13/viper"
)

var config *viper.Viper
var logger *log.Logger

func Init(env string) error {
	var err error

	logger = log.NewLogger()
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(env)
	config.AddConfigPath("../config/")
	config.AddConfigPath("config/")
	err = config.ReadInConfig()
	if err != nil {
		e.Error(e.ERROR_PARSING_CONFIG_FILE)
	}

	return err
}

func relativePath(basedir string, path *string) {
	p := *path
	if len(p) > 0 && p[0] != '/' {
		*path = filepath.Join(basedir, p)
	}
}

func GetConfig() *viper.Viper {
	return config
}
