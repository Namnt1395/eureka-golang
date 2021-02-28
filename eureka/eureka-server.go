package eureka

import (
	"eureka-golang/util"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

const (
	NameApplication = "call-api-demo"
	ServerEureka    = "http://157.230.53.38:8761/eureka/apps/"
	//InstanceId      = "Demo1"
)

//var instanceId string

//func RegisterEureka() {
//	instanceId = util.GetUUID()
//
//	dir, _ := os.Getwd()
//	data, err1 := ioutil.ReadFile(dir + "/templates/regtpl.json")
//	if err1 != nil {
//		fmt.Println(err1.Error())
//	}
//	fmt.Println("data...", data)
//	tpl := string(data)
//	tpl = strings.Replace(tpl, "${ipAddress}", util.GetLocalIP(), -1) // Replace some placeholders
//	tpl = strings.Replace(tpl, "${port}", "8080", -1)
//	tpl = strings.Replace(tpl, "${instanceId}", instanceId, -1)
//
//	// RegisterEureka.
//	registerAction := HttpAction{ // Build a HttpAction struct
//		Url:         "http://192.168.100.123:8761/eureka/apps/api", // Note hard-coded path to Eureka...
//		Method:      "POST",
//		ContentType: "application/json",
//		Body:        tpl,
//	}
//	var result bool
//	for {
//		result = DoHttpRequest(registerAction) // Execute the HTTP request. result == true if req went OK
//		fmt.Println("resutl....", result)
//		if result {
//			// Update
//			break // Success, end registration loop
//		} else {
//			time.Sleep(time.Second * 1) // Registration failed (usually, Eureka isn't up yet),
//		} // retry in 5 seconds.
//	}
//}

func getTemplateEureka() string {
	dir, _ := os.Getwd()
	data, err1 := ioutil.ReadFile(dir + "/templates/regtpl.json")
	if err1 != nil {
		fmt.Println(err1.Error())
	}
	return string(data)
}

func formatTemplate(status string, portApplication string) string {
	templateEureka := getTemplateEureka()
	templateEureka = strings.Replace(templateEureka, "${ipAddress}", util.GetLocalIP(), -1) // Replace some placeholders
	templateEureka = strings.Replace(templateEureka, "${port}", portApplication, -1)
	templateEureka = strings.Replace(templateEureka, "${app-name}", NameApplication, -1)
	templateEureka = strings.Replace(templateEureka, "${statusEureka}", status, -1)

	return templateEureka
}

func RegisterEureka(status string, portApplication string) {
	templateEureka := formatTemplate(status, portApplication)
	// RegisterEureka.
	registerAction := HttpAction{ // Build a HttpAction struct
		Url:         fmt.Sprintf("%s%s", ServerEureka, NameApplication), // Note hard-coded path to Eureka...
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
		Url:         fmt.Sprintf("%s%s", ServerEureka, NameApplication), // Note hard-coded path to Eureka...
		Method:      "POST",
		ContentType: "application/json",
		Body:        templateEureka,
	}
	result := DoHttpRequest(upAction) // Execute the HTTP request. result == true if req went OK
	if !result {
		fmt.Println("Error exception")
	}
}

func StartHeartbeat() {
	for {
		time.Sleep(time.Second * 1)
		heartbeat()
	}
}

func GetInstanceService(serviceName string)  {
	
}

func heartbeat() {
	heartbeatAction := HttpAction{
		//Url:    "http://192.168.100.123:8761/eureka/apps/api/" + util.GetLocalIP() + ":api:" + InstanceId,
		Url:    fmt.Sprintf("%s%s/%s:%s", ServerEureka, NameApplication, util.GetLocalIP(), NameApplication),
		Method: "PUT",
	}
	DoHttpRequest(heartbeatAction)
}

func Deregister() {
	fmt.Println("Trying to deregister application...")
	deregisterAction := HttpAction{
		Url:    fmt.Sprintf("%s%s/%s:%s", ServerEureka, NameApplication, util.GetLocalIP(), NameApplication),
		Method: "DELETE",
	}
	DoHttpRequest(deregisterAction)
	fmt.Println("Deregistered application, exiting. Check Eureka...")
}
