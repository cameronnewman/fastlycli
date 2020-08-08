package fastly

import (
	"io"
	"io/ioutil"
	"net/http"
)

const (
	//HeaderContentType for content type
	HeaderContentType string = "Content-Type"

	//HeaderFastlyKey Fastly API Key Header
	HeaderFastlyKey string = "Fastly-Key"

	//MIMEApplicationJSON content type for JSON
	MIMEApplicationJSON string = "application/json"

	//HTTPMethodPurge Fastly extension http method for purging the CDN
	HTTPMethodPurge string = "PURGE"
)

func (f *Fastly) get(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set(HeaderContentType, MIMEApplicationJSON)
	req.Header.Set(HeaderFastlyKey, f.apiKey)
	if err != nil {
		return nil, err
	}

	res, err := f.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusForbidden {
		return nil, ErrAPIKeyDoesntHaveAccess
	}

	if res.StatusCode != http.StatusOK {
		return nil, errHTTPRequestFailed
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (f *Fastly) post(url string, payload io.Reader) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, url, payload)
	req.Header.Set(HeaderContentType, MIMEApplicationJSON)
	req.Header.Set(HeaderFastlyKey, f.apiKey)
	if err != nil {
		return nil, err
	}

	res, err := f.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusForbidden {
		return nil, ErrAPIKeyDoesntHaveAccess
	}

	if res.StatusCode != http.StatusOK {
		return nil, errHTTPRequestFailed
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (f *Fastly) purge(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, url, nil)
	req.Header.Set(HeaderContentType, MIMEApplicationJSON)
	req.Header.Set(HeaderFastlyKey, f.apiKey)
	if err != nil {
		return nil, err
	}

	res, err := f.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusForbidden {
		return nil, ErrAPIKeyDoesntHaveAccess
	}

	if res.StatusCode != http.StatusOK {
		return nil, errHTTPRequestFailed
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
