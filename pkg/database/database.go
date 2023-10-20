package database

import (
	"os"

	"github.com/harunalbayrak/golang-gin-boilerplate/config"
	"github.com/harunalbayrak/golang-gin-boilerplate/pkg/e"
	"github.com/jinzhu/gorm"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func GetDBConnectionString() string {
	config := config.GetConfig()
	user := config.GetString("databaseUser")
	host := config.GetString("databaseHost")
	port := config.GetString("databasePort")
	databaseName := config.GetString("databaseName")
	password := os.Getenv("DATABASE_PASSWORD")
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, databaseName)

	return connectionString
}

func CheckDatabaseConnection() error {
	db, err := gorm.Open("mysql", GetDBConnectionString())
	if err != nil {
		return e.Error(e.ERROR_CONNECTING_DATABASE)
	}
	defer db.Close()

	return err
}
