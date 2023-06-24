package handler

import (
	"errors"
	"strings"

	"github.com/KuroNeko6666/sc-backend/database"
	"github.com/KuroNeko6666/sc-backend/helper"
	"github.com/KuroNeko6666/sc-backend/interface/model"
	"github.com/gofiber/fiber/v2"
)

// get all order
func GetAllOrderAdmin(c *fiber.Ctx) error {
	var order model.Order

	status := c.Query("status", "")
	status = strings.ToLower(status)
	limit := c.QueryInt("limit", 10)
	page := c.QueryInt("page", 1)
	offset := (page * limit) - limit
	userId := c.Query("user_id")

	if userId == "" && status == "" {
		goto all
	}
	if userId == "" && status != "" {
		goto withStatus
	}
	if userId != "" && status == "" {
		goto withoutStatus
	}

	if row := database.Client.Model(&order).
		Preload("Items").Preload("Items.Address").Preload("Items.Users").Preload("Items.Data").
		Limit(limit).Offset(offset).
		Where(database.Client.Where("user_id = ?", userId).Where("status = ?", status)).Find(&order); row.RowsAffected < 1 {
		return NotFound(c)
	}
	goto End

withStatus:
	if row := database.Client.Model(&order).
		Preload("Items").Preload("Items.Address").Preload("Items.Users").Preload("Items.Data").
		Limit(limit).Offset(offset).
		Where("status = ?", status).Find(&order); row.RowsAffected < 1 {
		return NotFound(c)
	}
	goto End
withoutStatus:
	if row := database.Client.Model(&order).
		Preload("Items").Preload("Items.Address").Preload("Items.Users").Preload("Items.Data").
		Limit(limit).Offset(offset).
		Where("user_id = ?", userId).Find(&order); row.RowsAffected < 1 {
		return NotFound(c)
	}
	goto End

all:
	if row := database.Client.Model(&order).
		Preload("Items").Preload("Items.Address").Preload("Items.Users").Preload("Items.Data").
		Limit(limit).Offset(offset).
		Find(&order); row.RowsAffected < 1 {
		return NotFound(c)
	}
	goto End

End:
	return SuccessPage(c, order, int64(limit), int64(page))
}

func UpdateOrderStatusAdmin(c *fiber.Ctx) error {
	var order model.Order

	orderId := c.Query("order_id")
	status := c.Query("status")

	ok := helper.IsStatusOrder(status)
	if !ok {
		err := errors.New("invalid status, must be waiting, accept, failure")
		return InternalServerData(c, err.Error())
	}

	if err := database.Client.Model(&order).
		Preload("Items").Preload("Items.Address").Preload("Items.Users").Preload("Items.Data").
		Where("id = ?", orderId).Update("status", status).Error; err != nil {
		return InternalServerData(c, err.Error())
	}

	return Success(c)
}
