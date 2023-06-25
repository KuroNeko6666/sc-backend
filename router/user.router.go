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
	device.Get(":id", handler.FindDevice)

	dashboard := base.Group("/dashboard", middleware.AuthUser)
	dashboard.Get("/device-data/:id", handler.ChartDeviceDataCreated)
	dashboard.Get("/device-speed/:id", handler.ChartDeviceSpeed)
	dashboard.Get("/device-total", handler.DahsboardTotal)

	order := base.Group("/order")
	order.Get("", handler.GetOrderListFromUSer)
	order.Get("/:id", handler.FindOrder)
	order.Post("", handler.CartToOrder)

	cart := base.Group("/cart")
	cart.Get("", handler.GetCartListFromUser)
	cart.Put("/:id", handler.AddDeviceToCart)
	cart.Delete("/:id", handler.RemoveDeviceFromCart)

}
