package app

import (
	"log"

	"github.com/KuroNeko6666/sc-backend/config"
	"github.com/KuroNeko6666/sc-backend/database"
	"github.com/KuroNeko6666/sc-backend/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func RunApp() {

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOriginsFunc: func(origin string) bool {
			return true
		},
		AllowCredentials: true,
		AllowOrigins:     "*",
		AllowHeaders:     "*, content-type, authorization",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS, PATCH",
	}))
	database.ConnectDB()
	router.Router(app)

	log.Fatal(app.Listen(config.ServerHost + config.ServerPort))

}
