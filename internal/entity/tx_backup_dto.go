package entity

import (
	"noneland/backend/interview/pkg"
)

type QryTransactionBackupParam struct {
	FilterTransactionBackupParam
	// PageParam
	// SortParam
}

func (q *QryTransactionBackupParam) SetUserId(userId string) {
	q.UserId = userId
}

type FilterTransactionBackupParam struct {
	UserId string `json:"user_id" form:"user_id"`
	pkg.TimestampRangeEndTimeLessThanEqual
}
