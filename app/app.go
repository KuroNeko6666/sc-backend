package app

import (
	"log"

	"github.com/KuroNeko6666/sc-backend/config"
	"github.com/KuroNeko6666/sc-backend/database"
	"github.com/KuroNeko6666/sc-backend/router"
	"github.com/gofiber/fiber/v2"
)

func RunApp() {

	app := fiber.New()
	database.ConnectDB()
	router.Router(app)

	log.Fatal(app.Listen(config.ServerHost + config.ServerPort))

}
