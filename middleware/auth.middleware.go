package middleware

import (
	"github.com/KuroNeko6666/sc-backend/config"
	"github.com/KuroNeko6666/sc-backend/database"
	"github.com/KuroNeko6666/sc-backend/handler"
	"github.com/KuroNeko6666/sc-backend/helper"
	"github.com/KuroNeko6666/sc-backend/interface/model"
	"github.com/gofiber/fiber/v2"
)

func AuthUser(c *fiber.Ctx) error {
	var user model.User

	token := c.Cookies("token", "")
	if token == "" {
		return handler.UnAuthorized(c)
	}

	if err := helper.GetUserFromToken(token, config.SecretKeyApp, &user); err != nil {
		return handler.UnAuthorized(c)
	}

	if row := database.Client.Model(&user).First(&user).RowsAffected; row == 0 {
		return handler.UnAuthorized(c)
	}

	return c.Next()
}

func AuthAdmin(c *fiber.Ctx) error {
	var admin model.Admin

	token := c.Cookies("token", "")
	if token == "" {
		return handler.UnAuthorized(c)
	}

	if err := helper.GetAdminFromToken(token, config.SecretKeyApp, &admin); err != nil {
		return handler.UnAuthorized(c)
	}

	if row := database.Client.Model(&admin).First(&admin).RowsAffected; row == 0 {
		return handler.UnAuthorized(c)
	}

	return c.Next()
}
