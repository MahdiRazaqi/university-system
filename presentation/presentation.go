package presentation

import (
	"errors"
	"time"

	"github.com/MahdiRazaqi/university-system/database"
	"github.com/jinzhu/gorm"
)

// Presentation model
type Presentation struct {
	gorm.Model
	TID      uint
	CID      uint
	Time     time.Time
	Capacity int
}

func (p *Presentation) table() *gorm.DB {
	if !database.Connection.HasTable(p) {
		return database.Connection.CreateTable(p)
	}
	return database.Connection
}

// FindOne strudent from database
func FindOne(cond interface{}, args ...interface{}) (*Presentation, error) {
	p := &Presentation{}
	return p, p.table().Where(cond, args...).First(p).Error
}

// Find strudents from database
func Find(limit int, page int, cond interface{}, args ...interface{}) (*[]Presentation, error) {
	p := &Presentation{}
	presentations := &[]Presentation{}
	return presentations, p.table().Where(cond, args...).Limit(limit).Offset(page - 1).Find(presentations).Error
}

// Save presentation to database
func (p *Presentation) Save() error {
	return p.table().Save(p).Error
}

// Delete presentation from database
func Delete(cond interface{}, args ...interface{}) error {
	p := &Presentation{}
	query := p.table().Where(cond, args...).Delete(p)

	if query.Error == nil && query.RowsAffected == 0 {
		return errors.New("record not found")
	}
	return query.Error
}
