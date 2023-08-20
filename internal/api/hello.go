package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"noneland/backend/interview/internal/app"
	"noneland/backend/interview/pkg"
)

func NewUserHandler(svc *app.ApplicationGroup) *UserHandler {
	return &UserHandler{service: svc.UserUseCase}
}

type UserHandler struct {
	service *app.UserUseCase
}

func (h *UserHandler) Hello(c *gin.Context) {
	h.service.Hello()
	pkg.ReplySuccessResponse(c, http.StatusOK, nil)
}
