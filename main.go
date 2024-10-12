package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println(".env file not loaded, continuing with existing environment variables.")
	}

	ctx := context.Background()

	redisDb, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		panic(err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       redisDb,
	})

	pong, err := rdb.Ping(ctx).Result()

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to redis: " + pong)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_SSLMODE"), os.Getenv("POSTGRES_TIMEZONE"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.WithContext(ctx).Select(1)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected to postgres.")
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
