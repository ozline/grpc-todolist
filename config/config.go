package config

import (
	"log"

	"github.com/spf13/viper"
)

var (
	Server    *server
	Mysql     *mysql
	Snowflake *snowflake
	Service   *service
	Etcd      *etcd
)

func Init(path string, service string) {
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("could not find config files")
		} else {
			log.Panicln("read config error")
		}
		log.Fatal(err)
	}

	configMapping(service) // 映射配置
}

func configMapping(srv string) {

	Server = &server{
		Secret:  []byte(viper.GetString("server.jwt-secret")),
		Version: viper.GetString("server.version"),
		Name:    viper.GetString("server.name"),
	}

	Snowflake = &snowflake{
		WorkerID:      viper.GetInt64("snowflake.worker-id"),
		DatancenterID: viper.GetInt64("snowflake.datancenter-id"),
	}

	Service = &service{
		Name: viper.GetString("services." + srv + ".name"),
		Addr: viper.GetString("services." + srv + ".addr"),
	}

	Mysql = &mysql{
		Addr:     viper.GetString("mysql.addr"),
		Database: viper.GetString("mysql.database"),
		Username: viper.GetString("mysql.username"),
		Password: viper.GetString("mysql.password"),
		Charset:  viper.GetString("mysql.charset"),
	}

	Etcd = &etcd{
		Addr: viper.GetString("etcd.addr"),
	}
}

func GetService(srvname string) *service {
	return &service{
		Name: viper.GetString("services." + srvname + ".name"),
		Addr: viper.GetString("services." + srvname + ".addr"),
	}
}
