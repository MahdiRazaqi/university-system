package v1

import (
	"errors"
	"strconv"

	"github.com/MahdiRazaqi/university-system/obtained"
	"github.com/MahdiRazaqi/university-system/presentation"
	"github.com/MahdiRazaqi/university-system/student"
	"github.com/labstack/echo"
)

/**
* @api {post} /api/v1/presentation/:id add to obtained
* @apiVersion 1.0.0
* @apiName addToObtained
* @apiGroup presentation
*
*
* @apiSuccess {String} message success message
*
* @apiError {String} error error message
 */
func addToObtained(c echo.Context) error {
	u, err := userBinding(c.Get("user"))
	if err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	if _, ok := u["StudentID"]; !ok {
		return c.JSON(403, echo.Map{"error": errors.New("don't have permission").Error()})
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	s, err := student.FindOne("student_id = ?", u["StudentID"])
	if err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	p, err := presentation.FindOne("id= ?", id)
	if err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	if p.Capacity <= 0 {
		return c.JSON(400, echo.Map{"error": errors.New("No capacity").Error()})
	}

	o := &obtained.Obtained{
		SID: s.ID,
		PID: uint(id),
	}

	if err := o.Save(); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	p.Capacity--

	if err := p.Save(); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	return c.JSON(200, echo.Map{
		"message": "presentation obtained successfully",
	})
}
