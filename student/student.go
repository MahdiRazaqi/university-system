package student

import (
	"errors"

	"github.com/MahdiRazaqi/university-system/database"
	"github.com/jeyem/passwd"
	"github.com/jinzhu/gorm"
)

// Student model
type Student struct {
	gorm.Model
	FirstName string
	LastName  string
	StudentID int
	Password  string
}

func (s *Student) table() *gorm.DB {
	if !database.Connection.HasTable(s) {
		return database.Connection.CreateTable(s)
	}
	return database.Connection
}

// FindOne strudent from database
func FindOne(cond interface{}, args ...interface{}) (*Student, error) {
	s := &Student{}
	return s, s.table().Where(cond, args...).First(s).Error
}

// Find strudents from database
func Find(limit int, page int, cond interface{}, args ...interface{}) (*[]Student, error) {
	s := &Student{}
	students := &[]Student{}
	return students, s.table().Where(cond, args...).Limit(limit).Offset(page - 1).Find(students).Error
}

// Save student to database
func (s *Student) Save() error {
	return s.table().Save(s).Error
}

// Delete student from database
func Delete(cond interface{}, args ...interface{}) error {
	s := &Student{}
	query := s.table().Where(cond, args...).Delete(s)

	if query.Error == nil && query.RowsAffected == 0 {
		return errors.New("record not found")
	}
	return query.Error
}

// Mini student data
func (s *Student) Mini() map[string]interface{} {
	return map[string]interface{}{
		"id":         s.Model.ID,
		"created_at": s.Model.CreatedAt,
		"first_name": s.FirstName,
		"last_name":  s.LastName,
		"student_id": s.StudentID,
	}
}

// LoadByStudentID load student with student id
func LoadByStudentID(studentID int) (*Student, error) {
	return FindOne("student_id = ?", studentID)
}

// Auth authenticate student with student id and password
func Auth(studentID int, password string) (*Student, error) {
	err := errors.New("student id or password not matched")

	s, err := LoadByStudentID(studentID)
	if err != nil {
		return nil, err
	}

	if !passwd.Check(password, s.Password) {
		return nil, err
	}
	return s, nil
}
