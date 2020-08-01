package v1

import (
	"errors"
	"time"

	"github.com/MahdiRazaqi/university-system/presentation"
	"github.com/MahdiRazaqi/university-system/teacher"
	"github.com/labstack/echo"
)

type presentationForm struct {
	CID      uint   `json:"cid" form:"cid"`
	Time     string `json:"time" form:"time"`
	Capacity int    `json:"capacity" form:"capacity"`
}

/**
* @api {post} /api/v1/presentation add presentation
* @apiVersion 1.0.0
* @apiName addPresentation
* @apiGroup presentation
*
* @apiParam {Number} cid course id
* @apiParam {String} time example 12:30
* @apiParam {Number} capacity presentation capacity
*
* @apiSuccess {String} message success message
*
* @apiError {String} error error message
 */
func addPresentation(c echo.Context) error {
	u, err := userBinding(c.Get("user"))
	if err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	if _, ok := u["TeacherID"]; !ok {
		return c.JSON(403, echo.Map{"error": errors.New("don't have permission").Error()})
	}

	formData := new(presentationForm)
	if err := c.Bind(formData); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	time, err := time.Parse("15:04", formData.Time)
	if err := c.Bind(formData); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	t, err := teacher.FindOne("teacher_id = ?", u["TeacherID"])
	if err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	p := &presentation.Presentation{
		TID:      t.ID,
		CID:      formData.CID,
		Time:     time,
		Capacity: formData.Capacity,
	}

	if err := p.Save(); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	return c.JSON(200, echo.Map{
		"message": "presentation created successfully",
	})
}
