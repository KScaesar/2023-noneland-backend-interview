package app

import (
	"noneland/backend/interview/internal/entity"
)

func NewUserUseCase(repo entity.UserRepository) *UserUseCase {
	return &UserUseCase{repo: repo}
}

type UserUseCase struct {
	repo entity.UserRepository
}

func (uc *UserUseCase) Hello() {
	// use repo sample
	// repo, err := di.NewRepo()
	// if err != nil {
	//	errResponse(c)
	//	return
	// }
	// users, err := repo.GetUsers()
}
