package main

import (
	"eureka-golang/eureka"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	linkApi := eureka.GetInstanceService("api-demo", "api")
	if len(linkApi) > 0 {
		fmt.Println("linkApi...", linkApi)
		resp, err := http.Get(linkApi)
		if err != nil {
			log.Fatalln(err)
		}
		//We Read the response body on the line below.
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("data...", string(body))
	}

}
