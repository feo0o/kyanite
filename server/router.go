package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func route(eng *gin.Engine) {
	healthzAPI := eng.Group("/healthz")
	{
		healthzAPI.HEAD("", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "OK")
		})

		healthzAPI.GET("", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "OK")
		})
	}
}
