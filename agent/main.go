package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"

	"github.com/dhamith93/systats"
	"github.com/viper-00/nothing/internal/api"
	"github.com/viper-00/nothing/internal/auth"
	"github.com/viper-00/nothing/internal/config"
	"github.com/viper-00/nothing/internal/logger"
	"github.com/viper-00/nothing/internal/monitor"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/internal/metadata"
	"google.golang.org/grpc/metadata"
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
	conn, c, ctx, cancel := createClient(config)
	if conn == nil {
		logger.Log("error", "error crating connection")
		return
	}

	defer conn.Close()
	defer cancel()

	syStats := systats.New()
	response, err := c.InitAgent(ctx, &api.ServerInfo{
		ServerName: config.ServerId,
		Timezone:   monitor.GetSystem(&syStats).TimeZone,
	})

	if err != nil {
		logger.Log("error", "error adding agent: "+err.Error())
		os.Exit(1)
	}
	fmt.Printf("%s \n", response.Body)
}

func sendCustomMetric(name, unit, value string, config *config.Config) {

}

func createClient(config *config.Config) (*grpc.ClientConn, api.MonitorDataServiceClient, context.Context, context.CancelFunc) {
	var (
		conn     *grpc.ClientConn
		tlsCreds credentials.TransportCredentials
		err      error
	)

	if len(config.CollectorEndpointCACertPath) > 0 {
		tlsCreds, err = loadTLSCreds(config)
		if err != nil {
			log.Fatal("cannot load TLS credentials: ", err)
		}
		conn, err = grpc.Dial(config.CollectorEndpoint, grpc.WithTransportCredentials(tlsCreds))
	} else {
		conn, err = grpc.Dial(config.CollectorEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	if err != nil {
		logger.Log("error", "connection error: "+err.Error())
		return nil, nil, nil, nil
	}

	c := api.NewMonitorDataServiceClient(conn)
	token := generateToken()
	ctx, cancel := context.WithTimeout(metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{"jwt": token})), time.Second*10)
	return conn, c, ctx, cancel
}

func loadTLSCreds(config *config.Config) (credentials.TransportCredentials, error) {
	cert, err := ioutil.ReadFile(config.CollectorEndpointCACertPath)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(cert) {
		return nil, fmt.Errorf("failed to add server CA cert")
	}

	tlsConfig := &tls.Config{
		RootCAs: certPool,
	}

	return credentials.NewTLS(tlsConfig), nil
}

func generateToken() string {
	token, err := auth.GenerateJWT()
	if err != nil {
		logger.Log("error", "error generating token: "+err.Error())
		os.Exit(1)
	}
	return token
}
