package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/dhamith93/systats"
	"github.com/viper-00/nothing/collector/internal/config"
	"github.com/viper-00/nothing/internal/alerts"
	"github.com/viper-00/nothing/internal/alertstatus"
	"github.com/viper-00/nothing/internal/database"
	"github.com/viper-00/nothing/internal/logger"
	"github.com/viper-00/nothing/internal/monitor"
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
	metricType := alert.MetricName
	metricName := ""
	if metricType == monitor.DISKS {
		metricName = alert.Disk
	}
	if metricType == monitor.SERVICES {
		metricName = alert.Service
	}

	alertStatus := buildAlertStatus(alert, &server, config, mysql)

	alertFormatDbForStartEvent := mysql.GetAlertByStartEvent(strconv.FormatInt(alertStatus.StartEvent, 10))
	if alertFormatDbForStartEvent != nil {
		return
	}

	logger.Log(metricName, metricName)
}

func buildAlertStatus(alert *alerts.AlertConfig, server *string, config *config.Config, mysql *database.MySql) alertstatus.AlertStatus {
	var alertStatus alertstatus.AlertStatus
	logName := ""

	switch alert.MetricName {
	case monitor.DISKS:
		logName = alert.Disk
	case monitor.SERVICES:
		logName = alert.Service
	}

	metricLogs := mysql.GetLogFromDBWithId(*server, alert.MetricName, logName, 0, 0)
	logId := metricLogs[0][0]
	alertStatus.Alert = *alert
	alertStatus.Server = *server
	alertStatus.Type = alertstatus.Normal

	switch alert.MetricName {
	case monitor.PROC_USAGE:
		var cpu systats.CPU
		err := json.Unmarshal([]byte(metricLogs[0][1]), &cpu)
		if err != nil {
			logger.Log("error", err.Error())
			return alertStatus
		}
		alertStatus.UnixTime = strconv.FormatInt(cpu.Time, 10)
		alertStatus.Value = float32(cpu.LoadAvg)
		alertStatus.Type = getAlertType(alert, float64(cpu.LoadAvg))
	case monitor.MEMORY:
		var memory systats.Memory
		err := json.Unmarshal([]byte(metricLogs[0][1]), &memory)
		if err != nil {
			logger.Log("error", err.Error())
			return alertStatus
		}
		alertStatus.UnixTime = strconv.FormatInt(memory.Time, 10)
		alertStatus.Value = float32(memory.PercentageUsed)
		alertStatus.Type = getAlertType(alert, memory.PercentageUsed)
	case monitor.SWAP:
		var swap systats.Swap
		err := json.Unmarshal([]byte(metricLogs[0][1]), &swap)
		if err != nil {
			logger.Log("error", err.Error())
			return alertStatus
		}
		alertStatus.UnixTime = strconv.FormatInt(swap.Time, 10)
		alertStatus.Value = float32(swap.PercentageUsed)
		alertStatus.Type = getAlertType(alert, swap.PercentageUsed)
	case monitor.DISKS:
		var disk systats.Disk
		err := json.Unmarshal([]byte(metricLogs[0][1]), &disk)
		if err != nil {
			logger.Log("error", err.Error())
			return alertStatus
		}

		if disk.FileSystem == alert.Disk {
			valStr := strings.Replace(disk.Usage.Usage, "%", "", -1)
			val, err := strconv.ParseFloat(valStr, 32)
			if err != nil {
				logger.Log("error", err.Error())
				return alertStatus
			}
			alertStatus.UnixTime = strconv.FormatInt(disk.Time, 10)
			alertStatus.Value = float32(val)
			alertStatus.Type = getAlertType(alert, val)
		}
	case monitor.SERVICES:
		var service monitor.Service
		err := json.Unmarshal([]byte(metricLogs[0][1]), &service)
		if err != nil {
			logger.Log("error", err.Error())
			return alertStatus
		}

		if service.Name == alert.Service {
			val := 0.0
			if service.Running {
				val = 1.0
			}
			alertStatus.UnixTime = service.Time
			alertStatus.Value = float32(val)
			alertStatus.Type = getAlertType(alert, val)
		}
	}

	logIdInt, err := strconv.ParseInt(logId, 10, 64)
	if err != nil {
		logger.Log("error", err.Error())
		return alertStatus
	}

	alertStatus.StartEvent = logIdInt
	return alertStatus
}

func getAlertType(alert *alerts.AlertConfig, val float64) alertstatus.StatusType {
	switch alert.Op {
	case "==":
		if val == float64(alert.CriticalThreshold) {
			return alertstatus.Critical
		} else if val == float64(alert.WarnThreshold) {
			return alertstatus.Warning
		} else {
			return alertstatus.Normal
		}
	case "!=":
		if val != float64(alert.CriticalThreshold) {
			return alertstatus.Critical
		} else if val != float64(alert.WarnThreshold) {
			return alertstatus.Warning
		} else {
			return alertstatus.Normal
		}
	case ">":
		if val > float64(alert.CriticalThreshold) {
			return alertstatus.Critical
		} else if val > float64(alert.WarnThreshold) {
			return alertstatus.Warning
		} else {
			return alertstatus.Normal
		}
	case "<":
		if val < float64(alert.CriticalThreshold) {
			return alertstatus.Critical
		} else if val < float64(alert.WarnThreshold) {
			return alertstatus.Warning
		} else {
			return alertstatus.Normal
		}
	case ">=":
		if val >= float64(alert.CriticalThreshold) {
			return alertstatus.Critical
		} else if val >= float64(alert.WarnThreshold) {
			return alertstatus.Warning
		} else {
			return alertstatus.Normal
		}
	case "<=":
		if val <= float64(alert.CriticalThreshold) {
			return alertstatus.Critical
		} else if val <= float64(alert.WarnThreshold) {
			return alertstatus.Warning
		} else {
			return alertstatus.Normal
		}
	case "inactive":
		if val == 0.0 {
			return alertstatus.Critical
		}
	case "active":
		if val == 1.0 {
			return alertstatus.Critical
		}
	}
	return alertstatus.Normal
}
