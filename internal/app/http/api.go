package http

import (
	"fmt"

	"github.com/claudiomozer/gouser/internal/app"
	"github.com/gin-gonic/gin"
)

func Start() *app.Container {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	engine.ContextWithFallback = true

	container := app.LoadContainer()
	BuildRoutes(engine, container)

	runErr := engine.Run(fmt.Sprintf(":%d", app.ENV.AppPort))
	if runErr != nil {
		panic(runErr)
	}

	return container
}
