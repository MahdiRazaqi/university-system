package v1

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/MahdiRazaqi/university-system/course"
	"github.com/MahdiRazaqi/university-system/database"
	"github.com/hb-go/json"
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

	crs := &course.Course{
		Name:     formData.Name,
		CourseID: formData.CourseID,
		Unit:     formData.Unit,
	}
	if err := crs.Save(); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	return c.JSON(200, echo.Map{
		"message": "course created successfully",
		"course":  crs.Mini(),
	})
}

/**
* @api {put} /api/v1/course/:id edit course
* @apiVersion 1.0.0
* @apiName editCourse
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

	crs := &course.Course{
		Model:    gorm.Model{ID: uint(id)},
		Name:     formData.Name,
		CourseID: formData.CourseID,
		Unit:     formData.Unit,
	}

	if err := crs.Save(); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	return c.JSON(200, echo.Map{
		"message": "course updated successfully",
		"course":  crs.Mini(),
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

	id, _ := strconv.Atoi(c.Param("id"))

	if err := course.Delete("id = ?", id); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	return c.JSON(200, echo.Map{
		"message": "course deleted successfully",
	})
}

/**
* @api {get} /api/v1/course/:id get one course
* @apiVersion 1.0.0
* @apiName getCourse
* @apiGroup Course
*
* @apiSuccess {Object} course course model
*
* @apiError {String} error error message
 */
func getCourse(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	crs, err := course.FindOne("id = ?", id)
	if err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	return c.JSON(200, echo.Map{
		"course": crs.Mini(),
	})
}

/**
* @api {get} /api/v1/course get courses list
* @apiVersion 1.0.0
* @apiName getCoursesList
* @apiGroup Course
*
* @apiParam {Number} page list page
* @apiParam {Number} limit list limit
*
* @apiSuccess {Object} course course model
*
* @apiError {String} error error message
 */
func getCoursesList(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	redisID := fmt.Sprintf("course-%v-%v", page, limit)
	var result []map[string]interface{}

	cacheData, err := database.Redis.Get(redisID).Result()
	if err == nil {
		json.Unmarshal([]byte(cacheData), &result)

		return c.JSON(200, echo.Map{
			"courses": result,
		})
	}

	courses, err := course.Find(limit, page, "")
	if err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	for _, crs := range *courses {
		result = append(result, crs.Mini())
	}

	jsonData, _ := json.Marshal(result)
	database.Redis.Set(redisID, jsonData, 1*time.Hour)

	return c.JSON(200, echo.Map{
		"courses": result,
	})
}
