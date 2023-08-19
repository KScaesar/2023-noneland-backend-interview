package entity

import (
	"context"
)

type TransactionBackupRepository interface {
	GetSpotTransactionBackupAllByUserId(ctx context.Context, dto *QryTransactionBackupParam) ([]ExchangeTransactionResponse, error)
	CreatBulkTransactionBackup(ctx context.Context, txAll []TransactionBackup) error
}

func NewTransactionBackup(userId string, vm *ExchangeTransactionResponse) TransactionBackup {
	return TransactionBackup{
		TxId:      vm.TxId,
		UserId:    userId,
		Amount:    vm.Amount,
		Asset:     vm.Asset,
		Status:    vm.Status,
		Timestamp: vm.Timestamp,
		Type:      vm.Type,
	}
}

type TransactionBackup struct {
	TxId      int64  `gorm:"column:tx_id"`
	UserId    string `gorm:"column:user_id"`
	Amount    string `gorm:"column:amount"`
	Asset     string `gorm:"column:asset"`
	Status    string `gorm:"column:status"`
	Timestamp int    `gorm:"column:timestamp"`
	Type      string `gorm:"column:type"`
}
