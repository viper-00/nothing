package main

import (
	"sync"

	"github.com/viper-00/nothing/client/internal/server"
	"github.com/viper-00/nothing/internal/config"
)

func main() {
	config := config.GetConfig("config.json")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		server.Run(":" + config.Port)
		wg.Done()
	}()
	wg.Wait()
}
