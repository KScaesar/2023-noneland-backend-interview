package pkg

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"noneland/backend/interview/configs"
	"noneland/backend/interview/internal/api"
)

// SetupHttp2 為了分成測試用與正式用，所以把 gin 的初始化抽出來
func SetupHttp2(engine *gin.Engine) (h http.Handler) {
	return h2c.NewHandler(engine, &http2.Server{})
}

func NewGin(
	cfg *configs.Config,
	userH *api.UserHandler,
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
