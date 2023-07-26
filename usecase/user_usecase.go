package usecase

import (
	"github.com/HiteshKumarMeghwar/L-M-S/domain"
	"github.com/HiteshKumarMeghwar/L-M-S/model"
)

type userUsecase struct {
	userRepo domain.UserRepo
}

// createuser implements domain.userusecase
func (n *userUsecase) CreateUser(createUser model.User) error {
	err := n.userRepo.CreateUser(createUser)
	return err
}

func NewUserUsecase(userRepo domain.UserRepo) domain.UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}
