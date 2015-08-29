package fastlyservice

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const FastlyAPIEndPoint = "https://api.fastly.com"

type FastlyClient struct {
}

var FastlyAPIKey string

func NewFastlyService() *FastlyClient {
	fastlyClient := &FastlyClient{}
	return fastlyClient
}

func (fs *FastlyClient) PurgeObjects(serviceName string, objects string) {

	if (*fs).CheckAPIKey() {

		var service FastlyServiceModel
		service = (*fs).lookupServiceDetails(serviceName)

		if service.ID != "" {
			println("Purging " + serviceName + " service")

			//POST /service/ekjhsdfkjhsdfouejk/purge_all
			req, err := http.NewRequest("POST", FastlyAPIEndPoint+"/service/"+service.ID+"/purge_all", nil)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Fastly-Key", FastlyAPIKey)

			client := &http.Client{}
			response, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer response.Body.Close()

			if response.Status == "200 OK" {
				println("Service " + serviceName + " successfully purged")
			} else {
				println("Service " + serviceName + " failed to purge cached objects")
			}
		}
	}
}

func (fs *FastlyClient) ReturnServiceDetails(serviceName string) {

	if (*fs).CheckAPIKey() {

		var service FastlyServiceModel
		service = (*fs).lookupServiceDetails(serviceName)

		if service.ID != "" {
			service.Versions = nil
			result, err := json.Marshal(service)
			if err != nil {
				panic(err)
			}

			fmt.Println(string(result))
		}
	}
}

func (fs *FastlyClient) lookupServiceDetails(serviceName string) FastlyServiceModel {

	var service FastlyServiceModel

	if (*fs).CheckAPIKey() {

		println("Getting " + serviceName + " service details")

		//GET /service/search?name={serviceName}
		req, err := http.NewRequest("GET", FastlyAPIEndPoint+"/service/search?name="+serviceName, nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Fastly-Key", FastlyAPIKey)

		client := &http.Client{}
		response, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)

		if response.Status == "200 OK" {

			err := json.Unmarshal(body, &service)
			println("Service " + serviceName + " found. ID=" + service.ID)

			if err != nil {
				panic(err)
			}
		} else {
			println("No service called " + serviceName + " exists")
		}
	}
	return service
}

func (fs *FastlyClient) CheckAPIKey() bool {
	if os.Getenv("FASTLYAPIKEY") == "" {
		println("Fastly API key not set. Please export $FASTLYAPIKEY=x")
		return false
	} else {
		FastlyAPIKey = os.Getenv("FASTLYAPIKEY")
		return true
	}
}
