package router

import (
	"github.com/KuroNeko6666/sc-backend/handler"
	"github.com/KuroNeko6666/sc-backend/middleware"
	"github.com/gofiber/fiber/v2"
)

func admin(route fiber.Router) {
	base := route.Group("/admin")

	auth := base.Group("/auth")
	auth.Post("", handler.LoginAdmin)
	auth.Put("", handler.RegisterAdmin)
	auth.Get("", handler.ValidateTokenAdmin)
	auth.Delete("", handler.Logout)

	device := base.Group("/device", middleware.AuthAdmin)
	device.Post("/user", handler.AddUserToDevice)
	device.Delete("/user", handler.RemoveUserFromDevice)
	device.Get("/user/:id", handler.FindDeviceUser)

	device.Get("/data/:id", handler.FindDeviceDataFromAdmin)
	device.Post("/data", handler.CreateDeviceData)

	device.Post("", handler.CreateDevice)
	device.Get("", handler.GetDeviceFromAdmin)
	device.Put(":id", handler.UpdateDevice)
	device.Delete(":id", handler.DeleteDevice)

}
