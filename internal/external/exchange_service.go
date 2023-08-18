package external

import (
	"context"
	"net/http"

	"github.com/avast/retry-go/v4"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"

	"noneland/backend/interview/configs"
	"noneland/backend/interview/internal/app"
	"noneland/backend/interview/pkg"
	"noneland/backend/interview/pkg/errors"
)

func NewHttpExchangeQryService(client *http.Client, cfg *configs.Config) *HttpExchangeQryService {
	return &HttpExchangeQryService{
		client: client,
		url:    cfg.ExchangeUrl,
	}
}

type HttpExchangeQryService struct {
	client *http.Client
	url    string
}

func (svc *HttpExchangeQryService) GetBalanceByUserId(ctx context.Context, usrId string) (resp app.BalanceResponse, Err error) {
	type balance struct {
		Free decimal.Decimal `json:"free"`
	}

	mqSpot := make(chan lo.Tuple2[decimal.Decimal, error], 1)
	mqFuture := make(chan lo.Tuple2[decimal.Decimal, error], 1)

	go func() {
		var payload balance
		var err error
		err = retry.Do(
			func() error {
				req, err := http.NewRequest(http.MethodGet, svc.url+"/spot/balance", nil)
				if err != nil {
					return err
				}
				payload, err = pkg.HttpDoReturnType[balance](svc.client, req)
				return err
			},
			retry.Attempts(3),
		)
		mqSpot <- lo.Tuple2[decimal.Decimal, error]{A: payload.Free, B: err}
	}()
	go func() {
		var payload balance
		var err error
		err = retry.Do(
			func() error {
				req, err := http.NewRequest(http.MethodGet, svc.url+"/futures/balance", nil)
				if err != nil {
					return err
				}
				payload, err = pkg.HttpDoReturnType[balance](svc.client, req)
				return err
			},
			retry.Attempts(3),
		)
		mqFuture <- lo.Tuple2[decimal.Decimal, error]{A: payload.Free, B: err}
	}()

	qryCount := 2
	for i := 0; i < qryCount; i++ {
		select {
		case spot := <-mqSpot:
			if spot.B != nil {
				return app.BalanceResponse{}, spot.B
			}
			resp.SpotFee = spot.A

		case future := <-mqFuture:
			if future.B != nil {
				return app.BalanceResponse{}, future.B
			}
			resp.FuturesFee = future.A

		case <-ctx.Done():
			return app.BalanceResponse{}, errors.ErrTimeout
		}
	}
	return
}

func (HttpExchangeQryService) GetTransactionListByUserId(
	ctx context.Context, userId string, dtoPage pkg.PageParam, tRange pkg.TimestampRangeEndTimeLessThanEqual,
) (pkg.ListResponse[app.TransactionResponse], error,
) {
	// TODO implement me
	panic("implement me")
}
