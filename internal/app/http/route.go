package http

import (
	"github.com/claudiomozer/gouser/internal/app"
	"github.com/claudiomozer/gouser/internal/app/http/middleware"
	"github.com/gin-gonic/gin"
)

func BuildRoutes(engine *gin.Engine, container *app.Container) {
	v1 := engine.Group("/v1")
	user := v1.Group("/user")
	user.Use(middleware.LoggerMiddleware)
	{
		userHandler := NewUserHandler(container.UserService)
		user.POST("", userHandler.Create)
		user.GET("", userHandler.Get)
		user.PATCH("/role", userHandler.UpdateRole)
		user.DELETE("/:id", userHandler.Delete)
	}
}
