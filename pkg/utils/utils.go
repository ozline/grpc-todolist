package utils

import (
	"strings"

	"github.com/ozline/grpc-todolist/config"
)

// WARNING: please ensure config load successfully
func GetMysqlDSN() string {
	dsn := strings.Join([]string{config.Mysql.Username, ":", config.Mysql.Password, "@tcp(", config.Mysql.Addr, ")/", config.Mysql.Database, "?charset=" + config.Mysql.Charset + "&parseTime=true"}, "")

	return dsn
}
