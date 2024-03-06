package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"log"
	"time"
)

type app struct {
	server *fiber.App
}

func main() {

	fib := fiber.New(fiber.Config{
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	})

	application := app{
		server: fib,
	}

	err := InitDB()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := CloseDB()
		if err != nil {
			log.Fatal(err)
		}
	}()

	application.server.Use(logger.New())
	application.server.Use(recover2.New())
	application.server.Use(cors.New())

	BookGroup(application.server)

	err = application.server.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
