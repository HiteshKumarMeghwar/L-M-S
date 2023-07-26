package repo

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/HiteshKumarMeghwar/L-M-S/domain"
	"github.com/HiteshKumarMeghwar/L-M-S/model"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type userRepo struct {
	db  *gorm.DB
	rdb *redis.Client
}

// GetUserById implements domain.UserRepo
func (n *userRepo) GetUserById(id int) (model.User, error) {
	var users model.User
	var ctx = context.Background()

	// first check data is avaiable in redis
	result, err := n.rdb.Get(ctx, "user"+strconv.Itoa(id)).Result()
	if err != nil && err != redis.Nil {
		return users, err
	}

	// if data avaiable in redis, decode it from JSON, and return it
	if len(result) > 0 {
		err := json.Unmarshal([]byte(result), &users)
		return users, err
	}

	// if data was not avaiable in redis, get it from database
	err = n.db.Model(model.User{}).Select(
		"id",
		"name",
		"description",
		"author",
	).Where("id=?", id).Find(&users).Error
	if err != nil {
		return users, err
	}

	// encode that slice into json before saving into redis
	jsonBytes, err := json.Marshal(users)
	if err != nil {
		return users, err
	}
	jsonString := string(jsonBytes)

	// set the json-encoded value in redis
	err = n.rdb.Set(ctx, "user"+strconv.Itoa(id), jsonString, 24*time.Hour).Err()
	if err != nil {
		return users, err
	}

	return users, nil
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
