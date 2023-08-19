package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"noneland/backend/interview/internal/app"
	"noneland/backend/interview/internal/entity"
	"noneland/backend/interview/pkg"
)

func NewExchangeHandler(exService app.ExchangeQryService, backupService *app.TransactionBackupUseCase) ExchangeHandler {
	return ExchangeHandler{exService: exService, backupService: backupService}
}

type ExchangeHandler struct {
	exService     app.ExchangeQryService
	backupService *app.TransactionBackupUseCase
}

func (h *ExchangeHandler) GetSummaryBalance(c *gin.Context) {
	resp, err := h.exService.GetBalanceByUserId(c.Request.Context(), "")
	if err != nil {
		pkg.ReplyErrorResponse(c, err)
		return
	}
	pkg.ReplySuccessResponse(c, http.StatusOK, resp)
}

func (h *ExchangeHandler) GetSpotTransactionRecordAll(c *gin.Context) {
	var dto entity.QryTransactionBackupParam
	if pkg.BindQueryStringOrPostFormRequest(c, &dto) {
		return
	}

	ctx := c.Request.Context()
	dto.SetUserId("userId from token")

	list, err := h.backupService.GetSpotTransactionBackupAll(ctx, &dto)
	if err != nil {
		pkg.ReplyErrorResponse(c, err)
		return
	}
	pkg.ReplySuccessResponse(c, http.StatusOK, list)
}
