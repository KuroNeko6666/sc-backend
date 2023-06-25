package handler

import (
	"math"

	"github.com/KuroNeko6666/sc-backend/config"
	"github.com/KuroNeko6666/sc-backend/database"
	"github.com/KuroNeko6666/sc-backend/helper"
	"github.com/KuroNeko6666/sc-backend/interface/form"
	"github.com/KuroNeko6666/sc-backend/interface/model"
	"github.com/KuroNeko6666/sc-backend/interface/response"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
)

func CreateDevice(c *fiber.Ctx) error {
	var form form.Device
	var device model.Device
	var address model.DeviceAddress

	if err := c.BodyParser(&form); err != nil {
		return BadRequestData(c, err.Error())
	}

	if err := copier.CopyWithOption(&device, &form, copier.Option{IgnoreEmpty: true}); err != nil {
		return InternalServerData(c, err.Error())
	}

	if err := copier.CopyWithOption(&address, &form, copier.Option{IgnoreEmpty: true}); err != nil {
		return InternalServerData(c, err.Error())
	}

	device.Address = address

	if err := database.Client.Model(&device).Create(&device).Error; err != nil {
		return InternalServerData(c, err.Error())
	}

	return SuccessData(c, device)
}

func UpdateDevice(c *fiber.Ctx) error {
	var form form.UpdateDevice
	var device model.Device

	deviceID := c.Params("id")

	if err := c.BodyParser(&form); err != nil {
		return BadRequestData(c, err.Error())
	}

	if row := database.Client.Model(&device).Preload("Address").Where("id = ?", deviceID).Find(&device).RowsAffected; row == 0 {
		return NotFound(c)
	}

	if err := copier.CopyWithOption(&device, &form, copier.Option{IgnoreEmpty: true}); err != nil {
		return InternalServerData(c, err.Error())
	}

	if err := database.Client.Model(&device).Omit("Address").Updates(&device).Error; err != nil {
		return InternalServerData(c, err.Error())
	}

	device.Address.Address = form.Address
	device.Address.City = form.City
	device.Address.Province = form.Province
	device.Address.Longitude = form.Longitude
	device.Address.Latitude = form.Latitude

	if err := database.Client.Model(&device.Address).Updates(&device.Address).Error; err != nil {
		return InternalServerData(c, err.Error())
	}

	return SuccessData(c, device)
}

func DeleteDevice(c *fiber.Ctx) error {
	var device model.Device

	deviceID := c.Params("id")

	if row := database.Client.Model(&device).
		Preload("Address").Preload("Users").
		Preload("Data").Where("id = ?", deviceID).
		Find(&device).RowsAffected; row == 0 {
		return NotFound(c)
	}

	if err := database.Client.Model(&device).Association("Users").Clear(); err != nil {
		return InternalServerData(c, err.Error())
	}

	if err := database.Client.Model(&device).Association("Data").Delete(); err != nil {
		return InternalServerData(c, err.Error())
	}

	if err := database.Client.Model(&device.Address).Where("address_id = ?", device.Address.AddressID).Delete(&device.Address).Error; err != nil {
		return InternalServerData(c, err.Error())
	}

	if err := database.Client.Model(&device).Delete(&device).Error; err != nil {
		return InternalServerData(c, err.Error())
	}

	return SuccessData(c, device)
}

func GetDeviceFromAdmin(c *fiber.Ctx) error {
	var devices []model.Device
	var data []response.Devices
	var count int64

	search := helper.SearchString(c.Query("search", ""))
	limit := c.QueryInt("limit", 10)
	page := c.QueryInt("page", 1)
	offset := (page * limit) - limit

	if err := database.Client.Model(&model.Device{}).
		Joins("Address").
		Limit(limit).Offset(offset).
		Where("id LIKE ?", search).
		Or("model LIKE ?", search).
		Or("Address.address LIKE ?", search).
		Or("Address.city LIKE ?", search).
		Or("Address.province LIKE ?", search).
		Find(&devices).Error; err != nil {
		return InternalServerData(c, err.Error())
	}

	if err := database.Client.Model(&model.Device{}).Count(&count).Error; err != nil {
		return InternalServerData(c, err.Error())
	}

	for _, device := range devices {
		var res response.Devices

		if err := copier.CopyWithOption(&res, &device, copier.Option{IgnoreEmpty: true}); err != nil {
			return InternalServerData(c, err.Error())
		}

		res.Subcribers = database.Client.Model(&device).Association("Users").Count()
		data = append(data, res)
	}

	total := math.Ceil(float64(count) / float64(limit))

	return SuccessPage(c, data, int64(total), int64(page))
}

func GetDeviceFromUser(c *fiber.Ctx) error {
	var devices []model.Device
	var data []response.DevicesMarket

	var user model.User

	token := c.Cookies("token", "")

	if err := helper.GetUserFromToken(token, config.SecretKeyApp, &user); err != nil {
		return UnAuthorized(c)
	}

	search := helper.SearchString(c.Query("search", ""))
	limit := c.QueryInt("limit", 10)
	page := c.QueryInt("page", 1)
	offset := (page * limit) - limit

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

func GetDeviceForMarket(c *fiber.Ctx) error {
	var devices []model.Device
	var data []response.DevicesMarket

	var user model.User

	token := c.Cookies("token", "")

	if err := helper.GetUserFromToken(token, config.SecretKeyApp, &user); err != nil {
		return UnAuthorized(c)
	}

	search := helper.SearchString(c.Query("search", ""))
	limit := c.QueryInt("limit", 10)
	page := c.QueryInt("page", 1)
	offset := (page * limit) - limit

	if err := database.Client.Model(&model.Device{}).
		Joins("Address").
		Limit(limit).Offset(offset).
		Where("id LIKE ?", search).
		Or("model LIKE ?", search).
		Or("Address.address LIKE ?", search).
		Or("Address.city LIKE ?", search).
		Or("Address.province LIKE ?", search).Find(&devices).Error; err != nil {
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

func GetDeviceNotHave(c *fiber.Ctx) error {
	var devices []model.Device
	var data []response.DevicesMarket

	var user model.User
	user.ID = c.Params("id", "")

	search := helper.SearchString(c.Query("search", ""))
	limit := c.QueryInt("limit", 10)
	page := c.QueryInt("page", 1)
	offset := (page * limit) - limit

	if err := database.Client.Model(&model.Device{}).
		Joins("Address").
		Limit(limit).Offset(offset).
		Where("id LIKE ?", search).
		Or("model LIKE ?", search).
		Or("Address.address LIKE ?", search).
		Or("Address.city LIKE ?", search).
		Or("Address.province LIKE ?", search).Find(&devices).Error; err != nil {
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

func FindDevice(c *fiber.Ctx) error {
	var devices model.Device
	deviceID := c.Params("id", "")

	if err := database.Client.Model(&model.Device{}).
		Preload("Address").
		Where("id = ?", deviceID).
		Find(&devices).Error; err != nil {
		return InternalServerData(c, err.Error())
	}

	return SuccessData(c, devices)
}
