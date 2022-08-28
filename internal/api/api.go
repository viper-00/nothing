package api

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/viper-00/nothing/internal/config"
	"github.com/viper-00/nothing/internal/database"
	"github.com/viper-00/nothing/internal/logger"
	"github.com/viper-00/nothing/internal/monitor"
)

type Server struct{}

type Agents struct {
	AgentIDs []string
}

type CustomMetrics struct {
	CustomMetrics []string
}

func (s *Server) HandlePing(ctx context.Context, info *ServerInfo) (*Message, error) {
	config := config.GetConfig("config.json")
	err := handlePing(info.ServerName, &config)
	if err != nil {
		return &Message{Body: err.Error()}, err
	}

	return &Message{Body: "pong"}, nil
}

func handlePing(serverName string, config *config.Config) error {
	mysql := getMySQLConnection(config)
	defer mysql.Close()

	unixTime := strconv.FormatInt(time.Now().Unix(), 10)
	err := mysql.Ping(serverName, unixTime)
	if err != nil {
		logger.Log("error", err.Error())
		return fmt.Errorf("error saving ping from %s", serverName)
	}

	return nil
}

func (s *Server) IsUp(ctx context.Context, info *ServerInfo) (*IsActive, error) {
	config := config.GetConfig("config.json")
	upAddRunning, err := isUp(info.ServerName, &config)
	if err != nil {
		return &IsActive{IsUp: false}, err
	}

	return &IsActive{IsUp: upAddRunning}, nil
}

func isUp(serverName string, config *config.Config) (bool, error) {
	mysql := getMySQLConnection(config)
	defer mysql.Close()

	serverPingTimeStr, err := mysql.ServerPingTime(serverName)
	if err != nil {
		return false, fmt.Errorf("error loading ping time of %s", serverName)
	}

	serverPingTime, err := strconv.ParseInt(serverPingTimeStr, 10, 64)
	if err != nil {
		logger.Log("error", err.Error())
		return false, fmt.Errorf("error loading ping time of %s", serverName)
	}

	return time.Now().Unix()-serverPingTime <= 61, nil
}

func (s *Server) InitAgent(ctx context.Context, info *ServerInfo) (*Message, error) {
	config := config.GetConfig("config.json")
	err := initAgent(info.ServerName, info.Timezone, &config)
	if err != nil {
		return &Message{Body: err.Error()}, err
	}

	return &Message{Body: "agent added"}, nil
}

func initAgent(serverName, timeZone string, config *config.Config) error {
	logger.Log("info", "Initializing agent for "+serverName)

	mysql := getMySQLConnection(config)
	defer mysql.Close()

	if mysql.AgentIDExists(serverName) {
		logger.Log("error", "agent id "+serverName+" exists")
		return fmt.Errorf("agent id " + serverName + " exists")
	}

	err := mysql.AddAgent(serverName, timeZone)
	if err != nil {
		logger.Log("error", err.Error())
		return fmt.Errorf("error adding agent")
	}

	return nil
}

func (s *Server) HandleMonitorData(ctx context.Context, data *MonitorData) (*Message, error) {
	var monitorData = monitor.MonitorData{}
	err := json.Unmarshal([]byte(data.MonitorData), &monitorData)
	if err != nil {
		return &Message{Body: err.Error()}, err
	}

	err = handleMonitorData(&monitorData)
	if err != nil {
		return &Message{Body: err.Error()}, err
	}

	return &Message{Body: "ok"}, nil
}

func handleMonitorData(monitorData *monitor.MonitorData) error {
	serverName := monitorData.ServerId
	time := monitorData.UnixTime
	config := config.GetConfig("config.json")
	mysql := getMySQLConnection(&config)
	defer mysql.Close()

	data := make(map[string]interface{})
	data["system"] = &monitorData.System
	data["memory"] = &monitorData.Memory
	data["swap"] = &monitorData.Swap
	data["procUsage"] = &monitorData.ProcUsage
	data["processes"] = &monitorData.Processes

	for key, item := range data {
		err := saveToDB(item, mysql, serverName, time, key, "")
		if err != nil {
			return err
		}
	}

	for _, disk := range monitorData.Disk {
		err := saveToDB(disk, mysql, serverName, time, monitor.DISKS, disk.FileSystem)
		if err != nil {
			return err
		}
	}

	for _, service := range monitorData.Services {
		err := saveToDB(service, mysql, serverName, time, monitor.SERVICES, service.Name)
		if err != nil {
			return err
		}
	}

	for _, network := range monitorData.Networks {
		err := saveToDB(network, mysql, serverName, time, monitor.NETWORKS, network.Interface)
		if err != nil {
			return err
		}
	}

	return nil
}

func saveToDB(item interface{}, mysql database.MySql, serverName, time, key, logName string) error {
	res, err := json.Marshal(item)
	if err != nil {
		return err
	}

	err = mysql.SaveLogToDB(serverName, time, string(res), key, logName, false)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) HandleCustomMonitorData(ctx context.Context, data *MonitorData) (*Message, error) {
	return nil, nil
}

func (s *Server) HandleMonitorDataRequest(ctx context.Context, data *MonitorDataRequest) (*MonitorData, error) {
	return nil, nil
}

func (s *Server) HandleCustomMetricNameRequest(ctx context.Context, info *ServerInfo) (*Message, error) {
	return nil, nil
}

func (s *Server) HandleAgentIdsRequest(ctx context.Context, void *Void) (*Message, error) {
	return nil, nil
}

func (s *Server) mustEmbedUnimplementedMonitorDataServiceServer() {}

func getMySQLConnection(config *config.Config) database.MySql {
	mysql := database.MySql{}
	password := os.Getenv("NOTHING_MYSQL_PSWD")
	mysql.Connect(config.MySQLHost, config.MySQLDatabaseName, config.MySQLUserName, password, false)
	return mysql
}
