package main

import (
	"fmt"

	"github.com/PawlikMateusz/TWF0ZXVzeiBHb2dvQXBwcyBOQVNB/internal/config"
	"github.com/PawlikMateusz/TWF0ZXVzeiBHb2dvQXBwcyBOQVNB/internal/server/http/handlers"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	// load configuration
	log.Debug("Staring app")
	log.Debug("Loading configuration")
	config.SetDefaults()
	if err := config.LoadEnvVars(); err != nil {
		log.Fatal("Failed to load env variables")
	}

	// setup routes
	picturesHandler := handlers.PicturesHandler{}

	router := gin.Default()
	router.GET("/pictures", picturesHandler.Get)

	// run http server
	if err := router.Run(fmt.Sprintf(":%d", config.GetPort())); err != nil {
		log.Fatal(err)
	}
}
