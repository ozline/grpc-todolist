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
	c := new(config)
	if err := viper.Unmarshal(&c); err != nil {
		log.Fatal(err)
	}
	Snowflake = &c.Snowflake

	Server = &c.Server
	Server.Secret = []byte(viper.GetString("server.jwt-secret"))

	Etcd = &c.Etcd

	Mysql = &c.Mysql

	Service = GetService(srv)
}

func GetService(srvname string) *service {
	return &service{
		Name: viper.GetString("services." + srvname + ".name"),
		Addr: viper.GetString("services." + srvname + ".addr"),
	}
}
