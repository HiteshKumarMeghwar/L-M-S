package domain

import "github.com/HiteshKumarMeghwar/L-M-S/model"

type UserRepo interface {
	CreateUser(createUser model.User) error
	GetUserById(id int) (model.User, error)
}

type UserUsecase interface {
	CreateUser(createUser model.User) error
	GetUserById(id int) (model.User, error)
}
