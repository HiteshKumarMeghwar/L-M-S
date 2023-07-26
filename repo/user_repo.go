package repo

import (
	"errors"

	"github.com/HiteshKumarMeghwar/L-M-S/domain"
	"github.com/HiteshKumarMeghwar/L-M-S/model"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type userRepo struct {
	db  *gorm.DB
	rdb *redis.Client
}

// createUser implements domain.UserRepo
func (n *userRepo) CreateUser(createUser model.User) error {
	if err := n.db.Create(&createUser).Error; err != nil {
		return errors.New("Internal Server Error: cannot create user")
	}
	return nil
}

func NewUserRepo(db *gorm.DB, rdb *redis.Client) domain.UserRepo {
	return &userRepo{
		db:  db,
		rdb: rdb,
	}
}
