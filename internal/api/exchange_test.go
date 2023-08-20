package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"noneland/backend/interview/configs"
	"noneland/backend/interview/internal/api"
	"noneland/backend/interview/internal/di"
	"noneland/backend/interview/internal/entity"
	"noneland/backend/interview/internal/xMock"
	"noneland/backend/interview/pkg"
)

func TestExchangeHandler_GetSummaryBalance(t *testing.T) {
	// arrange
	cfg := configs.NewConfig("template-dev")
	cfg.DebugHttp = false

	apps := di.NewApplication(cfg)
	mockExService := xMock.NewMockExchangeQryService(t)
	apps.ExchangeQryService = mockExService // comment the row to disable mock

	handlers := di.NewHttpHandler(apps)
	router := api.NewRouter(cfg, handlers)
	ts := httptest.NewServer(router)
	defer ts.Close()
	returnValue := entity.BalanceResponse{
		SpotFee:    decimal.NewFromFloat(123.456),
		FuturesFee: decimal.NewFromFloat(12.456),
	}
	mockExService.EXPECT().
		GetBalanceByUserId(mock.Anything, mock.AnythingOfType("string")).
		Return(returnValue, nil)

	// expect
	expectedBody := `
{
  "code": 0,
  "msg": "ok",
  "payload": {
    "spot_fee": "123.456",
    "futures_fee": "12.456"
  }
}
`

	// action
	req, _ := http.NewRequest(http.MethodGet, ts.URL+"/api/v1/exchange/summary/balance", nil)
	actualBody, actualResp, err := pkg.HttpDoReturnString(http.DefaultClient, req)

	// assert
	require.NoError(t, err)
	require.JSONEq(t, expectedBody, actualBody)
	require.Equal(t, http.StatusOK, actualResp.StatusCode)
}
