package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Serve manage all api routes here
func Serve(config *Config) {
	mode := gin.ReleaseMode
	if config.Env == "local" {
		mode = gin.DebugMode
	}
	gin.SetMode(mode)

	r := gin.New()

	r.RedirectTrailingSlash = false
	r.RedirectFixedPath = false

	r.Use(gin.CustomRecovery(config.PanicRecovery))

	apiGroup := r.Group(fmt.Sprintf("/api/v%s", config.Version))
	apiGroup.GET("/products")

	apiGroup.GET("/ping", func(c *gin.Context) {
		var msg string
		if config.Env == "production" {
			msg = fmt.Sprintf("Pong! The %s Application Version is %s.", config.ServiceName, config.Version)
		} else {
			msg = "pong"
		}
		c.JSON(200, gin.H{"message": msg})
	})

	r.HandleMethodNotAllowed = true
	if err := r.Run(fmt.Sprintf(":%s", config.ServicePort)); err != nil {
		logrus.Fatalf("api listening problem: %s", err)
	}
}
