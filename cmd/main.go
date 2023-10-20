package main

import (
	"github.com/harunalbayrak/golang-gin-boilerplate/config"
	"github.com/harunalbayrak/golang-gin-boilerplate/log"
	"github.com/harunalbayrak/golang-gin-boilerplate/models"
	"github.com/harunalbayrak/golang-gin-boilerplate/pkg/database"
	"github.com/harunalbayrak/golang-gin-boilerplate/server"
)

var logger *log.Logger

func init() {
	logger = log.NewLogger()
}

func main() {
	logger.Info("Starting golang gin boilerplate")

	err := config.Init("development")
	if err != nil {
		panic(err)
	}

	err = database.CheckDatabaseConnection()
	if err != nil {
		panic(err)
	}

	models.Setup()
	server.Init()
	err = server.Run()
	if err != nil {
		panic(err)
	}
}
