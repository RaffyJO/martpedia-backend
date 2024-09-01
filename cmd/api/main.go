package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

// NewServer initializes a new Fiber application with configuration settings
func NewServer() *fiber.App {
	return fiber.New(fiber.Config{
		IdleTimeout:  time.Second * 60,
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
	})
}

func main() {
	server := InitializedServer()
	if err := server.Listen("localhost:3000"); err != nil {
		log.Fatal(err)
	}
}
