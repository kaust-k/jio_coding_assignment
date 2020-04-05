package db

import (
	"fmt"
	"jwt_server/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var dbClient *gorm.DB

func init() {
	cfg := config.GetConfig()
	url := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		cfg.DatabaseHost, cfg.DatabasePort, cfg.DatabaseUser, cfg.DatabaseName, cfg.DatabasePassword)
	var err error
	dbClient, err = gorm.Open("postgres", url)
	if err != nil {
		panic(err)
	}
}


// GetClient gets Postgres client
func GetClient() *gorm.DB {
	return dbClient
}
