package xGorm

import (
	"gorm.io/gorm"

	"noneland/backend/interview/configs"
	"noneland/backend/interview/internal/entity"
)

// 原本的程式是 return interface
// 但我覺得 return struct
// 進行測試的時候, 比較容易操控某些變因
// 且進行 wire inject 的時候, 如有分 read write database
// 可以分別定義 read write struct, bind 型別進行注入
//
// 需要 repo interface 的物件
// constructor function 的 input 定義為 interface 就好

func NewUserRepository(db *gorm.DB, config *configs.Config) *UserRepository {
	return &UserRepository{
		db:     db,
		config: config,
	}
}

type UserRepository struct {
	db     *gorm.DB
	config *configs.Config
}

func (repo *UserRepository) GetUsers() (users []entity.User, err error) {
	datas := []*userGorm{}
	err = repo.db.Model(&userGorm{}).Find(&datas).Error
	if err != nil {
		return nil, err
	}

	users = make([]entity.User, len(datas))
	for i, v := range datas {
		item := userGormToEntity(v)
		if item != nil {
			users[i] = *item
		}
	}

	return nil, err
}
