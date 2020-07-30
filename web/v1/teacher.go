package v1

import (
	"github.com/MahdiRazaqi/university-system/teacher"
	"github.com/jeyem/passwd"
	"github.com/labstack/echo"
)

type teacherForm struct {
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
	TeacherID int    `json:"teacher_id" form:"teacher_id"`
	Password  string `json:"password" form:"password"`
}

/**
* @api {post} /api/v1/teacher/register Teacher register
* @apiVersion 1.0.0
* @apiName teacherRegister
* @apiGroup Teacher
*
* @apiParam {String} first_name first name
* @apiParam {String} last_name last name
* @apiParam {Number} teacher_id unique teacher id
* @apiParam {String} password password
*
* @apiSuccess {String} message success message
* @apiSuccess {Object} teacher teacher model
*
* @apiError {String} error error message
 */
func teacherRegister(c echo.Context) error {
	formData := new(teacherForm)
	if err := c.Bind(formData); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	t := &teacher.Teacher{
		FirstName: formData.FirstName,
		LastName:  formData.LastName,
		TeacherID: formData.TeacherID,
		Password:  passwd.Make(formData.Password),
	}
	if err := t.Save(); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	token, err := t.CreateToken()
	if err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	return c.JSON(200, echo.Map{
		"message": "registered successfully",
		"token":   token,
		"user":    t.Mini(),
	})
}

/**
* @api {post} /api/v1/teacher/login Teacher login
* @apiVersion 1.0.0
* @apiName teacherLogin
* @apiGroup Teacher
*
* @apiParam {Number} teacher_id unique teacher id
* @apiParam {String} password password
*
* @apiSuccess {String} message success message
* @apiSuccess {Object} teacher teacher model
*
* @apiError {String} error error message
 */
func teacherLogin(c echo.Context) error {
	formData := new(teacherForm)
	if err := c.Bind(formData); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	t, err := teacher.Auth(formData.TeacherID, formData.Password)
	if err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	token, err := t.CreateToken()
	if err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	return c.JSON(200, echo.Map{
		"message": "registered successfully",
		"token":   token,
		"user":    t.Mini(),
	})
}
