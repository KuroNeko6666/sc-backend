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
	device.Get("/data/:id", handler.FindDeviceDataFromAdmin)
	device.Post("/data", handler.CreateDeviceData)

	device.Post("", handler.CreateDevice)
	device.Get("", handler.GetDeviceFromAdmin)
	device.Put(":id", handler.UpdateDevice)
	device.Delete(":id", handler.DeleteDevice)
	device.Get(":id", handler.FindDevice)

	min := base.Group("/admin")
	min.Post("", handler.CreateAdmin)
	min.Get("", handler.GetAdmin)
	min.Get(":id", handler.FindAdmin)
	min.Put(":id", handler.UpdateAdmin)
	min.Delete(":id", handler.DeleteAdmin)

	user := base.Group("/user")

	user.Post("/device", handler.AddUserToDevice)
	user.Get("/device/user/:id", handler.GetDeviceNotHave)
	user.Delete("/device/:device/:user", handler.RemoveUserFromDevice)
	user.Get("/device/:id", handler.GetDeviceUser)

	user.Post("", handler.CreateUser)
	user.Get("", handler.GetUser)
	user.Get(":id", handler.FindUser)
	user.Put(":id", handler.UpdateUser)
	user.Delete(":id", handler.DeleteUser)

	dashboard := base.Group("/dashboard")
	dashboard.Get("/user", handler.ChartUserCreated)
	dashboard.Get("/admin", handler.ChartAdminCreated)

	order := base.Group("/order", middleware.AuthAdmin)
	order.Get("", handler.GetAllOrderAdmin)
	order.Put("", handler.UpdateOrderStatusAdmin)
}
