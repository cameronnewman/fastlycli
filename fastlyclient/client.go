package fastlyclient

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

//FastlyAPIEndPoint Default API hostname
const (
	FastlyAPIEndPoint = "https://api.fastly.com"
	FastlyPurgeAll    = "*"
)

//Client struct for client connection
type Client struct {
}

//FastlyAPIKey api key
var FastlyAPIKey string

//NewFastlyClient constructor
func NewFastlyClient() *Client {
	fastlyClient := &Client{}
	return fastlyClient
}

//GetServiceDetails Get Fastly service by friendly name
func (c *Client) GetServiceDetails(serviceName string) {

	if c.checkAPIKey() {

		var result SearchResultModel
		result = (*c).lookupServiceByName(serviceName)

		if result.ID != "" {

			//GET /service/search?name={serviceName}
			req, err := http.NewRequest("GET", FastlyAPIEndPoint+"/service/"+result.ID+"/details", nil)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Fastly-Key", FastlyAPIKey)

			client := &http.Client{}
			response, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer response.Body.Close()
			body, _ := ioutil.ReadAll(response.Body)

			if response.StatusCode == http.StatusOK {

				var service ServiceModel
				err := json.Unmarshal(body, &service)

				if err != nil {
					panic(err)
				}

				result2, err := json.MarshalIndent(service, "", "\t")
				if err != nil {
					panic(err)
				}

				println(string(result2))

			}
		}
	}
}

//GetServiceDomains Service public cnames
func (c *Client) GetServiceDomains(serviceName string) {

	if c.checkAPIKey() {

		var result SearchResultModel
		result = c.lookupServiceByName(serviceName)

		if result.ID != "" {

			//println("Getting " + serviceName + " service domains")

			//GET /service/search?name={serviceName}
			req, err := http.NewRequest("GET", FastlyAPIEndPoint+"/service/"+result.ID+"/details", nil)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Fastly-Key", FastlyAPIKey)

			client := &http.Client{}
			response, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer response.Body.Close()
			body, _ := ioutil.ReadAll(response.Body)

			if response.StatusCode == http.StatusOK {

				var service ServiceModel
				err := json.Unmarshal(body, &service)

				if err != nil {
					panic(err)
				}
			}
		}
	}
}

//GetServiceBackends Get all the Service backends
func (c *Client) GetServiceBackends(serviceName string) {

	if c.checkAPIKey() {

		var result SearchResultModel
		result = c.lookupServiceByName(serviceName)

		if result.ID != "" {

			//println("Getting " + serviceName + " service domains")

			//GET /service/search?name={serviceName}
			req, err := http.NewRequest("GET", FastlyAPIEndPoint+"/service/"+result.ID+"/details", nil)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Fastly-Key", FastlyAPIKey)

			client := &http.Client{}
			response, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer response.Body.Close()

			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				panic(err)
			}

			if response.StatusCode == http.StatusOK {

				var service ServiceModel
				err := json.Unmarshal(body, &service)
				if err != nil {
					panic(err)
				}
			}
		}
	}
}

//PurgeObjects purge object
func (c *Client) PurgeObjects(serviceName string, objects string) {

	if c.checkAPIKey() && objects != "" {

		if objects != "" {
			var service SearchResultModel
			service = c.lookupServiceByName(serviceName)

			if service.ID != "" {

				if objects == FastlyPurgeAll {
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

					if response.StatusCode == http.StatusOK {
						println("Service " + serviceName + " successfully purged")
					} else {
						println("Service " + serviceName + " failed to purge cached objects")
					}

				} else {

					//POST /service/ekjhsdfkjhsdfouejk/purge??
					req, err := http.NewRequest("PURGE", FastlyAPIEndPoint+"/service/"+service.ID+"/"+objects, nil)
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

		} else {
			println("No object or wildcard set")
		}
	}
}

func (c *Client) checkAPIKey() bool {
	if os.Getenv("FASTLYAPIKEY") == "" {
		println("Fastly API key not set. Please export $FASTLYAPIKEY=x")
		return false
	}

	FastlyAPIKey = os.Getenv("FASTLYAPIKEY")
	return true
}

func (c *Client) lookupServiceByName(serviceName string) SearchResultModel {

	var service SearchResultModel

	if c.checkAPIKey() {

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
		if err != nil {
			panic(err)
		}

		if response.Status == "200 OK" {

			err := json.Unmarshal(body, &service)
			//println("Service " + serviceName + " found. ID=" + service.ID)

			if err != nil {
				panic(err)
			}
		} else {
			println("No service called " + serviceName + " exists")
		}
	}
	return service
}
