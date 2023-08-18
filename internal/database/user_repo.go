package database

import (
	"gorm.io/gorm"

	"noneland/backend/interview/configs"
	"noneland/backend/interview/internal/entity"
)

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
