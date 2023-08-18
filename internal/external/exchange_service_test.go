package external

import (
	"context"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/require"

	"noneland/backend/interview/configs"
	"noneland/backend/interview/internal/app"
)

func TestHttpExchangeQryService_GetBalanceByUserId(t *testing.T) {
	client := http.DefaultClient
	cfg := configs.NewConfig("template-dev")

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		cfg.ExchangeUrl+"/spot/balance",
		httpmock.NewStringResponder(http.StatusOK, `{ "free": "10.12345" }`),
	)
	httpmock.RegisterResponder(
		http.MethodGet,
		cfg.ExchangeUrl+"/futures/balance",
		httpmock.NewStringResponder(http.StatusOK, `{ "free": "10.145" }`),
	)

	//

	ctx := context.Background()
	usrId := ""
	service := NewHttpExchangeQryService(client, cfg)

	//

	expectedResp := app.BalanceResponse{
		SpotFee:    decimal.NewFromFloat32(10.12345),
		FuturesFee: decimal.NewFromFloat32(10.145),
	}
	expectedCallCount := 2

	//

	actualResp, err := service.GetBalanceByUserId(ctx, usrId)
	actualCallCount := httpmock.GetTotalCallCount()

	//

	require.NoError(t, err)
	require.Equal(t, expectedResp, actualResp)
	require.Equal(t, expectedCallCount, actualCallCount)
}
