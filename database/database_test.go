package database

import (
	"testing"

	"github.com/MahdiRazaqi/university-system/config"
)

func TestConnet(t *testing.T) {
	path := "../config.json"

	if err := config.Load(path); err != nil {
		t.Error()
	}

	if err := Connect(); err != nil {
		t.Error(err)
	}
}
