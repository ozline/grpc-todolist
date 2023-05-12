package config

import (
	"log"

	"github.com/spf13/viper"
)

func Init(path string) {
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
}
