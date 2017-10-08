package fastlyclient

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"
)

const (
	httpMaxIdleConnections   int    = 30
	httpRequestTimeout       int    = 60
	fastlyAPIEndPoint        string = "https://api.fastly.com"
	fastlyAPIEnvironmentName string = "FASTLYAPIKEY"
)

//Client struct for fastly client connection
type Client struct {
	httpClient *http.Client
	apiKey     string
}

//NewFastlyClient constructor
func NewFastlyClient() *Client {
	client := &Client{}

	key, err := client.getAPIKey()
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	client.apiKey = key
	client.httpClient = client.initHTTPClient()
	return client
}

//GetServiceDetails get service details via friendly name
func (c *Client) GetServiceDetails(serviceName string) {

	var result SearchResultModel
	result, err := c.lookupServiceByName(serviceName)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	//GET /service/search?name={serviceName}
	req, err := http.NewRequest("GET", fastlyAPIEndPoint+"/service/"+result.ID+"/details", nil)
	req.Header.Set(HeaderContentType, MIMEApplicationJSON)
	req.Header.Set(HeaderFastlyKey, c.apiKey)

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
		return
	}
	println(ErrorFailedToRetrieveServiceDetails.Error())
}

//GetServiceDomains get public cnames for service
func (c *Client) GetServiceDomains(serviceName string) {

	var result SearchResultModel
	result, err := c.lookupServiceByName(serviceName)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	//GET /service/search?name={serviceName}
	req, err := http.NewRequest(http.MethodGet, fastlyAPIEndPoint+"/service/"+result.ID+"/details", nil)
	req.Header.Set(HeaderContentType, MIMEApplicationJSON)
	req.Header.Set(HeaderFastlyKey, c.apiKey)

	response, err := c.httpClient.Do(req)
	if err != nil {
		println(err.Error())
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	if response.StatusCode == http.StatusOK {

		var service ServiceModel
		err := json.Unmarshal(body, &service)
		if err != nil {
			println(err.Error())
		}

		result, err := json.MarshalIndent(service.ActiveVersion.Domains, "", "\t")
		if err != nil {
			panic(err)
		}
		println(string(result))
		return
	}
	println(ErrorFailedToRetrieveServiceDetails.Error())
}

//GetServiceBackends get all backends for a service
func (c *Client) GetServiceBackends(serviceName string) {

	var result SearchResultModel
	result, err := c.lookupServiceByName(serviceName)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	//GET /service/search?name={serviceName}
	req, err := http.NewRequest(http.MethodGet, fastlyAPIEndPoint+"/service/"+result.ID+"/details", nil)
	req.Header.Set(HeaderContentType, MIMEApplicationJSON)
	req.Header.Set(HeaderFastlyKey, c.apiKey)

	response, err := c.httpClient.Do(req)
	if err != nil {
		println(err.Error())
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		println(err.Error())
	}

	if response.StatusCode == http.StatusOK {
		var service ServiceModel
		err := json.Unmarshal(body, &service)
		if err != nil {
			println(err.Error())
		}

		result, err := json.MarshalIndent(service.ActiveVersion.Backends, "", "\t")
		if err != nil {
			panic(err)
		}
		println(string(result))
		return
	}
	println(ErrorFailedToRetrieveServiceDetails.Error())
}

//PurgeObject an object from the service
func (c *Client) PurgeObject(serviceName string, object string) {

	var result SearchResultModel
	result, err := c.lookupServiceByName(serviceName)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	if c.isStringEmpty(object) {
		println(ErrorNoObjectSetForPurge.Error())
		os.Exit(1)
	}

	_, err = url.ParseRequestURI(object)
	if err != nil {
		println(ErrorObjectIsNotVaildateURI.Error())
		os.Exit(1)
	}

	//POST /service/ekjhsdfkjhsdfouejk/purge??
	req, err := http.NewRequest(HTTPMethodPurge, fastlyAPIEndPoint+"/service/"+result.ID+"/"+object, nil)
	req.Header.Set(HeaderContentType, MIMEApplicationJSON)
	req.Header.Set(HeaderFastlyKey, c.apiKey)

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		println(err.Error())
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		println(ResponseSucessfullyPurgedObject)
		return
	}
	println(ErrorFailedToPurgeObject.Error())
}

//PurgeAllObjects for service
func (c *Client) PurgeAllObjects(serviceName string) {

	var result SearchResultModel
	result, err := c.lookupServiceByName(serviceName)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	//POST /service/ekjhsdfkjhsdfouejk/purge_all
	req, err := http.NewRequest(http.MethodPost, fastlyAPIEndPoint+"/service/"+result.ID+"/purge_all", nil)
	req.Header.Set(HeaderContentType, MIMEApplicationJSON)
	req.Header.Set(HeaderFastlyKey, c.apiKey)

	response, err := c.httpClient.Do(req)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		println(ResponseSucessfullyPurgedAllObjectsFromService)
		return
	}
	println(ErrorFailedToPurgeService.Error())
}

func (c *Client) lookupServiceByName(serviceName string) (SearchResultModel, error) {

	var service SearchResultModel

	//GET /service/search?name={serviceName}
	req, err := http.NewRequest(http.MethodGet, fastlyAPIEndPoint+"/service/search?name="+serviceName, nil)
	req.Header.Set(HeaderContentType, MIMEApplicationJSON)
	req.Header.Set(HeaderFastlyKey, c.apiKey)

	response, err := c.httpClient.Do(req)
	if err != nil {
		return service, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return service, err
	}

	if response.StatusCode == http.StatusForbidden {
		return service, ErrorAPIKeyDoesntHaveAccess
	}

	if response.StatusCode == http.StatusOK {
		err := json.Unmarshal(body, &service)
		if err != nil {
			return service, err
		}
		return service, nil
	}

	return service, ErrorNoServiceWithNameExists
}

func (c *Client) isStringEmpty(s string) bool {
	return len(s) == 0
}

func (c *Client) getAPIKey() (string, error) {
	if os.Getenv(fastlyAPIEnvironmentName) == "" {
		return "", ErrorNoAPIKeySet
	}
	return os.Getenv(fastlyAPIEnvironmentName), nil
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
