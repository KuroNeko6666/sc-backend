package handler

import (
	"log"
	"time"

	"github.com/KuroNeko6666/sc-backend/config"
	"github.com/KuroNeko6666/sc-backend/database"
	"github.com/KuroNeko6666/sc-backend/helper"
	"github.com/KuroNeko6666/sc-backend/interface/form"
	"github.com/KuroNeko6666/sc-backend/interface/model"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
)

func LoginUser(c *fiber.Ctx) error {
	var form form.Login
	var user model.User
	var token string

	if err := c.BodyParser(&form); err != nil {
		return BadRequestData(c, err.Error())
	}

	if row := database.Client.Model(&user).
		Where("username = ?", form.Email).
		Or("email = ?", form.Email).Find(&user).RowsAffected; row == 0 {
		return UnAuthorized(c)
	}

	if err := helper.CompareHash(user.Password, form.Password); err != nil {
		return UnAuthorized(c)
	}

	if err := helper.GetTokenFromUser(user, config.SecretKeyApp, &token); err != nil {
		return UnAuthorized(c)
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.Cookie(cookie)

	return SuccessData(c, user)
}

func LoginAdmin(c *fiber.Ctx) error {
	var form form.Login
	var admin model.Admin
	var token string

	if err := c.BodyParser(&form); err != nil {
		return BadRequestData(c, err.Error())
	}

	if row := database.Client.Model(&admin).
		Where("username = ?", form.Email).
		Or("email = ?", form.Email).Find(&admin).RowsAffected; row == 0 {
		return UnAuthorized(c)
	}

	if err := helper.CompareHash(admin.Password, form.Password); err != nil {
		return UnAuthorized(c)
	}

	if err := helper.GetTokenFromAdmin(admin, config.SecretKeyApp, &token); err != nil {
		return UnAuthorized(c)
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.Cookie(cookie)

	return SuccessData(c, admin)
}

func RegisterUser(c *fiber.Ctx) error {
	var form form.User
	var data model.User

	if err := c.BodyParser(&form); err != nil {
		return BadRequestData(c, err.Error())
	}

	log.Println(&form)

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

func RegisterAdmin(c *fiber.Ctx) error {
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

func ActivateUser(c *fiber.Ctx) error {
	var user model.User
	userID := c.Params("id", "")

	if row := database.Client.Model(&user).Where("id = ?", userID).Find(&user).RowsAffected; row == 0 {
		return NotFound(c)
	}

	user.Status = !user.Status

	if err := database.Client.Model(&user).Updates(&user).Error; err != nil {
		return InternalServerData(c, err.Error())
	}

	return Success(c)
}

func Logout(c *fiber.Ctx) error {
	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now().Add(-1 * time.Hour)
	c.Cookie(cookie)

	return Success(c)
}

func ValidateTokenUser(c *fiber.Ctx) error {
	var user model.User

	token := c.Cookies("token", "")

	if err := helper.GetUserFromToken(token, config.SecretKeyApp, &user); err != nil {
		return Validate(c, false)
	}

	if row := database.Client.Model(&user).Find(&user).RowsAffected; row == 0 {
		return Validate(c, false)
	}

	return Validate(c, true)
}

func ValidateTokenAdmin(c *fiber.Ctx) error {
	var admin model.Admin

	token := c.Cookies("token", "")

	if err := helper.GetAdminFromToken(token, config.SecretKeyApp, &admin); err != nil {
		return Validate(c, false)
	}

	if row := database.Client.Model(&admin).Find(&admin).RowsAffected; row == 0 {
		return Validate(c, false)
	}

	return Validate(c, true)
}
