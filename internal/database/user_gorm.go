package database

import "noneland/backend/interview/internal/entity"

// 理想很豐滿, 現實很骨感
// 如果領域模型的欄位和資料庫一致
// 我不會另外定義 data layer(userGorm) 的物件轉換
// 會選擇直接在 entity obj 加上 gorm struct tag
// 因為翻譯層寫起來太煩
// 有一定的認知, 知道商業邏輯要以 domain obj 為主就好

type userGorm struct {
	Name string `gorm:"type:varchar(255)"`
}

func userGormToEntity(input *userGorm) *entity.User {
	return &entity.User{
		Name: input.Name,
	}
}

func userEntityToGorm(input *entity.User) *userGorm {
	return &userGorm{
		Name: input.Name,
	}
}
