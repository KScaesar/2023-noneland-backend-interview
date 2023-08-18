package pkg

import (
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"noneland/backend/interview/pkg/errors"
)

func NewSqliteGorm() *gorm.DB {
	db, err := gorm.Open(
		sqlite.Open(":memory:"), &gorm.Config{
			Logger: logger.Discard,
		},
	)
	if err != nil {
		log.Fatalf("gorm.Open 無法連接到記憶體數據庫：%v", err)
	}

	return db
}

func NewMySqlGorm() *gorm.DB {
	return nil
}

func GormError(err error) error {
	// var pgErr *pgconn.PgError
	// if errors.As(err, &pgErr) {
	// 	return PgsqlError(pgErr)
	// }

	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return errors.ErrNotFound
	default:
		return errors.Join3rdParty(errors.ErrSystem, err)
	}
}
