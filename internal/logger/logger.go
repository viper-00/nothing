package logger

import (
	"log"

	"github.com/viper-00/nothing/internal/config"
)

func Log(prefix string, msg string) {
	if !config.GetConfig("config.json").LogFileEnabled {
		return
	}
	log.Println(prefix + " " + msg)
}
