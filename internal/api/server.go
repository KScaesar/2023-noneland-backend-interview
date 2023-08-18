package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"noneland/backend/interview/configs"
	"noneland/backend/interview/pkg"
)

// NewRouter 可用在 httptest.NewServer 進行 integration test
func NewRouter(cfg *configs.Config, hg HandlerGroup) *gin.Engine {
	router := gin.New()
	setupServer(router, cfg)
	registerRoute(router, hg)
	return router
}

func NewServer(cfg *configs.Config, hg HandlerGroup) *http.Server {
	router := gin.New()
	server := setupServer(router, cfg)
	registerRoute(router, hg)
	return server
}

func setupServer(router *gin.Engine, cfg *configs.Config) *http.Server {
	if !cfg.DebugHttp {
		gin.SetMode(gin.ReleaseMode)
	}

	var mux http.Handler
	if cfg.EnableHttp2 {
		mux = pkg.SetupHttp2(router)
	} else {
		mux = router
	}

	s := &http.Server{
		Addr:           fmt.Sprintf(":%s", cfg.Port),
		Handler:        mux,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return s
}

func registerRoute(router *gin.Engine, hg HandlerGroup) {
	v1 := router.Group("/api/v1")

	// TODO: api router
	v1.GET("/hello", hg.UserH.Hello)
}

type HandlerGroup struct {
	UserH UserHandler
}
