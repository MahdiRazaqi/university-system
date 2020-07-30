package v1

import "github.com/labstack/echo"

// Register routes
func Register(e *echo.Echo) {
	v1 := e.Group("/api/v1")

	teacherGroup := v1.Group("/teacher")
	teacherGroup.POST("/register", teacherRegister)
	teacherGroup.POST("/login", teacherLogin)

	studentGroup := v1.Group("/student")
	studentGroup.POST("/register", studentRegister)
	studentGroup.POST("/login", studentLogin)
}
