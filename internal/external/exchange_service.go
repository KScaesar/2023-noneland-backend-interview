package external

import (
	"context"
	"net/http"
	"net/url"

	"github.com/avast/retry-go/v4"
	"github.com/google/go-querystring/query"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"

	"noneland/backend/interview/configs"
	"noneland/backend/interview/internal/entity"
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

func (svc *HttpExchangeQryService) GetBalanceByUserId(ctx context.Context, usrId string) (resp entity.BalanceResponse, Err error) {
	type balance struct {
		Free decimal.Decimal `json:"free"`
	}

	mqSpot := make(chan lo.Tuple2[decimal.Decimal, error], 1)
	mqFuture := make(chan lo.Tuple2[decimal.Decimal, error], 1)
	mqAll := []chan lo.Tuple2[decimal.Decimal, error]{mqSpot, mqFuture}

	urlAll := []string{svc.url + "/spot/balance", svc.url + "/futures/balance"}

	qryCount := len(urlAll)
	for i := 0; i < qryCount; i++ {
		go func(i int) {
			var payload balance
			var qryErr error
			qryErr = retry.Do(
				func() error {
					req, err := http.NewRequest(http.MethodGet, urlAll[i], nil)
					if err != nil {
						return errors.Join3rdParty(errors.ErrSystem, err)
					}
					payload, err = pkg.HttpDoReturnType[balance](svc.client, req)
					return err
				},
				retry.Attempts(3),
			)
			mqAll[i] <- lo.Tuple2[decimal.Decimal, error]{A: payload.Free, B: qryErr}
		}(i)
	}

	for i := 0; i < qryCount; i++ {
		select {
		case spot := <-mqSpot:
			if spot.B != nil {
				return entity.BalanceResponse{}, spot.B
			}
			resp.SpotFee = spot.A

		case future := <-mqFuture:
			if future.B != nil {
				return entity.BalanceResponse{}, future.B
			}
			resp.FuturesFee = future.A

		case <-ctx.Done():
			return entity.BalanceResponse{}, errors.Join3rdParty(errors.ErrSystem, ctx.Err())
		}
	}
	return
}

func (svc *HttpExchangeQryService) GetTransactionListByUserId(
	_ context.Context, userId string, dtoPage pkg.PageParam, tRange pkg.TimestampRangeEndTimeLessThan,
) (
	resp pkg.ListResponse[entity.ExchangeTransactionResponse], Err error,
) {

	qs := svc.transformQueryString(dtoPage, tRange)
	qryErr := retry.Do(
		func() error {
			req, err := http.NewRequest(http.MethodGet, svc.url+"/spot/transfer/records?"+qs.Encode(), nil)
			if err != nil {
				return errors.Join3rdParty(errors.ErrSystem, err)
			}
			resp, err = pkg.HttpDoReturnType[pkg.ListResponse[entity.ExchangeTransactionResponse]](svc.client, req)
			return err
		},
		retry.Attempts(3),
	)
	if qryErr != nil {
		return pkg.ListResponse[entity.ExchangeTransactionResponse]{}, errors.WrapWithMessage(qryErr, "call 3rd exchange api")
	}
	return
}

func (svc *HttpExchangeQryService) transformQueryString(page pkg.PageParam, tRange pkg.TimestampRangeEndTimeLessThan) url.Values {
	type QueryString struct {
		StartTime int64  `url:"startTime,omitempty"`
		EndTime   int64  `url:"endTime,omitempty"`
		Current   uint64 `url:"current,omitempty"`
		Size      uint64 `url:"size,omitempty"`
	}

	page.SetDefaultAndMaxSizeIfInvalid(10, 100)
	qs := QueryString{
		StartTime: tRange.StartTime,
		EndTime:   tRange.EndTime,
		Current:   page.Page,
		Size:      page.Size,
	}
	values, _ := query.Values(qs)
	return values
}
