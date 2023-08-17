package pkg

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"noneland/backend/interview/internal/pkg/errors"
)

// SetupHttp2 為了分成測試用與正式用，所以把 gin 的初始化抽出來
func SetupHttp2(engine *gin.Engine) (h http.Handler) {
	return h2c.NewHandler(engine, &http2.Server{})
}

func BindJsonRequest(c *gin.Context, obj any) bool {
	if err := c.ShouldBindJSON(obj); err != nil {
		Err := errors.Join3rdPartyWithMsg(errors.ErrInvalidParams, err, "bind json payload")
		ReplyErrorResponse(c, Err)
		return false
	}
	return ValidateDtoRequest(c, obj)
}

func BindQueryStringOrPostFormRequest(c *gin.Context, obj any) bool {
	if err := c.ShouldBind(obj); err != nil {
		Err := errors.Join3rdPartyWithMsg(errors.ErrInvalidParams, err, "bind query string or form payload")
		ReplyErrorResponse(c, Err)
		return false
	}
	return ValidateDtoRequest(c, obj)
}

func ValidateDtoRequest(c *gin.Context, obj any) bool {
	var target *validator.InvalidValidationError
	err := Validator.StructCtx(c.Request.Context(), obj)
	if err != nil {
		if errors.As(err, &target) {
			return true
		}

		Err := errors.Join3rdParty(errors.ErrInvalidParams, err)
		ReplyErrorResponse(c, Err)
		return false
	}
	return true
}

func ReplyErrorResponse(c *gin.Context, err error) {
	customError, _ := errors.ExtractCustomError(err)
	resp := &Response{
		Code:    customError.MyCode(),
		Message: err.Error(),
		Payload: struct{}{},
	}
	c.JSON(customError.HttpCode(), resp)
	c.Abort()
}

func ReplySuccessResponse(c *gin.Context, httpCode int, payload any) {
	if payload == nil {
		payload = struct{}{}
	}
	resp := &Response{
		Code:    0,
		Message: "ok",
		Payload: payload,
	}
	c.JSON(httpCode, resp)
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
	Payload any    `json:"payload"`
}
