package main

import (
	"context"
	"crypto/tls"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/viper-00/nothing/collector/internal/config"
	"github.com/viper-00/nothing/internal/alerts"
	"github.com/viper-00/nothing/internal/api"
	"github.com/viper-00/nothing/internal/auth"
	"github.com/viper-00/nothing/internal/database"
	"github.com/viper-00/nothing/internal/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func main() {
	var configPath, alertConfigPath, removeAgentVal string
	var alertConfig []alerts.AlertConfig

	initPtr := flag.Bool("init", false, "Initialize the collector")
	flag.StringVar(&configPath, "config", "", "Path to config json file.")
	flag.StringVar(&alertConfigPath, "alerts", "", "Path to alerts json file.")
	flag.StringVar(&removeAgentVal, "remove-agent", "", "Remove agent info from collector DB. Agent monitor data is not deleted.")
	flag.Parse()

	if _, err := os.Stat(configPath); errors.Is(err, os.ErrNotExist) {
		log.Fatalf("cannot load config json: %v", err)
	}

	config := config.GetConfig(configPath)
	if config.LogFileEnabled {
		file, err := os.OpenFile(config.LogFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		log.SetOutput(file)
	}

	err := godotenv.Load()
	if err != nil {
		logger.Log("Error", "Error loading .env file")
	}

	if len(alertConfigPath) > 0 {
		if _, err := os.Stat(alertConfigPath); errors.Is(err, os.ErrNotExist) {
			logger.Log("cannot load alert config: ", err.Error())
		}
		alertConfig = alerts.GetAlertConfig(alertConfigPath)
	}

	if *initPtr {
		initCollector(&config)
	} else if len(removeAgentVal) > 0 {
		removeAgent(removeAgentVal, &config)
	} else {
		mysql := getMysqlConnection(&config, false)
		defer mysql.Close()

		if alertConfig != nil {
			go HandleAlerts(alertConfig, &config, &mysql)
		}

		// purge some data
		go HandleDataPurge(&config, &mysql)

		listen, err := net.Listen("tcp", ":"+config.Port)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		server := api.Server{}
		var grpcServer *grpc.Server

		if config.TLSEnabled {
			tlsCreds, err := loadTLSCreds(&config)
			if err != nil {
				log.Fatal("cannot load TLS credentials: ", err)
				// log.Fatalf("failed to load TLS cert %s, key %s: %v", config.KeyPath, config.KeyPath, err)
			}
			grpcServer = grpc.NewServer(grpc.Creds(tlsCreds), grpc.UnaryInterceptor(authInterceptor))
		} else {
			grpcServer = grpc.NewServer(grpc.UnaryInterceptor(authInterceptor))
		}

		api.RegisterMonitorDataServiceServer(grpcServer, &server)
		if err := grpcServer.Serve(listen); err != nil {
			log.Fatalf("failed to server: %s", err)
		}
	}
}

func initCollector(config *config.Config) {
	mysql := getMysqlConnection(config, true)
	defer mysql.Close()
	err := mysql.Init()
	if err != nil {
		fmt.Println(err.Error())
	}
	auth.GetKey()
}

func getMysqlConnection(config *config.Config, isMultiStatement bool) database.MySql {
	mysql := database.MySql{}
	password := os.Getenv("MYSQL_PSWD")
	mysql.Connect(config.MySQLHost, config.MySQLDatabaseName, config.MySQLUserName, password, isMultiStatement)
	return mysql
}

func removeAgent(removeAgentVal string, config *config.Config) {
	fmt.Println("Removing agent " + removeAgentVal)
	mysql := getMysqlConnection(config, false)
	defer mysql.Close()

	if mysql.SqlErr != nil {
		fmt.Println(mysql.SqlErr.Error())
		return
	}

	if !mysql.AgentIDExists(removeAgentVal) {
		fmt.Println("Agent ID " + removeAgentVal + " doesn't exists... ")
		return
	}

	err := mysql.RemoveAgent(removeAgentVal)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func loadTLSCreds(config *config.Config) (credentials.TransportCredentials, error) {
	cert, err := tls.LoadX509KeyPair(config.CertPath, config.KeyPath)
	if err != nil {
		return nil, err
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.NoClientCert,
	}

	return credentials.NewTLS(tlsConfig), nil
}

func authInterceptor(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	meta, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		logger.Log("error", "cannot parse meta")
		return nil, status.Error(codes.Unauthenticated, "INTERNAL_SERVER_ERROR")
	}

	if len(meta["jwt"]) != 1 {
		logger.Log("error", "cannot parse meta - token empty")
		return nil, status.Error(codes.Unauthenticated, "token empty")
	}

	if !auth.ValidToken(meta["jwt"][0]) {
		logger.Log("error", "auth error")
		return nil, status.Error(codes.PermissionDenied, "invalid auth token")
	}

	return handler(ctx, req)
}
