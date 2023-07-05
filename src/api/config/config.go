package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

const (
	minutesConnMaxLifetime  = 2
	maxIdleConn             = 150
	maxOpenConn             = 300
	userName				= "root"
	passWord				= ""
	host					= "localhost:3306"
	database				= "ida_test"
)

func CreateClient() (client *gorm.DB) {
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Error),
	}

	client, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true",
		userName, passWord, host, database)), gormConfig)

	if err != nil {
		panic(any(err.Error()))
	}

	sqlDB, err := client.DB()
	if err != nil {
		panic(any(err.Error()))
	}

	sqlDB.SetConnMaxLifetime(time.Minute * minutesConnMaxLifetime)
	sqlDB.SetMaxIdleConns(maxIdleConn)
	sqlDB.SetMaxOpenConns(maxOpenConn)

	return client
}
