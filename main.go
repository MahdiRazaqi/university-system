package main

import (
	"github.com/MahdiRazaqi/university-system/config"
	"github.com/MahdiRazaqi/university-system/web"
)

func main() {
	config.Load()
	web.Start()
}
