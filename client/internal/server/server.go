package server

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/viper-00/nothing/internal/logger"
)

// Run starts the server in given port
func Run(port string) {
	handleRequests(port)
}

func handleRequests(port string) {
	router := mux.NewRouter().StrictSlash(true)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend/")))

	server := http.Server{}
	server.Addr = port
	server.Handler = handlers.CompressHandler(router)
	server.SetKeepAlivesEnabled(false)

	logger.Log("info", "API started on port "+port)
	log.Fatal(server.ListenAndServe())
}
