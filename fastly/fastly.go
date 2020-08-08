package fastly

import (
	"net/http"
	"os"
	"time"
)

const (
	httpMaxIdleConnections   int    = 30
	httpRequestTimeout       int    = 60
	fastlyAPIEndPoint        string = "https://api.fastly.com"
	fastlyAPIEnvironmentName string = "FASTLYAPIKEY"
)

//Fastly struct for fastly client connection
type Fastly struct {
	httpClient *http.Client
	apiKey     string
}

//New constructor
func New() *Fastly {
	client := &Fastly{}

	key, err := client.getAPIKey()
	if err != nil {
		panic(err)
	}

	client.apiKey = key
	client.httpClient = client.initHTTPClient()
	return client
}

func (f *Fastly) isStringEmpty(s string) bool {
	return len(s) == 0
}

func (f *Fastly) getAPIKey() (string, error) {
	if os.Getenv(fastlyAPIEnvironmentName) == "" {
		return "", ErrNoAPIKeySet
	}
	return os.Getenv(fastlyAPIEnvironmentName), nil
}

func (*Fastly) initHTTPClient() *http.Client {
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
