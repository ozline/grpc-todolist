package db

import (
	"github.com/ozline/grpc-todolist/config"
	"github.com/ozline/grpc-todolist/pkg/utils"
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

	SF, err = utils.NewSnowflake(config.Snowflake.WorkerID, config.Snowflake.DatancenterID)

	if err != nil {
		panic(err)
	}
}
