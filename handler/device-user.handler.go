package handler

import (
	"log"
	"math"

	"github.com/KuroNeko6666/sc-backend/database"
	"github.com/KuroNeko6666/sc-backend/helper"
	"github.com/KuroNeko6666/sc-backend/interface/form"
	"github.com/KuroNeko6666/sc-backend/interface/model"
	"github.com/KuroNeko6666/sc-backend/interface/response"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
)

func AddUserToDevice(c *fiber.Ctx) error {
	var form form.UserDevice
	var user model.User
	var device model.Device

	if err := c.BodyParser(&form); err != nil {
		return BadRequestData(c, err.Error())
	}

	log.Println(form)

	if row := database.Client.Model(&user).Preload("Devices").Where("id = ?", form.UserID).First(&user).RowsAffected; row == 0 {
		return NotFound(c)
	}

	if row := database.Client.Model(&device).Where("id = ?", form.DeviceID).First(&device).RowsAffected; row == 0 {
		return NotFound(c)
	}

	if count := database.Client.Model(&device).Where("id = ?", user.ID).Association("Users").Count(); count != 0 {
		return BadRequest(c)
	}

	if err := database.Client.Model(&user).Association("Devices").Append(&device); err != nil {
		return InternalServerData(c, err.Error())
	}

	return Success(c)
}

func RemoveUserFromDevice(c *fiber.Ctx) error {
	var form form.UserDevice
	var user model.User
	var device model.Device

	form.DeviceID = c.Params("device")
	form.UserID = c.Params("user")

	if row := database.Client.Model(&user).Where("id = ?", form.UserID).First(&user).RowsAffected; row == 0 {
		return NotFound(c)
	}

	if row := database.Client.Model(&device).Where("id = ?", form.DeviceID).First(&device).RowsAffected; row == 0 {
		return NotFound(c)
	}

	if count := database.Client.Model(&device).Where("id = ?", user.ID).Association("Users").Count(); count == 0 {
		return BadRequest(c)
	}

	if err := database.Client.Model(&device).Association("Users").Delete(&user); err != nil {
		return InternalServerData(c, err.Error())
	}

	return Success(c)
}

func FindDeviceUser(c *fiber.Ctx) error {
	var device model.Device
	var users []model.User

	search := helper.SearchString(c.Query("search", ""))
	limit := c.QueryInt("limit", 10)
	page := c.QueryInt("page", 1)
	offset := (page * limit) - limit
	deviceID := c.Params("id", "")

	if row := database.Client.Model(&device).Where("id = ?", deviceID).First(&device).RowsAffected; row == 0 {
		return NotFound(c)
	}

	if err := database.Client.Model(&device).Limit(limit).Offset(offset).
		Where("name LIKE ?", search).Or("username LIKE ?", search).
		Or("email LIKE ?", search).Association("Users").Find(&users); err != nil {
		return InternalServerData(c, err.Error())
	}

	count := database.Client.Model(&device).Association("Users").Count()
	total := math.Ceil(float64(count) / float64(limit))

	return SuccessPage(c, users, int64(total), int64(page))

}

func GetDeviceUser(c *fiber.Ctx) error {
	var devices []model.Device
	var data []response.DevicesMarket
	var user model.User

	search := helper.SearchString(c.Query("search", ""))
	limit := c.QueryInt("limit", 10)
	page := c.QueryInt("page", 1)
	offset := (page * limit) - limit
	user.ID = c.Params("id", "")

	if err := database.Client.Model(&user).
		Joins("Address").
		Limit(limit).Offset(offset).
		Where("id LIKE ?", search).
		Or("model LIKE ?", search).
		Or("Address.address LIKE ?", search).
		Or("Address.city LIKE ?", search).
		Or("Address.province LIKE ?", search).
		Association("Devices").Find(&devices); err != nil {
		return InternalServerData(c, err.Error())
	}

	for _, device := range devices {
		var res response.DevicesMarket

		if err := copier.CopyWithOption(&res, &device, copier.Option{IgnoreEmpty: true}); err != nil {
			return InternalServerData(c, err.Error())
		}

		res.Subcribers = database.Client.Model(&device).Association("Users").Count()
		if count := database.Client.Model(&device).Where("id = ?", user.ID).Association("Users").Count(); count == 0 {
			res.Subcribe = false
		} else {
			res.Subcribe = true
		}
		data = append(data, res)
	}

	count := database.Client.Model(user).Association("Devices").Count()
	total := math.Ceil(float64(count) / float64(limit))

	return SuccessPage(c, data, int64(total), int64(page))
}
