package db

import (
	"github.com/ozline/grpc-todolist/pkg/utils"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	gormopentracing "gorm.io/plugin/opentracing"
)

var DB *gorm.DB
var SF *utils.Snowflake

func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(utils.GetMysqlDSN()),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		})

	if err != nil {
		panic(err)
	}

	if err = DB.Use(gormopentracing.New()); err != nil {
		panic(err)
	}

	DB = DB.Table(viper.GetString("services.task.name"))

	SF, err = utils.NewSnowflake(viper.GetInt64("snowflake.worker-id"), viper.GetInt64("snowflake.datancenter-id"))

	if err != nil {
		panic(err)
	}
}
