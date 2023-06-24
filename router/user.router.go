package router

import (
	"github.com/KuroNeko6666/sc-backend/handler"
	"github.com/KuroNeko6666/sc-backend/middleware"
	"github.com/gofiber/fiber/v2"
)

func user(route fiber.Router) {
	base := route.Group("/user")

	auth := base.Group("/auth")
	auth.Post("", handler.LoginUser)
	auth.Put("", handler.RegisterUser)
	auth.Get("", handler.ValidateTokenUser)
	auth.Delete("", handler.Logout)

	device := base.Group("/device", middleware.AuthUser)
	device.Get("/market", handler.GetDeviceForMarket)

	device.Get("", handler.GetDeviceFromUser)
	device.Get(":id", handler.FindDeviceDataFromUser)

	cart := base.Group("/cart", middleware.AuthUser)
	cart.Get("", handler.UserGetItemCart)
	cart.Post("/add-item", handler.UserAddItemToCart)
	cart.Delete("/remove-item", handler.RemoveUserFromDevice)

	order := base.Group("/order", middleware.AuthUser)
	order.Get("", handler.UserGetItemOrder)
	order.Post("/create-transaction", handler.UserCreateOrder)
	// order.Post("/update-transaction", handler.UserCreateOrder)
}
