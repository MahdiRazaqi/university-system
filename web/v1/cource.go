package v1

import (
	"errors"
	"strconv"

	"github.com/MahdiRazaqi/university-system/course"
	"github.com/jinzhu/gorm"
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
		"message": "course created successfully",
		"course":  course.Mini(),
	})
}

/**
* @api {put} /api/v1/course/:id edit course
* @apiVersion 1.0.0
* @apiName editCourse
* @apiGroup Course
*
* @apiSuccess {String} message success message
* @apiSuccess {Object} course course model
*
* @apiError {String} error error message
 */
func editCourse(c echo.Context) error {
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

	id, _ := strconv.Atoi(c.Param("id"))

	course := &course.Course{
		Model:    gorm.Model{ID: uint(id)},
		Name:     formData.Name,
		CourseID: formData.CourseID,
		Unit:     formData.Unit,
	}

	if err := course.Save(); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	return c.JSON(200, echo.Map{
		"message": "course updated successfully",
		"course":  course.Mini(),
	})
}

/**
* @api {delete} /api/v1/course/:id delete course
* @apiVersion 1.0.0
* @apiName deleteCourse
* @apiGroup Course
*
* @apiSuccess {String} message success message
*
* @apiError {String} error error message
 */
func deleteCourse(c echo.Context) error {
	u, err := userBinding(c.Get("user"))
	if err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	if _, ok := u["TeacherID"]; !ok {
		return c.JSON(403, echo.Map{"error": errors.New("don't have permission").Error()})
	}

	id := c.Param("id")

	if err := course.Delete("id = ?", id); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	return c.JSON(200, echo.Map{
		"message": "course deleted successfully",
	})
}
