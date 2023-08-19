package entity

import (
	"github.com/shopspring/decimal"
)

type BalanceResponse struct {
	SpotFee    decimal.Decimal `json:"spot_fee"`
	FuturesFee decimal.Decimal `json:"futures_fee"`
}

type TransactionResponse struct {
	TxId      int64  `json:"txId"`
	Amount    string `json:"amount"`
	Asset     string `json:"asset"`
	Status    string `json:"status"`
	Timestamp int    `json:"timestamp"`
	Type      string `json:"type"`
}
