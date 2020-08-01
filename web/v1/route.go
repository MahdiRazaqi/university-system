package v1

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Register routes
func Register(e *echo.Echo) {
	v1 := e.Group("/api/v1")

	teacherGroup := v1.Group("/teacher")
	teacherGroup.POST("/register", teacherRegister)
	teacherGroup.POST("/login", teacherLogin)

	studentGroup := v1.Group("/student")
	studentGroup.POST("/register", studentRegister)
	studentGroup.POST("/login", studentLogin)

	r := v1.Group("/")
	r.Use(middleware.JWT([]byte(signature)), userRequired)

	courceGroup := r.Group("cource")
	courceGroup.POST("", addCourse)
	// postGroup.GET("", listMyPosts)
	// postGroup.PUT("/:id", editPost)
	// postGroup.DELETE("/:id", removePost)

}

func userRequired(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		u, _ := LoadByToken(GetToken(c.Request()))
		c.Set("user", u)
		return next(c)
	}
}
