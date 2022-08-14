package main

import (
	"fmt"

	"github.com/viper-00/nothing/collector/internal/config"
	"github.com/viper-00/nothing/internal/alerts"
	"github.com/viper-00/nothing/internal/database"
)

func HandleAlerts(alertConfigs []alerts.AlertConfig, config *config.Config, mysql *database.MySql) {
	fmt.Println("Exiting")
}
