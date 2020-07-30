package teacher

import (
	"errors"

	"github.com/MahdiRazaqi/university-system/database"
	"github.com/jinzhu/gorm"
)

// Teacher model
type Teacher struct {
	gorm.Model
	FirstName string
	LastName  string
	TeacherID int
	Password  string
}

func (t *Teacher) table() *gorm.DB {
	if !database.Connection.HasTable(t) {
		return database.Connection.CreateTable(t)
	}
	return database.Connection
}

// FindOne teacher from database
func FindOne(cond interface{}, args ...interface{}) (*Teacher, error) {
	t := &Teacher{}
	return t, t.table().Where(cond, args...).First(t).Error
}

// Find teachers from database
func Find(limit int, page int, cond interface{}, args ...interface{}) (*[]Teacher, error) {
	t := &Teacher{}
	teachers := &[]Teacher{}
	return teachers, t.table().Where(cond, args...).Limit(limit).Offset(page - 1).Find(teachers).Error
}

// Save teacher to database
func (t *Teacher) Save() error {
	return t.table().Save(t).Error
}

// Delete teacher from database
func Delete(cond interface{}, args ...interface{}) error {
	t := &Teacher{}
	query := t.table().Where(cond, args...).Delete(t)

	if query.Error == nil && query.RowsAffected == 0 {
		return errors.New("record not found")
	}
	return query.Error
}

// Mini teacher data
func (t *Teacher) Mini() map[string]interface{} {
	return map[string]interface{}{
		"id":         t.Model.ID,
		"created_at": t.Model.CreatedAt,
		"first_name": t.FirstName,
		"last_name":  t.LastName,
		"teacher_id": t.TeacherID,
	}
}
