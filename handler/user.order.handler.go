package handler

import (
	"github.com/KuroNeko6666/sc-backend/database"
	"github.com/KuroNeko6666/sc-backend/interface/model"
	"github.com/KuroNeko6666/sc-backend/interface/response"
	"github.com/gofiber/fiber/v2"
)

func UserCreateOrder(c *fiber.Ctx) error {
	var cart model.Cart
	var order model.Order
	var res response.Order

	userId := c.Query("user_id")

	// Get Cart data
	if err := database.Client.Model(&cart).
		Preload("Items").Preload("Items.Address").Preload("Items.Users").Preload("Items.Data").Where("user_id = ?", userId).
		First(&cart).Error; err != nil {
		return InternalServerData(c, err.Error())
	}

	order.UserID = cart.UserID
	order.Items = cart.Items
	order.Status = "waiting"

	if err := database.Client.Model(&order).
		Preload("Users").
		Create(&order).Error; err != nil {
		return InternalServerData(c, err.Error())
	}

	res.ID = order.ID
	res.Status = order.Status
	res.Items = order.Items
	res.CreatedAt = order.CreatedAt
	res.UpdatedAt = order.UpdatedAt

	if err := database.Client.Model(&cart).Association("Items").Clear(); err != nil {
		return InternalServerData(c, err.Error())
	}
	return Success(c)
}

func UserGetItemOrder(c *fiber.Ctx) error {
	var order []model.Order

	limit := c.QueryInt("limit", 10)
	page := c.QueryInt("page", 1)
	offset := (page * limit) - limit
	userId := c.Query("user_id")

	if err := database.Client.Model(&order).
		Preload("Items").Preload("Items.Address").Preload("Items.Users").Preload("Items.Data").
		Limit(limit).Offset(offset).
		Where("user_id = ?", userId).Find(&order).Error; err != nil {
		return InternalServerData(c, err.Error())
	}

	return SuccessPage(c, order, int64(limit), int64(page))
}
