package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck health check
// @Summary health check
// @Description health check, no params needed
// @Tags healthz
// @Produce plain
// @Success 200 {string} string "OK"
// @Router /healthz [get]
func HealthChechk(c *gin.Context) {
	// todo: check detail
	c.String(http.StatusOK, "OK\n")
}
