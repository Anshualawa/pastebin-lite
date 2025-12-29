package main

import (
	"fmt"
	"log"
	"os"
	"pastebin-lite/handlers"
	"pastebin-lite/store"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println("RedisURL:", os.Getenv("REDIS_URL"))
	if err := store.InitRedis(); err != nil {
		log.Fatalf("redis connection failed: %v", err)
	}
	log.Println("redis connected")

	r := gin.Default()

	r.GET("/api/healthz", handlers.HealthCheck)

	r.POST("/api/pastes", handlers.CreatePaste)

	r.GET("/api/pastes/:id", handlers.GetPaste)

	r.GET("/p/:id", handlers.ViewPaste)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "9002"
	}
	log.Println("server running on :" + port)
	r.Run(":" + port)
}
