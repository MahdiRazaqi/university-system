package web

import (
	v1 "github.com/MahdiRazaqi/university-system/web/v1"
	"github.com/labstack/echo"
)

// Start web service
func Start() {
	e := echo.New()
	v1.Register(e)
	e.Start(":8080")
}
