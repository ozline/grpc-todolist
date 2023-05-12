package utils

import (
	"strings"

	"github.com/spf13/viper"
)

func GetMysqlDSN() string {
	host := viper.GetString("mysql.host")
	port := viper.GetString("mysql.port")
	database := viper.GetString("mysql.database")
	username := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	charset := viper.GetString("mysql.charset")

	dsn := strings.Join([]string{username, ":", password, "@tcp(", host, ":", port, ")/", database, "?charset=" + charset + "&parseTime=true"}, "")

	return dsn
}
