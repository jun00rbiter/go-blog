package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jun00rbiter/go-blog/app/types/rest"
)

// GetHealthCheck はヘルスチェックのController
func GetHealthCheck(c *gin.Context) {
	// 無条件でstatus:OKを返す
	res := rest.HealthCheckResponse{
		Status: "OK",
	}
	c.SecureJSON(http.StatusOK, res)
}
