package server

import (
	"fmt"

	"github.com/feo0o/kyanite/api"
	"github.com/feo0o/kyanite/app"
	docs "github.com/feo0o/kyanite/docs"
	"github.com/feo0o/kyanite/middleware"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func route(e *gin.Engine) {
	docs.SwaggerInfo.Title = fmt.Sprintf("%s API Docs", app.Name)
	docs.SwaggerInfo.Description = "A simple HTTP(S) server with Ops tools."
	docs.SwaggerInfo.Version = app.Version()
	docs.SwaggerInfo.BasePath = "/"

	healthzAPI := e.Group("/healthz", middleware.DelayBySecond())
	{
		healthzAPI.HEAD("", api.HealthChechk)
		healthzAPI.GET("", api.HealthChechk)
	}

	randomAPI := e.Group("/random")
	{
		randomAPI.HEAD("/string", api.GetRandomString)
		randomAPI.GET("/string", api.GetRandomString)

		randomAPI.HEAD("/int", api.GetRandomInt)
		randomAPI.GET("/int", api.GetRandomInt)

		randomAPI.HEAD("/decimal", api.GetRandomdecimal)
		randomAPI.GET("/decimal", api.GetRandomdecimal)
	}

	// Metrics
	e.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Docs API
	e.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
