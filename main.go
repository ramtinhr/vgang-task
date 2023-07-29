package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/ramtinhr/vgang-task/service"
	"github.com/sirupsen/logrus"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading env file")
	}
}

func main() {

	config := &service.Config{
		ServicePort: os.Getenv("PORT"),
		ServiceName: os.Getenv("NAME"),
		Version:     os.Getenv("VERSION"),
		Env:         os.Getenv("ENV"),
	}

	logrus.Infof("Listening on port %s", config.ServicePort)
	service.Serve(config)
}
