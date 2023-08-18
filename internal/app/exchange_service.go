package app

import (
	"context"

	"github.com/shopspring/decimal"

	"noneland/backend/interview/pkg"
)

//go:generate mockery --name ExchangeQryService --filename exchange_service.go --output ../mocks --with-expecter --quiet

type ExchangeQryService interface {
	GetBalanceByUserId(ctx context.Context, usrId string) (BalanceResponse, error)

	GetTransactionListByUserId(
		ctx context.Context, userId string, dtoPage pkg.PageParam, tRange pkg.TimestampRangeEndTimeLessThanEqual,
	) (
		pkg.ListResponse[TransactionResponse], error,
	)
}

type BalanceResponse struct {
	SpotFee    decimal.Decimal `json:"spot_fee"`
	FuturesFee decimal.Decimal `json:"futures_fee"`
}

type TransactionResponse struct {
	Amount    decimal.Decimal `json:"amount"`
	Asset     string          `json:"asset"`
	Status    string          `json:"status"`
	Timestamp int             `json:"timestamp"`
	TxId      int64           `json:"txId"`
	Type      string          `json:"type"`
}
