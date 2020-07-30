package v1

import (
	"github.com/labstack/echo"
)

// var signature = config.Config.SecretKey

// Register routes
func Register(e *echo.Echo) {
	v1 := e.Group("/api/v1")

	teacherGroup := v1.Group("/teacher")
	teacherGroup.POST("/register", teacherRegister)
	teacherGroup.POST("/login", teacherLogin)

	studentGroup := v1.Group("/student")
	studentGroup.POST("/register", studentRegister)
	studentGroup.POST("/login", studentLogin)

	// r := v1.Group("/")
	// r.Use(middleware.JWT([]byte(signature)), studentRequired)

	// postGroup := r.Group("post")
	// postGroup.POST("", addPost)
	// postGroup.GET("", listMyPosts)
	// postGroup.PUT("/:id", editPost)
	// postGroup.DELETE("/:id", removePost)

}

// func studentRequired(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		u := &user.User{}
// 		if err := u.LoadByToken(user.GetToken(c.Request())); err != nil {
// 			return c.JSON(400, echo.Map{"error": "loading user from token " + err.Error()})
// 		}
// 		c.Set("user", u)
// 		return next(c)
// 	}
// }
