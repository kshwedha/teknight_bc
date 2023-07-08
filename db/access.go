package db

import (
	"fmt"
	"teknight_bc/config"
)

func AccessString() (string, string) {
	driver := config.GetConfigVal("DB", "driver")
	username := config.GetConfigVal("DB", "username")
	passwd := config.GetConfigVal("DB", "password")
	protocol := "tcp"
	host := config.GetConfigVal("DB", "host")
	port := config.GetConfigVal("DB", "port")
	database := config.GetConfigVal("DB", "database")

	t_string := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", username, passwd, protocol, host, port, database)
	return driver, t_string
}
