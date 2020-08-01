package v1

import (
	"errors"

	"github.com/MahdiRazaqi/university-system/course"
	"github.com/labstack/echo"
)

type courseForm struct {
	Name     string `json:"name" form:"name"`
	CourseID int    `json:"course_id" form:"course_id"`
	Unit     int    `json:"unit" form:"unit"`
}

/**
* @api {post} /api/v1/course add course
* @apiVersion 1.0.0
* @apiName addCourse
* @apiGroup Course
*
* @apiParam {String} name course name
* @apiParam {Number} course_id course id
* @apiParam {Number} unit course unit
*
* @apiSuccess {String} message success message
* @apiSuccess {Object} course course model
*
* @apiError {String} error error message
 */
func addCourse(c echo.Context) error {
	u, err := userBinding(c.Get("user"))
	if err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	if _, ok := u["TeacherID"]; !ok {
		return c.JSON(403, echo.Map{"error": errors.New("don't have permission").Error()})
	}

	formData := new(courseForm)
	if err := c.Bind(formData); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	course := &course.Course{
		Name:     formData.Name,
		CourseID: formData.CourseID,
		Unit:     formData.Unit,
	}
	if err := course.Save(); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	return c.JSON(200, echo.Map{
		"message": "post created successfully",
		"course":  course.Mini(),
	})
}
