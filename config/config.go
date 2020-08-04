package config

import (
	"encoding/json"
	"io/ioutil"
)

type config struct {
	SecretKey string `json:"secret_key"`
	DBMS      struct {
		Name     string `json:"name"`
		Host     string `json:"host"`
		DB       string `json:"db"`
		User     string `json:"user"`
		Password string `json:"password"`
	} `json:"dbms"`
}

var configPath = "./config.json"

// Config data
var Config config

// Load config file
func Load(path ...string) error {
	if len(path) != 0 {
		configPath = path[0]
	}

	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(file, &Config); err != nil {
		return err
	}
	return nil
}
