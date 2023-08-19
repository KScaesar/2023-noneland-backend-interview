package app

import (
	"context"

	"noneland/backend/interview/internal/entity"
	"noneland/backend/interview/pkg"
)

func NewTransactionBackupUseCase(backupRepo entity.TransactionBackupRepository, exQryService ExchangeQryService) *TransactionBackupUseCase {
	return &TransactionBackupUseCase{backupRepo: backupRepo, exQryService: exQryService}
}

type TransactionBackupUseCase struct {
	backupRepo   entity.TransactionBackupRepository
	exQryService ExchangeQryService // for backup use case
}

func (uc *TransactionBackupUseCase) GetSpotTransactionBackupAll(
	ctx context.Context, dto *entity.QryTransactionBackupParam,
) (
	[]entity.ExchangeTransactionResponse, error,
) {
	return uc.backupRepo.GetSpotTransactionBackupAllByUserId(ctx, dto)
}

func (uc *TransactionBackupUseCase) BackupTransactionRecord(ctx context.Context, tRange pkg.TimestampRangeEndTimeLessThan) error {
	// 無法作答
	// 因為不了解貴公司 write tx 的情境
	// 所以 read 無法給準確的方案
	// 雖然我詢問過貴公司一次問題
	// 但是回覆的答案
	// 和我想像的情境有衝突, 應該是我不熟悉貴公司的產業關係
	// 所以回覆中所說的情境
	// 我有看沒有懂, 訊息太少
	panic("implement me")
}
