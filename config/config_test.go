package config

import "testing"

func TestLoad(t *testing.T) {
	path := "../config.json"

	if err := Load(path); err != nil {
		t.Error()
	}
}
