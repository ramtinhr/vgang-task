package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/ramtinhr/vgang-task/api"
	"github.com/ramtinhr/vgang-task/models"
	"github.com/ramtinhr/vgang-task/provider"
	"github.com/ramtinhr/vgang-task/service"
	"github.com/ramtinhr/vgang-task/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading env file")
	}

	if utils.PgsqlDB == nil {
		db, err := gorm.Open(postgres.New(postgres.Config{
			DSN:                  os.Getenv("DB_CONN"),
			PreferSimpleProtocol: true,
		}), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})

		if err != nil {
			logrus.Fatal(err)
		}

		utils.PgsqlDB = db

		db.AutoMigrate(models.Indexer{}, models.Product{})
	}
}

func main() {
	config := &service.Config{
		ServicePort: os.Getenv("PORT"),
		ServiceName: os.Getenv("NAME"),
		Version:     os.Getenv("VERSION"),
		Env:         os.Getenv("ENV"),
	}

	vgang := &provider.VgangUser{
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
	}

	indexer, err := vgang.Login()
	if err != nil {
		logrus.Errorf("Login Error %s", err)
		return
	}

	vgang.FetchProducts(indexer.AccessToken)

	logrus.Infof("Listening on port %s", config.ServicePort)
	api.Serve(config)
}
