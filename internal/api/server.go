package api

import (
	"github.com/gin-gonic/gin"

	"noneland/backend/interview/configs"
)

func NewGin(
	cfg *configs.Config,
	userH *UserHandler,
) *gin.Engine {
	if cfg.Mode == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	apiGroup := r.Group("/api")

	// TODO: api router
	apiGroup.GET("hello", userH.Hello)

	return r
}
