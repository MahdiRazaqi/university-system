package course

import (
	"errors"

	"github.com/MahdiRazaqi/university-system/database"
	"github.com/jinzhu/gorm"
)

// Course model
type Course struct {
	gorm.Model
	Name     string
	CourseID int
	Unit     int
}

func (c *Course) table() *gorm.DB {
	if !database.Connection.HasTable(c) {
		return database.Connection.CreateTable(c)
	}
	return database.Connection
}

// FindOne strudent from database
func FindOne(cond interface{}, args ...interface{}) (*Course, error) {
	c := &Course{}
	return c, c.table().Where(cond, args...).First(c).Error
}

// Find strudents from database
func Find(limit int, page int, cond interface{}, args ...interface{}) (*[]Course, error) {
	c := &Course{}
	courses := &[]Course{}
	return courses, c.table().Where(cond, args...).Limit(limit).Offset(page - 1).Find(courses).Error
}

// Save course to database
func (c *Course) Save() error {
	return c.table().Save(c).Error
}

// Delete course from database
func Delete(cond interface{}, args ...interface{}) error {
	c := &Course{}
	query := c.table().Where(cond, args...).Delete(c)

	if query.Error == nil && query.RowsAffected == 0 {
		return errors.New("record not found")
	}
	return query.Error
}

// Mini course data
func (c *Course) Mini() map[string]interface{} {
	return map[string]interface{}{
		"created_at": c.Model.CreatedAt,
		"name":       c.Name,
		"course_id":  c.CourseID,
		"unit":       c.Unit,
	}
}
