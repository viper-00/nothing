package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	LogFileEnabled    bool
	LogFilePath       string
	Port              string
	MySQLHost         string
	MySQLUserName     string
	MySQLDatabaseName string
	TLSEnabled        bool
	KeyPath           string
	CertPath          string
}

func GetConfig(path string) Config {
	config := Config{}
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return config
	}

	_ = json.Unmarshal([]byte(file), &config)
	return config
}
