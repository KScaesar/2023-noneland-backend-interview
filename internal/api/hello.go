package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"noneland/backend/interview/internal/app"
)

func NewUserHandler(service *app.UserUseCase) *UserHandler {
	return &UserHandler{service: service}
}

type UserHandler struct {
	service *app.UserUseCase
}

func (h *UserHandler) Hello(c *gin.Context) {
	h.service.Hello()
	c.JSON(http.StatusOK, okResponse{OK: true})
}
