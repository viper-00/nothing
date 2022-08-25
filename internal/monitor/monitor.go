package monitor

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/dhamith93/systats"
	"github.com/viper-00/nothing/internal/config"
	"github.com/viper-00/nothing/internal/logger"
)

const (
	SYSTEN     string = "system"
	MEMORY     string = "memory"
	SWAP       string = "swap"
	PROC_USAGE string = "procUsage"
	PROCESSES  string = "processes"
	DISKS      string = "disks"
	SERVICES   string = "services"
	NETWORKS   string = "networks"
)

// Service holds service activity information
type Service struct {
	Name    string
	Running bool
	Time    string
}

// Processes hold CPU and Memory usage data
type Processes struct {
	CPU    []systats.Process
	Memory []systats.Process
}

// MonitorData holds individual system stats
type MonitorData struct {
	UnixTime  string
	System    systats.System
	Memory    systats.Memory
	Swap      systats.Swap
	Disk      []systats.Disk
	ProcUsage systats.CPU
	Networks  []systats.Network
	Processes Processes
	Services  []Service
	ServerId  string
}

func MonitorAsJSON(config *config.Config) string {
	monitorData := Monitor(config)
	monitorData.ServerId = config.ServerId
	jsonData, err := json.Marshal(&monitorData)
	if err != nil {
		logger.Log("Error", err.Error())
		return ""
	}
	return string(jsonData)
}

func Monitor(config *config.Config) MonitorData {
	syStats := systats.New()
	unixTime := strconv.FormatInt(time.Now().Unix(), 10)
	system := GetSystem(&syStats)
	memory := GetMemory(&syStats)
	swap := GetSwap(&syStats)
	disks := GetDisks(&syStats, config)
	procUsage := GetProcessor(&syStats)
	networks := GetNetwork(&syStats)
	processes := GetProcesses(&syStats)
	services := GetServices(&syStats, unixTime, config)

	return MonitorData{
		UnixTime:  unixTime,
		System:    system,
		Memory:    memory,
		Swap:      swap,
		Disk:      disks,
		ProcUsage: procUsage,
		Networks:  networks,
		Processes: processes,
		Services:  services,
	}
}

// GetSystem returns a systats.SyStats struct with system info
func GetSystem(syStats *systats.SyStats) systats.System {
	// Operating System
	system, err := syStats.GetSystem()
	if err != nil {
		logger.Log("error", err.Error())
	}

	return system
}

// Getmemory returns a systats.Memory struct with memory usage
func GetMemory(syStats *systats.SyStats) systats.Memory {
	// Select byte format for kb or mb
	memory, err := syStats.GetMemory(systats.Megabyte)
	if err != nil {
		logger.Log("error", err.Error())
	}
	return memory
}

// GetSwap returns systats.Swap struct with swap usage
func GetSwap(syStats *systats.SyStats) systats.Swap {
	swap, err := syStats.GetSwap(systats.Megabyte)
	if err != nil {
		logger.Log("error", err.Error())
	}
	return swap
}

// GetDisk returns array of systats.Disk structs with disk info and usage data
func GetDisks(syStats *systats.SyStats, config *config.Config) []systats.Disk {
	disks, err := syStats.GetDisks()
	if err != nil {
		logger.Log("error", err.Error())
	}
	output := []systats.Disk{}
	disksTOIgnore := strings.Split(config.DisksTOIgnore, ",")

	for _, disk := range disks {
		ignore := false
		for _, diskToIgnore := range disksTOIgnore {
			if disk.FileSystem == strings.TrimSpace(diskToIgnore) {
				ignore = true
			}
		}
		if !ignore {
			output = append(output, disk)
		}
	}

	return output
}

// GetProcessor returns a systats.CPU struct with CPU info and usage data
func GetProcessor(syStats *systats.SyStats) systats.CPU {
	cpu, err := syStats.GetCPU()
	if err != nil {
		logger.Log("error", err.Error())
	}
	return cpu
}

// GetNetwork returns an array of systats.Network struct with network usage
func GetNetwork(syStats *systats.SyStats) []systats.Network {
	networks, err := syStats.GetNetworks()
	if err != nil {
		logger.Log("error", err.Error())
	}

	return networks
}

// GetProcessor returns a systats.Processor struct with process info
func GetProcesses(syStats *systats.SyStats) Processes {
	cpu, err := syStats.GetTopProcesses(10, "cpu")
	if err != nil {
		logger.Log("error", err.Error())
	}

	memory, err := syStats.GetTopProcesses(10, "memory")
	if err != nil {
		logger.Log("error", err.Error())
	}

	return Processes{
		CPU:    cpu,
		Memory: memory,
	}
}

// GetServices returns array of Service structs with service status
func GetServices(syStats *systats.SyStats, unixTime string, config *config.Config) []Service {
	servicesToCheck := config.Services
	var services []Service

	for _, serviceToCheck := range servicesToCheck {
		services = append(services, Service{
			Name:    serviceToCheck.Name,
			Running: syStats.IsServiceRunning(serviceToCheck.Name),
			Time:    unixTime,
		})
	}

	return services
}
