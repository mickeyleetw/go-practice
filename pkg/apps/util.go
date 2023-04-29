package apps

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommonHandler struct {
	version string
}

func NewCommonHandler(version string) *CommonHandler {
	return &CommonHandler{version: version}
}

func (o *CommonHandler) Healthz(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"version": o.version,
	})
}

func (o *CommonHandler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}
