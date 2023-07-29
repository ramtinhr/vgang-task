package service

import "github.com/sirupsen/logrus"

type Config struct {
	Env         string         `required:"true"`
	ServiceName string         `required:"true"`
	ServicePort string         `required:"true"`
	Version     string         `required:"true"`
	Logger      *logrus.Logger `ignored:"true"`
}
