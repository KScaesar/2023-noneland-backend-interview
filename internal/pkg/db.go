package pkg

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
