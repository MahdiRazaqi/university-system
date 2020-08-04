package database

import (
	"testing"

	"github.com/MahdiRazaqi/university-system/config"
)

func TestConnectToMysql(t *testing.T) {
	path := "../config.json"

	if err := config.Load(path); err != nil {
		t.Error(err)
	}

	if err := connectToMysql(); err != nil {
		t.Error(err)
	}
}
