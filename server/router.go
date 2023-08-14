package server

import (
	"fmt"

	"github.com/feo0o/kyanite/api"
	"github.com/feo0o/kyanite/app"
	docs "github.com/feo0o/kyanite/docs"
	"github.com/feo0o/kyanite/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func route(eng *gin.Engine) {
	docs.SwaggerInfo.Title = fmt.Sprintf("%s API Docs", app.Name)
	docs.SwaggerInfo.Description = "A simple HTTP(S) server with Ops tools."
	docs.SwaggerInfo.Version = app.Version()
	docs.SwaggerInfo.BasePath = "/"

	healthzAPI := eng.Group("/healthz", middleware.DelayBySecond())
	{
		healthzAPI.HEAD("", api.HealthChechk)
		healthzAPI.GET("", api.HealthChechk)
	}

	randomAPI := eng.Group("/random")
	{
		randomAPI.HEAD("/string", api.GetRandomString)
		randomAPI.GET("/string", api.GetRandomString)

		randomAPI.HEAD("/int", api.GetRandomInt)
		randomAPI.GET("/int", api.GetRandomInt)

		randomAPI.HEAD("/decimal", api.GetRandomdecimal)
		randomAPI.GET("/decimal", api.GetRandomdecimal)
	}

	// Docs API
	eng.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
