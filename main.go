package main

import (
	"fmt"
	"net/http"

	"github.com/xakepp35/xai_models/controllers"
	"github.com/xakepp35/xai_models/env"

	log "github.com/sirupsen/logrus"
)

func main() {
	fmt.Printf("Starting %s version %s on port: %s\n", env.ServiceName, env.ServiceVersion, env.Port)

	http.HandleFunc("/api/ping", controllers.Ping())
	http.HandleFunc("/api/graph", controllers.Graph())
	err := http.ListenAndServe(":"+env.Port, nil)
	if err != nil {
		log.Fatalf("Error http.ListenAndServe failed: %s", err)
	}
}
