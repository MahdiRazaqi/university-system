package obtained

import (
	"errors"

	"github.com/MahdiRazaqi/university-system/database"
	"github.com/jinzhu/gorm"
)

// Obtained model
type Obtained struct {
	gorm.Model
	SID uint
	PID uint
}

func (o *Obtained) table() *gorm.DB {
	if !database.Connection.HasTable(o) {
		return database.Connection.CreateTable(o)
	}
	return database.Connection
}

// FindOne strudent from database
func FindOne(cond interface{}, args ...interface{}) (*Obtained, error) {
	o := &Obtained{}
	return o, o.table().Where(cond, args...).First(o).Error
}

// Find strudents from database
func Find(limit int, page int, cond interface{}, args ...interface{}) (*[]Obtained, error) {
	o := &Obtained{}
	obtaineds := &[]Obtained{}
	return obtaineds, o.table().Where(cond, args...).Limit(limit).Offset(page - 1).Find(obtaineds).Error
}

// Save obtained to database
func (o *Obtained) Save() error {
	return o.table().Save(o).Error
}

// Delete obtained from database
func Delete(cond interface{}, args ...interface{}) error {
	o := &Obtained{}
	query := o.table().Where(cond, args...).Delete(o)

	if query.Error == nil && query.RowsAffected == 0 {
		return errors.New("record not found")
	}
	return query.Error
}
