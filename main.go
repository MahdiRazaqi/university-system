package main

import (
	"github.com/MahdiRazaqi/university-system/config"
	"github.com/MahdiRazaqi/university-system/database"
	"github.com/MahdiRazaqi/university-system/web"
)

func main() {
	config.Load()
	database.Connect()
	web.Start()
}
