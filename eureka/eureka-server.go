package eureka

import (
	"encoding/xml"
	"eureka-golang/object"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

const (
	ApplicationName = "api-demo"
	ServerEureka    = "http://157.230.53.38:8761/eureka/apps/"
	Prefix          = "http://"
)

func getTemplateEureka() string {
	data, err1 := ioutil.ReadFile("eureka/templates/regtpl.json")
	if err1 != nil {
		fmt.Println(err1.Error())
	}
	return string(data)
}

func formatTemplate(status string, portApplication string) string {
	templateEureka := getTemplateEureka()
	templateEureka = strings.Replace(templateEureka, "${ipAddress}", GetLocalIP(), -1) // Replace some placeholders
	templateEureka = strings.Replace(templateEureka, "${port}", portApplication, -1)
	templateEureka = strings.Replace(templateEureka, "${app-name}", ApplicationName, -1)
	templateEureka = strings.Replace(templateEureka, "${statusEureka}", status, -1)

	return templateEureka
}

func RegisterEureka(status string, portApplication string) {
	templateEureka := formatTemplate(status, portApplication)
	// RegisterEureka.
	registerAction := HttpAction{ // Build a HttpAction struct
		Url:         fmt.Sprintf("%s%s", ServerEureka, ApplicationName), // Note hard-coded path to Eureka...
		Method:      "POST",
		ContentType: "application/json",
		Body:        templateEureka,
	}
	var result bool
	for {
		result = DoHttpRequest(registerAction) // Execute the HTTP request. result == true if req went OK
		fmt.Println("result", result)
		if result {
			// Update
			break // Success, end registration loop
		} else {
			time.Sleep(time.Second * 1) // Registration failed (usually, Eureka isn't up yet),
		} // retry in 5 seconds.
	}
}
func SetEurekaStatus(status string, portApplication string) {
	templateEureka := formatTemplate(status, portApplication)
	// RegisterEureka.
	upAction := HttpAction{ // Build a HttpAction struct
		Url:         fmt.Sprintf("%s%s", ServerEureka, ApplicationName), // Note hard-coded path to Eureka...
		Method:      "POST",
		ContentType: "application/json",
		Body:        templateEureka,
	}
	result := DoHttpRequest(upAction) // Execute the HTTP request. result == true if req went OK
	if !result {
		fmt.Println("Error exception")
	}
}

func CheckServiceLive(port string, isStartChan chan bool) {
	isStart := <- isStartChan
	for {
		time.Sleep(time.Second * 5)
		if isStart && !checkInstanceService(ApplicationName) {
			RegisterEureka("UP", port)
		}
	}

}

func StartHeartbeat() {
	for {
		time.Sleep(time.Second * 1)
		heartbeat()
	}
}
func GetInstanceService(serviceName string, apiName string) string {
	urlService := fmt.Sprintf("%s%s", ServerEureka, serviceName)
	resp, err := http.Get(urlService)
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var application object.ApplicationsXml
	err1 := xml.Unmarshal(body, &application.Application)
	if err1 != nil {
		fmt.Println(err1.Error())
	}
	// Load balancer
	listInstance := application.Application.Instance
	if len(listInstance) <= 0 {
		return ""
	}
	var rws RandomWeightedSelector
	for index, instance := range listInstance {
		if instance.Status == "UP" {
			urlReal := fmt.Sprintf("%s%s:%s/%s", Prefix, instance.IpAddr, instance.Port, apiName)
			rws.AddEndpoint(Endpoint{Weight: index + 1, URL: urlReal})
		}
	}
	return rws.Select().URL
}

func heartbeat() {
	heartbeatAction := HttpAction{
		Url:    fmt.Sprintf("%s%s/%s:%s", ServerEureka, ApplicationName, GetLocalIP(), ApplicationName),
		Method: "PUT",
	}
	DoHttpRequest(heartbeatAction)
}
func checkInstanceService(serviceName string) bool {
	urlService := fmt.Sprintf("%s%s", ServerEureka, serviceName)
	resp, err := http.Get(urlService)
	if err != nil {
		//log.Fatalln(err)
		return false
	}
	if resp.StatusCode != http.StatusOK {
		return false
	}
	return true
}

func Deregister() {
	fmt.Println("Trying to deregister application...")
	deregisterAction := HttpAction{
		Url:    fmt.Sprintf("%s%s/%s:%s", ServerEureka, ApplicationName, GetLocalIP(), ApplicationName),
		Method: "DELETE",
	}
	DoHttpRequest(deregisterAction)
	fmt.Println("Deregistered application, exiting. Check Eureka...")
}
func HandleSigterm() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		//eureka.Deregister()
		SetEurekaStatus("DOWN", "8080")
		os.Exit(1)
	}()
}
