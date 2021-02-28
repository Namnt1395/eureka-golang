package main

import (
	"demo-eureka/eureka"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	//handleSigterm1() // Graceful shutdown on Ctrl+C or kill
	eureka.Register("STARTING", "8091") // Performs Eureka registration
	go eureka.StartHeartbeat()
	eureka.SetEurekaStatus("UP", "8091")
	time.Sleep(time.Second * 2)
	//for {
	//response, err := http.Get("http://157.230.53.38/api-demo")
	//if err != nil {
	//	fmt.Print(err.Error())
	//	os.Exit(1)
	//}

	//responseData, err := ioutil.ReadAll(response.Body)
	//if err != nil {
	//	log.Fatal(err)
	//}
	fmt.Println("IP", getIP())
	var client = &http.Client{}
	var transport http.RoundTripper = &http.Transport{
		DisableKeepAlives: true,
	}
	client.Transport = transport
	req, _ := http.NewRequest("GET", "http://10.42.243.138:8080/info", nil)
	client = &http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err1 := client.Do(req)
	if err1 != nil {
		fmt.Println("Err...", err1.Error())
		os.Exit(1)
	}
	fmt.Println("RESS.....", resp)
	//upAction := eureka.HttpAction{ // Build a HttpAction struct
	//	Url:    "http://10.42.146.75/", // Note hard-coded path to Eureka...
	//	Method: "GET",
	//}
	//result := eureka.DoHttpRequest(upAction)
	//fmt.Println("api....." ,result)
	//}

	wg := sync.WaitGroup{} // Use a WaitGroup to block main() exit
	wg.Add(1)
	wg.Wait()
}
func handleSigterm1() {
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

func getIP() string {
	conn, err := net.Dial("tcp", "http://10.42.243.138:8761/")
	fmt.Println("Conn", conn)
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "error"
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	panic("Unable to determine local IP address (non loopback). Exiting.")
}