package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/viper-00/nothing/internal/config"
	"github.com/viper-00/nothing/internal/monitor"
)

func main() {
	config := config.GetConfig("config.json")

	if config.LogFileEnabled {
		file, err := os.OpenFile(config.LogFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		log.SetOutput(file)
	}

	var name, value, unit string

	initPtr := flag.Bool("init", false, "Initialize agent")
	customPtr := flag.Bool("custom", false, "Send custom metrics")
	flag.StringVar(&name, "name", "", "Name of the metric")
	flag.StringVar(&value, "value", "", "Value of the metric")
	flag.StringVar(&unit, "unit", "", "Unit of the metric")
	flag.Parse()

	// Whether is the initialize
	if *initPtr {
		initAgent(&config)
		return
	} else if *customPtr {
		if len(name) > 0 && len(value) > 0 && len(unit) > 0 {
			sendCustomMetric(name, unit, value, &config)
		} else {
			fmt.Println("Metric name, unit, and value all must be required")
		}
		return
	}

	ticker := time.NewTicker(time.Duration(config.MonitorIntervalSeconds) * time.Second)
	tickerForPing := time.NewTicker(time.Minute)
	quit := make(chan struct{})
	quitForPing := make(chan bool)

	var wg sync.WaitGroup
	wg.Add(2)

	// Monitoring
	go func() {
		for {
			select {
			case <-ticker.C:
				monitorData := monitor.MonitorAsJSON(&config)
				sendMonitorData(monitorData, &config)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	// PING
	go func() {
		for {
			select {
			case <-tickerForPing.C:
				sendPing(&config)
			case <-quitForPing:
				ticker.Stop()
				return
			}
		}
	}()

	wg.Wait()
	fmt.Println("Exiting")
}

func initAgent(config *config.Config) {

}

func sendCustomMetric(name, unit, value string, config *config.Config) {

}
