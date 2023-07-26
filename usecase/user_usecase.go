package usecase

import (
	"errors"

	"github.com/HiteshKumarMeghwar/L-M-S/domain"
	"github.com/HiteshKumarMeghwar/L-M-S/model"
)

type userUsecase struct {
	userRepo domain.UserRepo
}

// GetUserById implements domain.UserUsecase
func (n *userUsecase) GetUserById(id int) (model.User, error) {
	res, err := n.userRepo.GetUserById(id)
	if err != nil {
		return model.User{}, errors.New("Internal Server Error: " + err.Error())
	}
	return res, nil
}

// createuser implements domain.userusecase
func (n *userUsecase) CreateUser(createUser model.User) error {
	err := n.userRepo.CreateUser(createUser)
	return errors.New("Internal Server Error: " + err.Error())
}

func NewUserUsecase(userRepo domain.UserRepo) domain.UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}
