package handler

import (
	"errors"
	"strings"

	"github.com/KuroNeko6666/sc-backend/database"
	"github.com/KuroNeko6666/sc-backend/interface/model"
	"github.com/gofiber/fiber/v2"
)

func UserGetItemCart(c *fiber.Ctx) error {
	var cart model.Cart

	userId := c.Query("user_id")

	if err := database.Client.Model(&cart).
		Preload("Items").Preload("Items.Address").Preload("Items.Users").Preload("Items.Data").
		Where("user_id = ?", userId).First(&cart).Error; err != nil {
		return InternalServerData(c, err.Error())
	}
	return SuccessData(c, cart)
}
func UserAddItemToCart(c *fiber.Ctx) error {
	var cart model.Cart
	var device model.Device

	deviceId := c.Query("device_id")
	userId := c.Query("user_id")

	if err := database.Client.Model(&device).
		Preload("Address").Preload("Users").Preload("Data").
		Where("id = ?", deviceId).First(&device).Error; err != nil {
		return InternalServerData(c, err.Error())
	}
	if err := database.Client.Model(&cart).Where("user_id = ?", userId).First(&cart).Error; err != nil {
		return InternalServerData(c, err.Error())
	}
	if err := database.Client.Debug().Model(&cart).Association("Items").Append(&device); err != nil {
		return InternalServerData(c, err.Error())
	}

	return Success(c)
}

func UserRemoveItemToCart(c *fiber.Ctx) error {
	var cart model.Cart
	var device model.Device

	deviceId := c.Query("device_id")
	userId := c.Query("user_id")
	usage := c.Query("usage")
	usage = strings.ToUpper(usage)

	switch usage {
	case "ONE":
		if err := database.Client.Model(&device).
			Preload("Address").Preload("Users").Preload("Data").
			Where("id = ?", deviceId).Where("id = ?", deviceId).First(&device).Error; err != nil {
			return InternalServerData(c, err.Error())
		}
		if err := database.Client.Model(&cart).Where("user_id = ?", userId).First(&cart).Error; err != nil {
			return InternalServerData(c, err.Error())
		}
		if err := database.Client.Debug().Model(&cart).Association("Items").Delete(&device); err != nil {
			return InternalServerData(c, err.Error())
		}
	case "ALL":
		if err := database.Client.Model(&cart).Where("user_id = ?", userId).First(&cart).Error; err != nil {
			return InternalServerData(c, err.Error())
		}
		if err := database.Client.Debug().Model(&cart).Association("Items").Clear(); err != nil {
			return InternalServerData(c, err.Error())
		}
	default:
		err := errors.New("invalid param usage")
		return InternalServerData(c, err.Error())
	}
	return Success(c)
}
