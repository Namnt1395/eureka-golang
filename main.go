package main

import (
	"encoding/xml"
	"eureka-golang/eureka"
	"eureka-golang/object"
	"eureka-golang/service"
	"fmt"
	"io/ioutil"
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
	resp, err := http.Get("http://157.230.53.38:8761/eureka/apps/api-demo")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// Convert the body to type string
	//sb :=  string(body)

	var instance object.ApplicationsXml
	err1 := xml.Unmarshal(body, &instance.Application)
	if err1 != nil {
		fmt.Println(err1.Error())
	}
	fmt.Println("Data...", instance.Application.Name)

	listInstance := instance.Application
	for _, v := range listInstance.Instance {
		fmt.Println("Data...", v.HostName)
	}
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
