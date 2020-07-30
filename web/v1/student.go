package v1

import (
	"github.com/MahdiRazaqi/university-system/student"
	"github.com/jeyem/passwd"
	"github.com/labstack/echo"
)

type studentForm struct {
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
	StudentID int    `json:"student_id" form:"student_id"`
	Password  string `json:"password" form:"password"`
}

/**
* @api {post} /api/v1/student/register Student register
* @apiVersion 1.0.0
* @apiName studentRegister
* @apiGroup Student
*
* @apiParam {String} first_name first name
* @apiParam {String} last_name last name
* @apiParam {Number} student_id unique student id
* @apiParam {String} password password
*
* @apiSuccess {String} message success message
* @apiSuccess {Object} student student model
*
* @apiError {String} error error message
 */
func studentRegister(c echo.Context) error {
	formData := new(studentForm)
	if err := c.Bind(formData); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	s := &student.Student{
		FirstName: formData.FirstName,
		LastName:  formData.LastName,
		StudentID: formData.StudentID,
		Password:  passwd.Make(formData.Password),
	}
	if err := s.Save(); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	token, err := s.CreateToken()
	if err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	return c.JSON(200, echo.Map{
		"message": "registered successfully",
		"token":   token,
		"user":    s.Mini(),
	})
}

/**
* @api {post} /api/v1/student/login Student login
* @apiVersion 1.0.0
* @apiName studentLogin
* @apiGroup Student
*
* @apiParam {Number} student_id unique student id
* @apiParam {String} password password
*
* @apiSuccess {String} message success message
* @apiSuccess {Object} student student model
*
* @apiError {String} error error message
 */
func studentLogin(c echo.Context) error {
	formData := new(studentForm)
	if err := c.Bind(formData); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	s, err := student.Auth(formData.StudentID, formData.Password)
	if err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	token, err := s.CreateToken()
	if err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	return c.JSON(200, echo.Map{
		"message": "registered successfully",
		"token":   token,
		"user":    s.Mini(),
	})
}
