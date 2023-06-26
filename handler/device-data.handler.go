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
	"github.com/jinzhu/copier"
)

func CreateDeviceData(c *fiber.Ctx) error {
	var form form.DeviceData
	var deviceData model.DeviceData

	if err := c.BodyParser(&form); err != nil {
		return BadRequestData(c, err.Error())
	}

	if err := copier.CopyWithOption(&deviceData, &form, copier.Option{IgnoreEmpty: true}); err != nil {
		return InternalServerData(c, err.Error())
	}

	if err := database.Client.Model(&deviceData).Create(&deviceData).Error; err != nil {
		return InternalServerData(c, err.Error())
	}

	return Success(c)
}

func FindDeviceDataFromAdmin(c *fiber.Ctx) error {
	var device model.Device
	var deviceData []model.DeviceData

	search := helper.SearchString(c.Query("search", ""))
	limit := c.QueryInt("limit", 10)
	page := c.QueryInt("page", 1)
	offset := (page * limit) - limit
	deviceID := c.Params("id", "")

	if row := database.Client.Model(&device).Where("id = ?", deviceID).First(&device).RowsAffected; row == 0 {
		return NotFound(c)
	}

	if err := database.Client.Model(&device).Limit(limit).Offset(offset).Where("speed LIKE ?", search).Or("distance LIKE ?", search).Association("Data").Find(&deviceData); err != nil {
		return InternalServerData(c, err.Error())
	}

	count := database.Client.Model(&device).Association("Data").Count()
	total := math.Ceil(float64(count) / float64(limit))

	return SuccessPage(c, deviceData, int64(total), int64(page))
}

func FindDeviceDataFromUser(c *fiber.Ctx) error {
	var device model.Device
	var deviceData []model.DeviceData
	var user model.User

	search := helper.SearchString(c.Query("search", ""))
	limit := c.QueryInt("limit", 10)
	page := c.QueryInt("page", 1)
	offset := (page * limit) - limit

	deviceID := c.Params("id", "")
	var token string

	token = c.Cookies("token", "")

	if token == "" {
		token = strings.Split(c.GetReqHeaders()["Authorization"], " ")[1]
	}

	if err := helper.GetUserFromToken(token, config.SecretKeyApp, &user); err != nil {
		return UnAuthorized(c)
	}

	if count := database.Client.Model(&user).Association("Devices").Count(); count == 0 {
		return NotFound(c)
	}

	if row := database.Client.Model(&device).Where("id = ?", deviceID).First(&device).RowsAffected; row == 0 {
		return NotFound(c)
	}

	if err := database.Client.Model(&device).Limit(limit).Offset(offset).Where("speed LIKE ?", search).Or("distance LIKE ?", search).Association("Data").Find(&deviceData); err != nil {
		return InternalServerData(c, err.Error())
	}

	count := database.Client.Model(&device).Association("Data").Count()
	total := math.Ceil(float64(count) / float64(limit))

	return SuccessPage(c, deviceData, int64(total), int64(page))
}
