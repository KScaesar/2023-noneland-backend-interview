package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"noneland/backend/interview/configs"
	"noneland/backend/interview/internal/entity"
	"noneland/backend/interview/internal/mocks"
	"noneland/backend/interview/pkg"
)

func TestExchangeHandler_GetSummaryBalance(t *testing.T) {
	cfg := &configs.Config{}
	mockExService := mocks.NewMockExchangeQryService(t)
	hg := HandlerGroup{
		ExchangeHandler: NewExchangeHandler(mockExService, nil),
	}
	router := NewRouter(cfg, hg)

	expectedBalance := entity.BalanceResponse{
		SpotFee:    decimal.NewFromFloat(123.456),
		FuturesFee: decimal.NewFromFloat(12.456),
	}

	// mockExService.On("GetBalanceByUserId", mock.Anything, mock.AnythingOfType("string")).
	// 	Return(expectedBalance, nil)
	mockExService.EXPECT().
		GetBalanceByUserId(mock.Anything, mock.AnythingOfType("string")).
		Return(expectedBalance, nil)

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

	ts := httptest.NewServer(router)
	defer ts.Close()
	req, err := http.NewRequest(http.MethodGet, ts.URL+"/api/v1/exchange/summary/balance", nil)
	require.NoError(t, err)

	actualBody, actualResp, err := pkg.HttpDoReturnString(http.DefaultClient, req)

	require.NoError(t, err)
	require.JSONEq(t, expectedBody, actualBody)
	require.Equal(t, http.StatusOK, actualResp.StatusCode)
}
