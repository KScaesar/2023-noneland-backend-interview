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
	v1 := r.Group("/api/v1")

	// TODO: api router
	v1.GET("/hello", userH.Hello)

	return r
}
