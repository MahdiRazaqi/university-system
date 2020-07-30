package database

import (
	"github.com/MahdiRazaqi/university-system/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"    //mysql driver
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres driver
	"github.com/sirupsen/logrus"
)

// Connection database
var Connection *gorm.DB

// Connect to database
func Connect() {
	conf := config.Config.DBMS

	db, err := gorm.Open(conf.Name, conf.User+":"+conf.Password+"@tcp("+conf.Host+")/"+conf.DB+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		logrus.Error(err)
	}
	Connection = db
}
