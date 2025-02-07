package database

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	connection *gorm.DB
)

func connect() {
	dsn := getDsn()
	c, err := getConnection(dsn)
	if err != nil {
		c.Error = err
		logrus.Errorf("error connecting to database %s", err.Error())

	}
	connection = c
}

func GetConnector() *gorm.DB {
	if connection == nil {
		connect()
	}
	return connection
}
