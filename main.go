package main

import (
	"apiweb/config"
	"apiweb/routing"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func main() {

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000, http://localhost:8088, https://*.isoide.co.id/",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, OPTIONS",
	}))
	config.ConnectDb()
	config.ConnectDbmaster()
	routing.Routes(app)

	// Corrected listen address
	log.Fatal(app.Listen(":8088"))
}
