package handler

import (
	"net/http"
	"time"

	"github.com/KuroNeko6666/sc-backend/config"
	"github.com/KuroNeko6666/sc-backend/interface/response"
	"github.com/gofiber/fiber/v2"
)

func NotFound(c *fiber.Ctx) error {
	status := http.StatusNotFound
	return c.Status(status).JSON(response.Base{
		Message: config.ResFailure,
		Data:    http.StatusText(status),
	})
}

func NotFoundData(c *fiber.Ctx, data string) error {
	status := http.StatusNotFound
	return c.Status(status).JSON(response.Base{
		Message: config.ResFailure,
		Data:    data,
	})
}

func BadRequest(c *fiber.Ctx) error {
	status := http.StatusBadRequest
	return c.Status(status).JSON(response.Base{
		Message: config.ResFailure,
		Data:    http.StatusText(status),
	})
}

func BadRequestData(c *fiber.Ctx, data string) error {
	status := http.StatusBadRequest
	return c.Status(status).JSON(response.Base{
		Message: config.ResFailure,
		Data:    data,
	})
}

func InternalServer(c *fiber.Ctx) error {
	status := http.StatusInternalServerError
	return c.Status(status).JSON(response.Base{
		Message: config.ResFailure,
		Data:    http.StatusText(status),
	})
}

func InternalServerData(c *fiber.Ctx, data string) error {
	status := http.StatusBadRequest
	return c.Status(status).JSON(response.Base{
		Message: config.ResFailure,
		Data:    data,
	})
}

func UnAuthorized(c *fiber.Ctx) error {
	status := http.StatusUnauthorized
	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now().Add(-1 * time.Hour)
	c.Cookie(cookie)
	return c.Status(status).JSON(response.Base{
		Message: config.ResFailure,
		Data:    http.StatusText(status),
	})
}

func UnAuthorizedData(c *fiber.Ctx, data string) error {
	status := http.StatusUnauthorized
	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now().Add(-1 * time.Hour)
	c.Cookie(cookie)
	return c.Status(status).JSON(response.Base{
		Message: config.ResFailure,
		Data:    data,
	})
}

func Success(c *fiber.Ctx) error {
	status := http.StatusOK
	return c.Status(status).JSON(response.Base{
		Message: config.ResSuccess,
		Data:    http.StatusText(status),
	})
}

func SuccessData(c *fiber.Ctx, data interface{}) error {
	status := http.StatusOK
	return c.Status(status).JSON(response.Base{
		Message: config.ResSuccess,
		Data:    data,
	})
}

func SuccessPage(c *fiber.Ctx, data interface{}, total int64, page int64) error {
	status := http.StatusOK
	return c.Status(status).JSON(response.Page{
		Message: config.ResSuccess,
		Data:    data,
		Total:   total,
		Page:    page,
	})
}

func Validate(c *fiber.Ctx, data bool) error {
	status := http.StatusOK
	return c.Status(status).JSON(data)
}
