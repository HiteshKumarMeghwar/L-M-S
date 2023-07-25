package main

import (
	"fmt"
	"log"

	"github.com/HiteshKumarMeghwar/L-M-S/config"
	"github.com/HiteshKumarMeghwar/L-M-S/database"
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
	rdb := database.ConnectionRedisDb(&loadConfig)

	startServer(db, rdb)
}

func startServer(db *gorm.DB, rdb *redis.Client) {
	app := fiber.New()

	err := app.Listen(":3400")
	if err != nil {
		panic(err)
	}
}
