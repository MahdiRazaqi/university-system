package database

import (
	"github.com/sirupsen/logrus"
)

// Connect to database
func Connect() {
	if err := connectToMysql(); err != nil {
		logrus.Error("connecting to mysql ", err)
	}

	if err := connectToRedis(); err != nil {
		logrus.Error("connecting to redis ", err)
	}
}
