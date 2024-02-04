package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
	"time"
	"unicore/pkg/constants"
)

var (
	DB          *gorm.DB
	MaxAttempts = 5
	TimeSleep   = 5
)

func Init() {
	var err error
	for attemps := 0; attemps < MaxAttempts; attemps++ {
		DB, err = gorm.Open(mysql.Open(constants.MySQLDSN),
			&gorm.Config{
				PrepareStmt:            true,
				SkipDefaultTransaction: true,
			},
		)
		if err == nil {
			break
		}
		time.Sleep(time.Duration(TimeSleep))
	}
	if err != nil {
		panic(err)
	}

	if err = DB.Use(gormopentracing.New()); err != nil {
		panic(err)
	}
}
