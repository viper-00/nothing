package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/viper-00/nothing/collector/internal/config"
	"github.com/viper-00/nothing/internal/alerts"
	"github.com/viper-00/nothing/internal/database"
	"github.com/viper-00/nothing/internal/logger"
	"github.com/viper-00/nothing/pkg/memdb"
)

func HandleAlerts(alertConfigs []alerts.AlertConfig, config *config.Config, mysql *database.MySql) {
	var (
		wg sync.WaitGroup
	)

	_ = mysql.ClearAllAlertsWithNullEnd()
	ticker := time.NewTicker(15 * time.Second)
	quit := make(chan struct{})
	wg.Add(1)
	go func() {
		incidentTracker := memdb.CreateDatabase("incident_tracker")
		err := incidentTracker.Create(
			"alert",
			memdb.Col{Name: "server_name", Type: memdb.String},
			memdb.Col{Name: "metric_type", Type: memdb.String},
			memdb.Col{Name: "metric_name", Type: memdb.String},
			memdb.Col{Name: "time", Type: memdb.String},
			memdb.Col{Name: "status", Type: memdb.Int},
			memdb.Col{Name: "value", Type: memdb.Float32},
		)
		if err != nil {
			logger.Log("error", "memdb"+err.Error())
		}

		for {
			select {
			case <-ticker.C:
				for _, alert := range alertConfigs {
					for _, server := range alert.Servers {
						processAlert(&alert, server, config, mysql, &incidentTracker)
					}
				}
				break
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
	wg.Wait()
	fmt.Println("Exsiting")
}

func processAlert(alert *alerts.AlertConfig, server string, config *config.Config, mysql *database.MySql, incidentTracker *memdb.Database) {

}
