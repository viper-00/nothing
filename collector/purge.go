package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/viper-00/nothing/collector/internal/config"
	"github.com/viper-00/nothing/internal/database"
	"github.com/viper-00/nothing/internal/logger"
)

func HandleDataPurge(config *config.Config, mysql *database.MySql) {
	var wg sync.WaitGroup

	ticker := time.NewTicker(6 * time.Second)
	quit := make(chan struct{})
	wg.Add(1)

	go func() {
		for {
			select {
			case <-ticker.C:
				purgeData := time.Now().AddDate(0, 0, -int(config.DataRetentionDays))
				unixTime := strconv.FormatInt(purgeData.Unix(), 10)
				affectedRows, err := mysql.PurgeMonitorDataOlderThan(unixTime)
				if err != nil {
					logger.Log("error", "data-purge: "+err.Error())
				}
				logger.Log("info", "data-purge: purged "+strconv.FormatInt(affectedRows, 10)+" rows")
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	wg.Wait()
	fmt.Println("Exiting")
}
