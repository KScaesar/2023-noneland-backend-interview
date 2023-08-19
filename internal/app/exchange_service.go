package app

import (
	"context"

	"noneland/backend/interview/internal/entity"
	"noneland/backend/interview/pkg"
)

type ExchangeQryService interface {
	GetBalanceByUserId(ctx context.Context, usrId string) (entity.BalanceResponse, error)

	GetTransactionListByUserId(
		ctx context.Context, userId string, dtoPage pkg.PageParam, tRange pkg.TimestampRangeEndTimeLessThan,
	) (
		pkg.ListResponse[entity.TransactionResponse], error,
	)
}
