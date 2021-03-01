package main

import (
	"eureka-golang/eureka"
	"eureka-golang/handle"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//var isStartService *bool

func main() {
	isStartService := make(chan bool)
	portApi := "8080"
	/**
	Connect Eureka service
	*/
	eureka.HandleSigterm()                     // Graceful shutdown on Ctrl+C or kill
	eureka.RegisterEureka("STARTING", portApi) // Performs Eureka registration
	go eureka.StartHeartbeat()
	// Reconnect when eureka server stop
	go eureka.CheckServiceLive(portApi, isStartService)
	/**
	End info Eureka
	*/
	startWebServer(portApi, isStartService)

}

func startWebServer(port string, isStart chan bool) {
	router := mux.NewRouter()
	log.Println("Starting HTTP service at 8080")
	router.HandleFunc("/api", handle.ApiDemo)
	eureka.SetEurekaStatus("UP", "8080")
	isStart <- true
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Println("An error occured starting HTTP listener at port 8080")
		log.Println("Error: " + err.Error())
	}
}
