package database

import (
	"github.com/MahdiRazaqi/university-system/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"    //mysql driver
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres driver
)

// Connection database
var Connection *gorm.DB

func connectToMysql() error {
	conf := config.Config.DBMS

	db, err := gorm.Open(conf.Name, conf.User+":"+conf.Password+"@tcp("+conf.Host+")/"+conf.DB+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return err
	}
	Connection = db

	return nil
}
