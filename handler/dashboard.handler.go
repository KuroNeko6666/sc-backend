package handler

import (
	"github.com/KuroNeko6666/sc-backend/config"
	"github.com/KuroNeko6666/sc-backend/database"
	"github.com/KuroNeko6666/sc-backend/helper"
	"github.com/KuroNeko6666/sc-backend/interface/model"
	"github.com/KuroNeko6666/sc-backend/interface/response"
	"github.com/gofiber/fiber/v2"
)

func DahsboardTotal(c *fiber.Ctx) error {
	var user model.User

	token := c.Cookies("token", "")
	if err := helper.GetUserFromToken(token, config.SecretKeyApp, &user); err != nil {
		return InternalServerData(c, err.Error())
	}

	count := database.Client.Model(&user).Association("Devices").Count()

	return SuccessData(c, count)

}

func ChartDeviceDataCreated(c *fiber.Ctx) error {
	var user model.User
	var response response.Chart
	deviceID := c.Params("id", "")
	dateType := c.Query("date_type", "day")
	token := c.Cookies("token", "")

	if err := helper.DateTypeValidate(dateType); err != nil {
		return BadRequestData(c, err.Error())
	}

	dateList := helper.TimeHandler(dateType)
	dateRanges := helper.RangeTimeHandler(dateList, dateType)
	dateLabels := helper.LabelHandler(dateList, dateType)

	if err := helper.GetUserFromToken(token, config.SecretKeyApp, &user); err != nil {
		return InternalServerData(c, err.Error())
	}

	if count := database.Client.Model(&user).Where("id = ?", deviceID).Association("Devices").Count(); count == 0 {
		return NotFound(c)
	}

	for index, date := range dateList {
		var count int64
		database.Client.Model(&model.DeviceData{}).Where("device_id", deviceID).Where("created_at BETWEEN ? AND ?", date, dateRanges[index]).Count(&count)
		response.Data = append(response.Data, count)
		response.Labels = append(response.Labels, dateLabels[index])
	}

	return SuccessData(c, response)
}

func ChartUserCreated(c *fiber.Ctx) error {
	var response response.Chart
	dateType := c.Query("date_type", "day")

	if err := helper.DateTypeValidate(dateType); err != nil {
		return BadRequestData(c, err.Error())
	}

	dateList := helper.TimeHandler(dateType)
	dateRanges := helper.RangeTimeHandler(dateList, dateType)
	dateLabels := helper.LabelHandler(dateList, dateType)

	for index, date := range dateList {
		var count int64
		database.Client.Model(&model.User{}).Where("created_at BETWEEN ? AND ?", date, dateRanges[index]).Count(&count)
		response.Data = append(response.Data, count)
		response.Labels = append(response.Labels, dateLabels[index])
	}

	return SuccessData(c, response)
}

func ChartAdminCreated(c *fiber.Ctx) error {
	var response response.Chart
	dateType := c.Query("date_type", "day")

	if err := helper.DateTypeValidate(dateType); err != nil {
		return BadRequestData(c, err.Error())
	}

	dateList := helper.TimeHandler(dateType)
	dateRanges := helper.RangeTimeHandler(dateList, dateType)
	dateLabels := helper.LabelHandler(dateList, dateType)

	for index, date := range dateList {
		var count int64
		database.Client.Model(&model.Admin{}).Where("created_at BETWEEN ? AND ?", date, dateRanges[index]).Count(&count)
		response.Data = append(response.Data, count)
		response.Labels = append(response.Labels, dateLabels[index])
	}

	return SuccessData(c, response)
}

func ChartDeviceSpeed(c *fiber.Ctx) error {
	var user model.User
	var response response.Chart
	deviceID := c.Params("id", "")
	dateType := c.Query("date_type", "day")
	token := c.Cookies("token", "")

	if err := helper.DateTypeValidate(dateType); err != nil {
		return BadRequestData(c, err.Error())
	}

	dateList := helper.SingleTimeHandler(dateType)
	speedRange := [3][2]int64{{0, 20}, {21, 40}, {41, 10000}}
	dateLabels := [3]string{"0km/d - 20km/d", "21km/d - 40km/d", "41km/d - 60km/d"}

	if err := helper.GetUserFromToken(token, config.SecretKeyApp, &user); err != nil {
		return InternalServerData(c, err.Error())
	}

	if count := database.Client.Model(&user).Where("id = ?", deviceID).Association("Devices").Count(); count == 0 {
		return NotFound(c)
	}

	for index, speed := range speedRange {
		var count int64
		database.Client.Model(&model.DeviceData{}).Where("device_id", deviceID).Where("created_at BETWEEN ? AND ?", dateList[0], dateList[1]).Where("speed BETWEEN ? AND ?", speed[0], speed[1]).Count(&count)
		response.Data = append(response.Data, count)
		response.Labels = append(response.Labels, dateLabels[index])
	}

	return SuccessData(c, response)
}
