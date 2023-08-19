package external

import (
	"context"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"noneland/backend/interview/configs"
	"noneland/backend/interview/internal/entity"
	"noneland/backend/interview/pkg"
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

	expectedResp := entity.BalanceResponse{
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

func TestHttpExchangeQryService_GetTransactionListByUserId(t *testing.T) {
	client := http.DefaultClient
	cfg := configs.NewConfig("template-dev")

	ctx := context.Background()
	usrId := ""
	tRange := pkg.TimestampRangeEndTimeLessThan{
		EndTime: pkg.MockTimeNow("2023-08-19T12:00:00Z")().UnixMilli(),
	}
	page := pkg.PageParam{
		Page: 2,
		Size: 123,
	}
	service := NewHttpExchangeQryService(client, cfg)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder(
		http.MethodGet,
		"=~/spot/transfer/records.?",
		httpmock.NewStringResponder(http.StatusOK, `
{
   "rows": [
      {
         "amount": "0.10000000",
         "asset": "BNB",
         "status": "CONFIRMED",
         "timestamp": 1566898617,
         "txId": 5240372201,
         "type": "IN"
      },
      {
         "amount": "5.00000000",
         "asset": "USDT",
         "status": "CONFIRMED",
         "timestamp": 1566888436,
         "txId": 5239810406,
         "type": "OUT"
      },
      {
         "amount": "1.00000000",
         "asset": "EOS",
         "status": "CONFIRMED",
         "timestamp": 1566888403,
         "txId": 5239808703,
         "type": "IN"
      }
   ],
   "total": 3
}
`))

	want := pkg.ListResponse[entity.TransactionResponse]{
		Rows: []entity.TransactionResponse{
			{
				Amount:    "0.10000000",
				Asset:     "BNB",
				Status:    "CONFIRMED",
				Timestamp: 1566898617,
				TxId:      5240372201,
				Type:      "IN",
			},
			{
				Amount:    "5.00000000",
				Asset:     "USDT",
				Status:    "CONFIRMED",
				Timestamp: 1566888436,
				TxId:      5239810406,
				Type:      "OUT",
			},
			{
				Amount:    "1.00000000",
				Asset:     "EOS",
				Status:    "CONFIRMED",
				Timestamp: 1566888403,
				TxId:      5239808703,
				Type:      "IN",
			},
		},
		Total: 3,
	}

	got, err := service.GetTransactionListByUserId(ctx, usrId, page, tRange)

	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestHttpExchangeQryService_transformQueryString(t *testing.T) {
	client := http.DefaultClient
	cfg := configs.NewConfig("template-dev")

	tRange := pkg.TimestampRangeEndTimeLessThan{
		EndTime: pkg.MockTimeNow("2023-08-19T12:00:00Z")().UnixMilli(),
	}
	page := pkg.PageParam{
		Page: 2,
		Size: 123,
	}
	service := NewHttpExchangeQryService(client, cfg)

	got := service.transformQueryString(page, tRange)

	assert.Equal(t, "2", got.Get("current"))
	assert.Equal(t, "100", got.Get("size"), "max size only 100")
	assert.Equal(t, false, got.Has("startTime"))
	assert.Equal(t, "1692446400000", got.Get("endTime"))
}
