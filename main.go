package main

import (
	"fmt"
	"net/http"

	"github.com/PawlikMateusz/TWF0ZXVzeiBHb2dvQXBwcyBOQVNB/internal/config"
	"github.com/PawlikMateusz/TWF0ZXVzeiBHb2dvQXBwcyBOQVNB/internal/server/http/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func main() {
	// load configuration
	log.Info("Loading configuration")
	config.SetDefaults()
	if err := config.LoadEnvVars(); err != nil {
		log.Fatal("Failed to load env variables")
	}

	// setup routes
	picturesHandler := handlers.PicturesHandler{}
	r := mux.NewRouter()
	r.HandleFunc("/pictures", picturesHandler.Get).Methods("GET")
	http.Handle("/", r)

	// run http server
	serverPort := config.GetPort()
	log.Infof("Starting server at port %d", serverPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", serverPort), nil); err != nil {
		log.Fatal(err)
	}
}
