package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/viper-00/nothing/internal/config"
	"github.com/viper-00/nothing/internal/logger"
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
	return "", nil
}
