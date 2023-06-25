package database

import (
	"log"

	"github.com/KuroNeko6666/sc-backend/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Client *gorm.DB

func ConnectDB() {
	var err error
	Client, err = gorm.Open(mysql.Open(config.DatabaseDSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal(err)
	}

	// err = Client.AutoMigrate(
	// 	&model.User{},
	// 	&model.Admin{},
	// 	&model.Device{},
	// 	&model.DeviceData{},
	// 	&model.DeviceAddress{},
	// 	&model.Cart{},
	// 	&model.Order{},
	// )

	if err != nil {
		log.Fatal(err)
	}

}
