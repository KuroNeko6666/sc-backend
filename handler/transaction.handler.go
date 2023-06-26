package handler

import (
	"math"
	"strings"

	"github.com/KuroNeko6666/sc-backend/config"
	"github.com/KuroNeko6666/sc-backend/database"
	"github.com/KuroNeko6666/sc-backend/helper"
	"github.com/KuroNeko6666/sc-backend/interface/form"
	"github.com/KuroNeko6666/sc-backend/interface/model"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AddDeviceToCart(c *fiber.Ctx) error {
	var cart model.Cart
	var user model.User
	var device model.Device
	deviceID := c.Params("id", "")
	var token string

	token = c.Cookies("token", "")

	if token == "" {
		token = strings.Split(c.GetReqHeaders()["Authorization"], " ")[1]
	}

	if err := helper.GetUserFromToken(token, config.SecretKeyApp, &user); err != nil {
		return BadRequestData(c, err.Error())
	}

	if row := database.Client.Model(&cart).Preload("Items").Where("user_id = ?", user.ID).First(&cart).RowsAffected; row == 0 {
		cart.UserID = user.ID
		if err := database.Client.Model(&cart).Create(&cart).Error; err != nil {
			return InternalServerData(c, err.Error())
		}
	}

	if count := database.Client.Model(&user).Where("id = ?", deviceID).Association("Devices").Count(); count != 0 {
		return BadRequest(c)
	}

	if count := database.Client.Model(&cart).Where("id = ?", deviceID).Association("Items").Count(); count != 0 {
		return BadRequest(c)
	}

	if row := database.Client.Model(&device).Where("id = ?", deviceID).First(&device).RowsAffected; row == 0 {
		return NotFound(c)
	}

	if err := database.Client.Model(&cart).Association("Items").Append(&device); err != nil {
		return InternalServerData(c, err.Error())
	}

	return Success(c)
}

func RemoveDeviceFromCart(c *fiber.Ctx) error {
	var cart model.Cart
	var user model.User
	var device model.Device
	deviceID := c.Params("id", "")
	var token string

	token = c.Cookies("token", "")

	if token == "" {
		token = strings.Split(c.GetReqHeaders()["Authorization"], " ")[1]
	}

	if err := helper.GetUserFromToken(token, config.SecretKeyApp, &user); err != nil {
		return BadRequestData(c, err.Error())
	}

	if row := database.Client.Model(&cart).Preload("Items").Where("user_id = ?", user.ID).First(&cart).RowsAffected; row == 0 {
		return BadRequest(c)
	}

	if count := database.Client.Model(&cart).Where("id = ?", deviceID).Association("Items").Count(); count == 0 {
		return NotFound(c)
	}

	if row := database.Client.Model(&device).Where("id = ?", deviceID).First(&device).RowsAffected; row == 0 {
		return NotFound(c)
	}

	if err := database.Client.Model(&cart).Association("Items").Delete(&device); err != nil {
		return InternalServerData(c, err.Error())
	}

	return Success(c)
}

func CartToOrder(c *fiber.Ctx) error {
	var cart model.Cart
	var order model.Order
	var user model.User

	var token string

	token = c.Cookies("token", "")

	if token == "" {
		token = strings.Split(c.GetReqHeaders()["Authorization"], " ")[1]
	}
	if err := helper.GetUserFromToken(token, config.SecretKeyApp, &user); err != nil {
		return BadRequestData(c, err.Error())
	}

	if row := database.Client.Model(&cart).Preload("Items").Where("user_id = ?", user.ID).First(&cart).RowsAffected; row == 0 {
		return BadRequest(c)
	}

	order.UserID = user.ID
	order.Status = "waiting"

	if err := database.Client.Model(&order).Create(&order).Error; err != nil {
		return InternalServerData(c, err.Error())
	}

	if err := database.Client.Model(&order).Association("Items").Append(&cart.Items); err != nil {
		return InternalServerData(c, err.Error())
	}

	if err := database.Client.Model(&cart).Association("Items").Clear(); err != nil {
		return InternalServerData(c, err.Error())
	}

	return Success(c)
}

func GetCartListFromUser(c *fiber.Ctx) error {
	var cart model.Cart
	var count int64
	var user model.User

	var token string

	token = c.Cookies("token", "")

	if token == "" {
		token = strings.Split(c.GetReqHeaders()["Authorization"], " ")[1]
	}
	limit := c.QueryInt("limit", 10)
	page := c.QueryInt("page", 1)
	search := helper.SearchString(c.Query("search", ""))
	offset := (page * limit) - limit

	if err := helper.GetUserFromToken(token, config.SecretKeyApp, &user); err != nil {
		return BadRequestData(c, err.Error())
	}

	if err := database.Client.Model(&model.Cart{}).
		Limit(limit).Offset(offset).Preload("Items.Address").
		Where("user_id = ?", user.ID).
		Where("id LIKE ?", search).
		Find(&cart).Error; err != nil {
		return InternalServerData(c, err.Error())
	}

	if err := database.Client.Model(&model.Cart{}).Where("user_id = ?", user.ID).Count(&count).Error; err != nil {
		return InternalServerData(c, err.Error())
	}

	total := math.Ceil(float64(count) / float64(limit))

	return SuccessPage(c, cart.Items, int64(total), int64(page))
}

func GetOrderList(c *fiber.Ctx) error {
	var list []model.Order
	var count int64
	limit := c.QueryInt("limit", 10)
	page := c.QueryInt("page", 1)
	search := helper.SearchString(c.Query("search", ""))
	offset := (page * limit) - limit

	if err := database.Client.Model(&model.Order{}).Preload("Items").Preload("User").
		Limit(limit).Offset(offset).Where("id LIKE ?", search).
		Or("status LIKE ?", search).
		Find(&list).Error; err != nil {
		return InternalServerData(c, err.Error())
	}

	if err := database.Client.Model(&model.Order{}).Count(&count).Error; err != nil {
		return InternalServerData(c, err.Error())
	}

	total := math.Ceil(float64(count) / float64(limit))

	return SuccessPage(c, list, int64(total), int64(page))
}

func FindOrder(c *fiber.Ctx) error {
	var order model.Order
	var count int64

	orderID := c.Params("id", "")
	limit := c.QueryInt("limit", 10)
	page := c.QueryInt("page", 0)
	search := helper.SearchString(c.Query("search", ""))
	offset := (page * limit) - limit

	if err := database.Client.Model(&order).
		Preload("Items", func(db *gorm.DB) *gorm.DB {
			return db.Preload("Address").Where(db.Where("id LIKE ?", search).
				Or("model LIKE ?", search).Limit(limit).Offset(offset))
		}).
		Preload("User").
		Where("id = ?", orderID).
		Find(&order).Error; err != nil {
		return InternalServerData(c, err.Error())
	}

	if err := database.Client.Model(&model.Order{}).Count(&count).Error; err != nil {
		return InternalServerData(c, err.Error())
	}

	total := math.Ceil(float64(count) / float64(limit))
	return SuccessPage(c, order, int64(total), int64(page))
}

func GetOrderListFromUSer(c *fiber.Ctx) error {
	var list []model.Order
	var count int64
	var user model.User

	var token string

	token = c.Cookies("token", "")

	if token == "" {
		token = strings.Split(c.GetReqHeaders()["Authorization"], " ")[1]
	}
	limit := c.QueryInt("limit", 10)
	page := c.QueryInt("page", 1)
	search := helper.SearchString(c.Query("search", ""))
	offset := (page * limit) - limit

	if err := helper.GetUserFromToken(token, config.SecretKeyApp, &user); err != nil {
		return BadRequestData(c, err.Error())
	}

	if err := database.Client.Model(&model.Order{}).
		Limit(limit).Offset(offset).Preload("Items").Preload("User").
		Where("user_id = ?", user.ID).
		Where(database.Client.Where("id LIKE ?", search).
			Or("status LIKE ?", search)).
		Find(&list).Error; err != nil {
		return InternalServerData(c, err.Error())
	}

	if err := database.Client.Model(&model.Order{}).Where("user_id = ?", user.ID).Count(&count).Error; err != nil {
		return InternalServerData(c, err.Error())
	}

	total := math.Ceil(float64(count) / float64(limit))

	return SuccessPage(c, list, int64(total), int64(page))
}

func OrderStatusChange(c *fiber.Ctx) error {
	var user model.User
	var order model.Order
	var form form.OrderStatus
	order.ID = c.Params("id", "")

	if err := c.BodyParser(&form); err != nil {
		return BadRequestData(c, err.Error())
	}

	if row := database.Client.Model(&order).Preload("Items").First(&order).RowsAffected; row == 0 {
		return NotFound(c)
	}

	if order.Status == "rijected" || order.Status == "accepted" {
		return BadRequest(c)
	}

	if row := database.Client.Model(&user).Preload("Devices").Where("id = ?", order.UserID).First(&user).RowsAffected; row == 0 {
		return NotFound(c)
	}

	if form.Status == "accepted" {
		if err := database.Client.Model(&user).Omit("Orders").Association("Devices").Append(&order.Items); err != nil {
			return InternalServerData(c, err.Error())
		}
	}

	if err := database.Client.Model(&order).Omit("Items").Update("status", form.Status).Error; err != nil {
		return InternalServerData(c, err.Error())
	}

	return Success(c)

}
