package http

import (
	"errors"
	"net/http"

	"github.com/claudiomozer/gouser/internal/domain/err"
	"github.com/claudiomozer/gouser/internal/domain/user"
	"github.com/claudiomozer/gouser/internal/infrastructure/logger"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *user.Service
}

func NewUserHandler(service *user.Service) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) Create(ctx *gin.Context) {
	var request *user.CreateRequest
	bindErr := ctx.ShouldBindJSON(&request)
	if bindErr != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "invalid body",
		})
		logger.FromContext(ctx).Error("create-user-error", "err", bindErr)
		return
	}

	createErr := h.service.Create(ctx, request)
	if createErr != nil {
		logger.FromContext(ctx).Error("create-user-error", "err", createErr)
		var domainErr *err.Error
		if errors.As(createErr, &domainErr) {
			var statusCode int
			switch domainErr.Code() {
			case err.ErrUserAlreadyExists:
				statusCode = http.StatusConflict
			default:
				statusCode = http.StatusBadRequest
			}
			ctx.AbortWithStatusJSON(statusCode, gin.H{
				"code":    domainErr.Code(),
				"message": domainErr.Message(),
			})
			return
		}
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}
}
