package server

import (
	"fmt"

	"github.com/harunalbayrak/golang-gin-boilerplate/config"
	"github.com/harunalbayrak/golang-gin-boilerplate/log"
)

var logger *log.Logger

func Init() {
	logger = log.NewLogger()
}

func Run() error {
	config := config.GetConfig()
	r := NewRouter()
	err := r.Run(fmt.Sprintf("%s:%s", config.GetString("listenAddress"), config.GetString("port")))

	return err
}
