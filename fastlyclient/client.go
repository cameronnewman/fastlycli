package fastlyclient

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

//FastlyAPIEndPoint Default API hostname
const (
	httpMaxIdleConnections   int    = 30
	httpRequestTimeout       int    = 60
	FastlyAPIEndPoint        string = "https://api.fastly.com"
	FastlyPurgeAll           string = "*"
	FastlyAPIEnvironmentName string = "FASTLYAPIKEY"
)

//Client struct for client connection
type Client struct {
	httpClient *http.Client
}

//FastlyAPIKey api key
var FastlyAPIKey string

//NewFastlyClient constructor
func NewFastlyClient() *Client {
	client := &Client{}
	client.httpClient = client.initHTTPClient()
	return client
}

//GetServiceDetails Get Fastly service by friendly name
func (c *Client) GetServiceDetails(serviceName string) {

	if c.checkAPIKey() {

		var result SearchResultModel
		result = (*c).lookupServiceByName(serviceName)

		if result.ID != "" {

			//GET /service/search?name={serviceName}
			req, err := http.NewRequest("GET", FastlyAPIEndPoint+"/service/"+result.ID+"/details", nil)
			req.Header.Set(HeaderContentType, MIMEApplicationJSON)
			req.Header.Set(HeaderFastlyKey, FastlyAPIKey)

			response, err := c.httpClient.Do(req)
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

				result, err := json.MarshalIndent(service, "", "\t")
				if err != nil {
					panic(err)
				}

				println(string(result))
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

			//GET /service/search?name={serviceName}
			req, err := http.NewRequest(http.MethodGet, FastlyAPIEndPoint+"/service/"+result.ID+"/details", nil)
			req.Header.Set(HeaderContentType, MIMEApplicationJSON)
			req.Header.Set(HeaderFastlyKey, FastlyAPIKey)

			response, err := c.httpClient.Do(req)
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
			//GET /service/search?name={serviceName}
			req, err := http.NewRequest(http.MethodGet, FastlyAPIEndPoint+"/service/"+result.ID+"/details", nil)
			req.Header.Set(HeaderContentType, MIMEApplicationJSON)
			req.Header.Set(HeaderFastlyKey, FastlyAPIKey)

			response, err := c.httpClient.Do(req)
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
					req, err := http.NewRequest(http.MethodPost, FastlyAPIEndPoint+"/service/"+service.ID+"/purge_all", nil)
					req.Header.Set(HeaderContentType, MIMEApplicationJSON)
					req.Header.Set(HeaderFastlyKey, FastlyAPIKey)

					response, err := c.httpClient.Do(req)
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
					req, err := http.NewRequest(HTTPMethodPurge, FastlyAPIEndPoint+"/service/"+service.ID+"/"+objects, nil)
					req.Header.Set(HeaderContentType, MIMEApplicationJSON)
					req.Header.Set(HeaderFastlyKey, FastlyAPIKey)

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
				}
			}

		} else {
			println("No object or wildcard set")
		}
	}
}

func (c *Client) checkAPIKey() bool {
	if os.Getenv(FastlyAPIEnvironmentName) == "" {
		println("Fastly API key not set. Please export $FASTLYAPIKEY=x")
		return false
	}

	FastlyAPIKey = os.Getenv(FastlyAPIEnvironmentName)
	return true
}

func (c *Client) lookupServiceByName(serviceName string) SearchResultModel {

	var service SearchResultModel

	if c.checkAPIKey() {

		//GET /service/search?name={serviceName}
		req, err := http.NewRequest(http.MethodGet, FastlyAPIEndPoint+"/service/search?name="+serviceName, nil)
		req.Header.Set(HeaderContentType, MIMEApplicationJSON)
		req.Header.Set(HeaderFastlyKey, FastlyAPIKey)

		response, err := c.httpClient.Do(req)
		if err != nil {
			panic(err)
		}
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}

		if response.StatusCode == http.StatusOK {
			err := json.Unmarshal(body, &service)
			if err != nil {
				panic(err)
			}
		} else {
			println("No service called " + serviceName + " exists")
		}
	}
	return service
}

func (*Client) initHTTPClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: httpMaxIdleConnections,
		},
		Timeout: time.Duration(httpRequestTimeout) * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
}
