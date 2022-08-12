package server

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/viper-00/nothing/internal/api"
	"github.com/viper-00/nothing/internal/auth"
	"github.com/viper-00/nothing/internal/config"
	"github.com/viper-00/nothing/internal/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

type output struct {
	Status string
	Data   interface{}
}

// Run starts the server in given port
func Run(port string) {
	handleRequests(port)
}

func handleRequests(port string) {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/network", returnNetwork)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend/")))

	server := http.Server{}
	server.Addr = port
	server.Handler = handlers.CompressHandler(router)
	server.SetKeepAlivesEnabled(false)

	logger.Log("info", "API started on port "+port)
	log.Fatal(server.ListenAndServe())
}

func returnNetwork(w http.ResponseWriter, r *http.Request) {
	handleRequest("networks", w, r, false)
}

func handleRequest(logType string, w http.ResponseWriter, r *http.Request, isCustomMetric bool) {
	w.Header().Set("Content-Type", "application/json")
	config := config.GetConfig("config.json")
	serverName, _ := parseGETForServerName(r)
	time, _ := parseGETForTime(r)
	from, to, _ := parseGETForDates(r)
	received, err := getMonitorData(serverName, logType, from, to, time, &config, isCustomMetric)

	// Return data
	var data interface{}
	var out output
	out.Status = "OK"
	if err != nil {
		out.Status = "ERR"
		json.NewEncoder(w).Encode(&out)
	}
	_ = json.Unmarshal([]byte(received), &data)
	out.Data = data
	json.NewEncoder(w).Encode(&out)
}

func parseGETForServerName(r *http.Request) (string, error) {
	serverIdArr, ok := r.URL.Query()["serverId"]
	if !ok || len(serverIdArr) == 0 {
		logger.Log("ERROR", "cannot parse for server ID")
		return "", fmt.Errorf("cannot parse for server id")
	}
	return serverIdArr[0], nil
}

func parseGETForTime(r *http.Request) (int64, error) {
	timeArr, ok := r.URL.Query()["time"]
	if !ok {
		return 0, fmt.Errorf("error parsing get vars")
	}
	timeInt, err := strconv.ParseInt(timeArr[0], 10, 64)
	if err != nil {
		return 0, fmt.Errorf("error parsing get vars")
	}
	return timeInt, nil
}

func parseGETForDates(r *http.Request) (int64, int64, error) {
	from, okFrom := r.URL.Query()["from"]
	to, okTo := r.URL.Query()["to"]

	if !okFrom || !okTo {
		return 0, 0, fmt.Errorf("error parsing get vars")
	}

	fromTime, err1 := strconv.ParseInt(from[0], 10, 64)
	toTime, err2 := strconv.ParseInt(to[0], 10, 64)

	if err1 != nil || err2 != nil {
		return 0, 0, fmt.Errorf("error parsing get vars")
	}

	return fromTime, toTime, nil
}

func getMonitorData(serverName string, logType string, from int64, to int64, time int64, config *config.Config, isCustomMetric bool) (string, error) {
	conn, c, ctx, cancel := createClient(config)
	defer conn.Close()
	defer cancel()
	monitorData, err := c.HandleMonitorDataRequest(ctx, &api.MonitorDataRequest{ServerName: serverName, LogType: logType, From: from, To: to, Time: time, IsCustomMetric: isCustomMetric})
	if err != nil {
		logger.Log("error", "error sending data: "+err.Error())
		return "", err
	}
	return monitorData.MonitorData, nil
}

func createClient(config *config.Config) (*grpc.ClientConn, api.MonitorDataServiceClient, context.Context, context.CancelFunc) {
	var (
		conn     *grpc.ClientConn
		tlsCreds credentials.TransportCredentials
		err      error
	)

	if len(config.CollectorEndpointCACertPath) > 0 {
		tlsCreds, err = loadTLSCreds(config.CollectorEndpointCACertPath)
		if err != nil {
			log.Fatal("cannot load TLS credentials: ", err)
		}
		conn, err = grpc.Dial(config.CollectorEndpoint, grpc.WithTransportCredentials(tlsCreds))
	} else {
		conn, err = grpc.Dial(config.CollectorEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	if err != nil {
		logger.Log("error", "connection error: "+err.Error())
		os.Exit(1)
	}

	c := api.NewMonitorDataServiceClient(conn)
	token := generateToken()
	ctx, cancel := context.WithTimeout(metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{"jwt": token})), time.Second*10)
	return conn, c, ctx, cancel
}

func generateToken() string {
	token, err := auth.GenerateJWT()
	if err != nil {
		logger.Log("error", "error generating token: "+err.Error())
	}
	return token
}

func loadTLSCreds(path string) (credentials.TransportCredentials, error) {
	cert, err := ioutil.ReadFile(path)
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
