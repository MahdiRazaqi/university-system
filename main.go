package main

import (
	"github.com/MahdiRazaqi/university-system/config"
	"github.com/MahdiRazaqi/university-system/database"
	"github.com/MahdiRazaqi/university-system/web"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := config.Load(); err != nil {
		logrus.Error("loading config file ", err)
	}

	database.Connect()
	web.Start()
}
