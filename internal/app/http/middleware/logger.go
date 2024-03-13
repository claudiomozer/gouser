package middleware

import (
	"github.com/claudiomozer/gouser/internal/infrastructure/logger"
	"github.com/gin-gonic/gin"
)

func LoggerMiddleware(ctx *gin.Context) {
	rCtx := logger.AddToContext(ctx.Request.Context())
	ctx.Request = ctx.Request.WithContext(rCtx)
	ctx.Next()
}
