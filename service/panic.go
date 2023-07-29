package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PanicRecovery Gin panic recovery and send 500 error
func (config *Config) PanicRecovery(c *gin.Context, recovered interface{}) {
	if err, ok := recovered.(string); ok {
		config.Logger.Error(err)
		c.JSON(http.StatusInternalServerError, ErrorResponse{Code: http.StatusInternalServerError, Cause: "Something went wrong"})
		return
	}
	c.AbortWithStatus(http.StatusInternalServerError)
}
