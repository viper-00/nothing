package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Port                        string
	LogFileEnabled              bool
	LogFilePath                 string
	CollectorEndpointCACertPath string
	CollectorEndpoint           string
	AlertEndpointCACertPath     string
	AlertEndpoint               string
	MonitorIntervalSeconds      int
	TLSEnabled                  bool
	CertPath                    string
	KeyPath                     string

	ServerId string
	Services []ServiceToMonitor

	DisksTOIgnore string

	MySQLUserName     string
	MySQLHost         string
	MySQLDatabaseName string
}

// ServiceToMonitor holds service info from config.json
type ServiceToMonitor struct {
	Name        string
	ServiceName string
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
