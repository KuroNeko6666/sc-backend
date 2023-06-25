package handler

import (
	"math"

	"github.com/KuroNeko6666/sc-backend/database"
	"github.com/KuroNeko6666/sc-backend/helper"
	"github.com/KuroNeko6666/sc-backend/interface/form"
	"github.com/KuroNeko6666/sc-backend/interface/model"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
)

func GetAdmin(c *fiber.Ctx) error {
	var list []model.Admin
	var count int64

	limit := c.QueryInt("limit", 10)
	page := c.QueryInt("page", 1)
	search := helper.SearchString(c.Query("search", ""))
	offset := (page * limit) - limit

	if err := database.Client.Model(&model.Admin{}).
		Limit(limit).Offset(offset).Where("name LIKE ?", search).
		Or("username LIKE ?", search).Or("email LIKE ?", search).
		Find(&list).Error; err != nil {
		return InternalServerData(c, err.Error())
	}

	if err := database.Client.Model(&model.Admin{}).Count(&count).Error; err != nil {
		return InternalServerData(c, err.Error())
	}

	total := math.Ceil(float64(count) / float64(limit))

	return SuccessPage(c, list, int64(total), int64(page))
}

func FindAdmin(c *fiber.Ctx) error {
	var data model.Admin

	dataID := c.Params("id", "")

	if row := database.Client.Model(&model.Admin{}).
		Where("id = ?", dataID).
		First(&data).RowsAffected; row == 0 {
		return NotFound(c)
	}

	return SuccessData(c, data)
}

func CreateAdmin(c *fiber.Ctx) error {
	var form form.Admin
	var data model.Admin

	if err := c.BodyParser(&form); err != nil {
		return BadRequestData(c, err.Error())
	}

	if err := helper.GenerateHash(&form.Password); err != nil {
		return InternalServerData(c, err.Error())
	}

	if err := copier.CopyWithOption(&data, &form, copier.Option{IgnoreEmpty: true}); err != nil {
		return InternalServerData(c, err.Error())
	}

	if err := database.Client.Model(&data).Create(&data).Error; err != nil {
		return InternalServerData(c, err.Error())
	}

	return Success(c)
}

func UpdateAdmin(c *fiber.Ctx) error {
	var form form.AdminUpdate
	var data model.Admin
	dataID := c.Params("id")

	if row := database.Client.Model(&model.Admin{}).
		Where("id = ?", dataID).
		First(&data).RowsAffected; row == 0 {
		return NotFound(c)
	}

	if err := c.BodyParser(&form); err != nil {
		return BadRequestData(c, err.Error())
	}

	if err := copier.CopyWithOption(&data, &form, copier.Option{IgnoreEmpty: true}); err != nil {
		return InternalServerData(c, err.Error())
	}

	if err := database.Client.Model(&data).Updates(&data).Error; err != nil {
		return InternalServerData(c, err.Error())
	}

	return Success(c)
}

func DeleteAdmin(c *fiber.Ctx) error {
	var data model.Admin

	dataID := c.Params("id", "")

	if row := database.Client.Model(&model.Admin{}).
		Where("id = ?", dataID).
		First(&data).RowsAffected; row == 0 {
		return NotFound(c)
	}

	if err := database.Client.Model(&data).
		Delete(&data).Error; err != nil {
		return InternalServerData(c, err.Error())
	}

	return Success(c)
}
