package cource

import (
	"errors"

	"github.com/MahdiRazaqi/university-system/database"
	"github.com/jinzhu/gorm"
)

// Cource model
type Cource struct {
	gorm.Model
	Name     string
	CourceID int
	Unit     int
}

func (c *Cource) table() *gorm.DB {
	if !database.Connection.HasTable(c) {
		return database.Connection.CreateTable(c)
	}
	return database.Connection
}

// FindOne strudent from database
func FindOne(cond interface{}, args ...interface{}) (*Cource, error) {
	c := &Cource{}
	return c, c.table().Where(cond, args...).First(c).Error
}

// Find strudents from database
func Find(limit int, page int, cond interface{}, args ...interface{}) (*[]Cource, error) {
	c := &Cource{}
	cources := &[]Cource{}
	return cources, c.table().Where(cond, args...).Limit(limit).Offset(page - 1).Find(cources).Error
}

// Save cource to database
func (c *Cource) Save() error {
	return c.table().Save(c).Error
}

// Delete cource from database
func Delete(cond interface{}, args ...interface{}) error {
	c := &Cource{}
	query := c.table().Where(cond, args...).Delete(c)

	if query.Error == nil && query.RowsAffected == 0 {
		return errors.New("record not found")
	}
	return query.Error
}
