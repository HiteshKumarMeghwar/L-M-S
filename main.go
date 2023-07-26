package main

import (
	"fmt"
	"log"

	"github.com/HiteshKumarMeghwar/L-M-S/config"
	"github.com/HiteshKumarMeghwar/L-M-S/controller"
	"github.com/HiteshKumarMeghwar/L-M-S/database"
	"github.com/HiteshKumarMeghwar/L-M-S/model"
	"github.com/HiteshKumarMeghwar/L-M-S/repo"
	"github.com/HiteshKumarMeghwar/L-M-S/router"
	"github.com/HiteshKumarMeghwar/L-M-S/usecase"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Hello, World")

	// this part we connect to database first
	// lets make env file

	// lets connecting the database mysql dan redis
	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load environment variables", err)
	}

	// mysql
	db := database.ConnectionDB(&loadConfig)
	db.AutoMigrate(&model.User{})

	// redis
	rdb := database.ConnectionRedisDb(&loadConfig)

	startServer(db, rdb)
}

func startServer(db *gorm.DB, rdb *redis.Client) {
	app := fiber.New()

	userRepo := repo.NewUserRepo(db, rdb)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userController := controller.NewUserController(userUsecase)

	routes := router.NewRouter(app, userController)

	err := routes.Listen(":3400")
	if err != nil {
		panic(err)
	}
}
