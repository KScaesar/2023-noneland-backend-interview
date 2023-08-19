package database

import (
	"context"

	"gorm.io/gorm"

	"noneland/backend/interview/internal/entity"
	"noneland/backend/interview/pkg"
)

func NewGormTransactionBackupRepository(db *gorm.DB) *GormTransactionBackupRepository {
	return &GormTransactionBackupRepository{db: db}
}

type GormTransactionBackupRepository struct {
	db *gorm.DB
}

func (repo *GormTransactionBackupRepository) GetSpotTransactionBackupAllByUserId(
	ctx context.Context, dto *entity.QryTransactionBackupParam,
) (
	result []entity.ExchangeTransactionResponse, err error,
) {
	where := repo.qryWhere(&dto.FilterTransactionBackupParam)

	err = repo.db.WithContext(ctx).
		Table("spot_tx_backup").
		Scopes(where...).
		Find(&result).Error
	if err != nil {
		return nil, pkg.GormError(err)
	}

	return result, nil
}

func (repo *GormTransactionBackupRepository) CreatBulkTransactionBackup(ctx context.Context, txAll []entity.TransactionBackup) error {
	err := repo.db.WithContext(ctx).
		Table("spot_tx_backup").
		Create(txAll).Error
	if err != nil {
		return pkg.GormError(err)
	}
	return nil
}

func (repo *GormTransactionBackupRepository) qryWhere(dto *entity.FilterTransactionBackupParam) []func(db *gorm.DB) *gorm.DB {
	var where []func(db *gorm.DB) *gorm.DB
	if dto.UserId != "" {
		where = append(where, func(db *gorm.DB) *gorm.DB {
			db.Where("user_id = ?", dto.UserId)
			return db
		})
	}
	if dto.StartTime != 0 {
		where = append(where, func(db *gorm.DB) *gorm.DB {
			db.Where("start_time >= ?", dto.StartTime)
			return db
		})
	}
	if dto.EndTime != 0 {
		where = append(where, func(db *gorm.DB) *gorm.DB {
			db.Where("end_time <= ?", dto.EndTime)
			return db
		})
	}
	return where
}
