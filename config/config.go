package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/sirupsen/logrus"
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

const configPath = "./config.json"

// Config data
var Config config

// Load config file
func Load() {
	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		logrus.Error("loading config file ", err)
	}

	if err := json.Unmarshal(file, &Config); err != nil {
		logrus.Error("loading config file ", err)
	}
}
