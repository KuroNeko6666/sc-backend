package config

const (
	ServerHost = "localhost"
	ServerPort = ":8000"

	ResFailure = "failure"
	ResSuccess = "success"

	DatabaseHost = "192.168.100.107"
	DatabasePort = ":3306"
	DatabaseUser = "mareca"
	DatabasePass = "mareca"
	DatabaseName = "speed_control_database"

	StorageHost = ""
	StoragePort = ""

	SecretKeyApp = "SP3EDC0ntR01"
)

func DatabaseDSN() string {
	option := "?charset=utf8mb4&parseTime=True&loc=Local"
	return DatabaseUser + ":" + DatabasePass + "@tcp(" + DatabaseHost + DatabasePort + ")/" + DatabaseName + option
}

func ServerAddress() string {
	return ServerHost + ServerPort
}
