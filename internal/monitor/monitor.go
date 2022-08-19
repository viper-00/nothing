package monitor

import "github.com/dhamith93/systats"

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

type Processes struct {
	CPU    []systats.Process
	Memory []systats.Process
}
