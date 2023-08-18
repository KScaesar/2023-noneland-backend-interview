package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"noneland/backend/interview/internal/app"
	"noneland/backend/interview/pkg"
)

func NewExchangeHandler(svc app.ExchangeQryService) ExchangeHandler {
	return ExchangeHandler{svc: svc}
}

type ExchangeHandler struct {
	svc app.ExchangeQryService
}

func (h *ExchangeHandler) GetSummaryBalance(c *gin.Context) {
	resp, err := h.svc.GetBalanceByUserId(c.Request.Context(), "")
	if err != nil {
		pkg.ReplyErrorResponse(c, err)
		return
	}
	pkg.ReplySuccessResponse(c, http.StatusOK, resp)
}
