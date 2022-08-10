package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Port           string
	LogFileEnabled bool
	LogFilePath    string
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
