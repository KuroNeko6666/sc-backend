package config

const (
	ServerHost = "0.0.0.0"
	ServerPort = ":8000"

	ResFailure = "failure"
	ResSuccess = "success"

	DatabaseHost = "188.166.245.130"
	DatabasePort = ":3306"
	DatabaseUser = "pub"
	DatabasePass = "Aks3sdb!@#"
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
