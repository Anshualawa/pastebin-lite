package main

import (
	"log"
	"os"
	"pastebin-lite/handlers"
	"pastebin-lite/store"

	"github.com/gin-gonic/gin"
)

func main() {
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
