package main

import (
	"countryes/router"
	"log"
	"net/http"
)

const (
	host = "127.0.0.1"
	port = ":8090"
)

func main() {
	r, err := router.Startup()
	if err != nil {
		log.Println("Router start failed: " + err.Error())
	}

	log.Println("Web server starded on port " + port)

	err = http.ListenAndServe(host+port, r)
	if err != nil {
		log.Fatal("Server failed to start")
	}
}
