package main

import (
	"eureka-golang/eureka"
	"eureka-golang/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//handleSigterm()                           // Graceful shutdown on Ctrl+C or kill
	//eureka.RegisterEureka("STARTING", "8080") // Performs Eureka registration
	//go eureka.StartHeartbeat()
	//
	//startWebServer("8080")
	//
	//wg := sync.WaitGroup{} // Use a WaitGroup to block main() exit
	//wg.Add(1)
	//wg.Wait()

	//fmt.Println(sb)

	//	log.Printf(sb)
}

func startWebServer(port string) {
	router := service.NewRouter()
	log.Println("Starting HTTP service at 8080")
	eureka.SetEurekaStatus("UP", "8080")
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Println("An error occured starting HTTP listener at port 8080")
		log.Println("Error: " + err.Error())
	}
	//router := mux.NewRouter()
	//log.Println("Starting HTTP service at 8080")
	//router.HandleFunc("/api", handle.ApiDemo)
	//eureka.SetEurekaStatus("UP", "8080")
	//err := http.ListenAndServe(":"+port, router)
	//if err != nil {
	//	log.Println("An error occured starting HTTP listener at port 8080")
	//	log.Println("Error: " + err.Error())
	//}
}

func handleSigterm() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		//eureka.Deregister()
		eureka.SetEurekaStatus("DOWN", "8080")
		os.Exit(1)
	}()
}
