package fastly

import (
	"encoding/json"
)

//GetService gets all details for a fastly service via it's friendly name
func (f *Fastly) GetService(serviceName string) (string, error) {
	service, err := f.getServiceByName(serviceName)
	if err != nil {
		return "", err
	}

	result, err := json.MarshalIndent(service, "", "\t")
	if err != nil {
		return "", err
	}
	return string(result), nil
}

//GetServiceDomains get public cnames for service
func (f *Fastly) GetServiceDomains(serviceName string) (string, error) {
	service, err := f.getServiceByName(serviceName)
	if err != nil {
		return "", err
	}

	result, err := json.MarshalIndent(service.ActiveVersion.Domains, "", "\t")
	if err != nil {
		return "", err
	}
	return string(result), nil
}

//GetServiceBackends get all backends for a service
func (f *Fastly) GetServiceBackends(serviceName string) (string, error) {
	service, err := f.getServiceByName(serviceName)
	if err != nil {
		return "", err
	}

	result, err := json.MarshalIndent(service.ActiveVersion.Backends, "", "\t")
	if err != nil {
		return "", err
	}
	return string(result), nil
}

func (f *Fastly) getServiceByName(serviceName string) (Service, error) {
	var serviceSearch ServiceSearchResult
	var service Service

	//GET /service/search?name={serviceName}
	body, err := f.get(fastlyAPIEndPoint + "/service/search?name=" + serviceName)
	if err != nil {
		return service, err
	}

	err = json.Unmarshal(body, &serviceSearch)
	if err != nil {
		return service, err
	}

	//GET /service/{serviceID}/details
	body, err = f.get(fastlyAPIEndPoint + "/service/" + serviceSearch.ID + "/details")
	if err != nil {
		return service, err
	}

	err = json.Unmarshal(body, &service)
	if err != nil {
		return service, err
	}

	return service, nil
}

//ServiceSearchResult is a result from search by service name
type ServiceSearchResult struct {
	Comment    string `json:"comment"`
	CustomerID string `json:"customer_id"`
	ID         string `json:"id"`
	Name       string `json:"name"`
	Versions   []struct {
		Active          bool          `json:"active"`
		Backend         []interface{} `json:"backend"`
		CacheSettings   []interface{} `json:"cache_settings"`
		Comment         string        `json:"comment"`
		Condition       []interface{} `json:"condition"`
		Created         string        `json:"created"`
		Deleted         string        `json:"deleted"`
		Deployed        bool          `json:"deployed"`
		Director        []interface{} `json:"director"`
		Domain          []interface{} `json:"domain"`
		Gzip            []interface{} `json:"gzip"`
		Header          []interface{} `json:"header"`
		Healthcheck     []interface{} `json:"healthcheck"`
		Locked          bool          `json:"locked"`
		LoggingSyslog   []interface{} `json:"logging_syslog"`
		Number          int           `json:"number"`
		RequestSettings []interface{} `json:"request_settings"`
		ResponseObject  []interface{} `json:"response_object"`
		ServiceID       string        `json:"service_id"`
		Settings        struct {
			GeneralDefaultHost string `json:"general.default_host"`
			GeneralDefaultPci  int    `json:"general.default_pci"`
			GeneralDefaultTTL  int    `json:"general.default_ttl"`
		} `json:"settings"`
		Staging   bool          `json:"staging"`
		Testing   bool          `json:"testing"`
		Updated   string        `json:"updated"`
		Vcl       []interface{} `json:"vcl"`
		Waf       []interface{} `json:"waf"`
		Wordpress []interface{} `json:"wordpress"`
	} `json:"versions"`
}

//Service is a fastly service from details lookup
type Service struct {
	ActiveVersion struct {
		Active   bool `json:"active"`
		Backends []struct {
			Address             string      `json:"address"`
			AutoLoadbalance     bool        `json:"auto_loadbalance"`
			BetweenBytesTimeout int         `json:"between_bytes_timeout"`
			ClientCert          interface{} `json:"client_cert"`
			Comment             string      `json:"comment"`
			ConnectTimeout      int         `json:"connect_timeout"`
			ErrorThreshold      int         `json:"error_threshold"`
			FirstByteTimeout    int         `json:"first_byte_timeout"`
			Healthcheck         string      `json:"healthcheck"`
			Hostname            string      `json:"hostname"`
			Ipv4                interface{} `json:"ipv4"`
			Ipv6                interface{} `json:"ipv6"`
			MaxConn             int         `json:"max_conn"`
			Name                string      `json:"name"`
			Port                int         `json:"port"`
			RequestCondition    string      `json:"request_condition"`
			Shield              string      `json:"shield"`
			SslCaCert           interface{} `json:"ssl_ca_cert"`
			SslClientCert       interface{} `json:"ssl_client_cert"`
			SslClientKey        interface{} `json:"ssl_client_key"`
			SslHostname         string      `json:"ssl_hostname"`
			UseSsl              bool        `json:"use_ssl"`
			Weight              int         `json:"weight"`
		} `json:"backends"`
		CacheSettings []interface{} `json:"cache_settings"`
		Comment       string        `json:"comment"`
		Conditions    []interface{} `json:"conditions"`
		Deployed      interface{}   `json:"deployed"`
		Directors     []interface{} `json:"directors"`
		Domains       []struct {
			Comment string `json:"comment"`
			Name    string `json:"name"`
		} `json:"domains"`
		Gzips []struct {
			CacheCondition string `json:"cache_condition"`
			ContentTypes   string `json:"content_types"`
			Extensions     string `json:"extensions"`
			Name           string `json:"name"`
		} `json:"gzips"`
		Headers          []interface{} `json:"headers"`
		Healthchecks     []interface{} `json:"healthchecks"`
		InheritServiceID interface{}   `json:"inherit_service_id"`
		Locked           bool          `json:"locked"`
		Matches          []interface{} `json:"matches"`
		Number           int           `json:"number"`
		Origins          []interface{} `json:"origins"`
		RequestSettings  []struct {
			Action           interface{} `json:"action"`
			BypassBusyWait   interface{} `json:"bypass_busy_wait"`
			DefaultHost      string      `json:"default_host"`
			ForceMiss        interface{} `json:"force_miss"`
			ForceSsl         string      `json:"force_ssl"`
			GeoHeaders       interface{} `json:"geo_headers"`
			HashKeys         string      `json:"hash_keys"`
			MaxStaleAge      string      `json:"max_stale_age"`
			Name             string      `json:"name"`
			RequestCondition string      `json:"request_condition"`
			TimerSupport     string      `json:"timer_support"`
			Xff              string      `json:"xff"`
		} `json:"request_settings"`
		ResponseObjects []interface{} `json:"response_objects"`
		ServiceID       string        `json:"service_id"`
		Settings        struct {
			GeneralDefaultHost string `json:"general.default_host"`
			GeneralDefaultPci  int    `json:"general.default_pci"`
			GeneralDefaultTTL  int    `json:"general.default_ttl"`
		} `json:"settings"`
		Staging   interface{}   `json:"staging"`
		Testing   interface{}   `json:"testing"`
		Vcls      []interface{} `json:"vcls"`
		Wordpress []interface{} `json:"wordpress"`
	} `json:"active_version"`
	Comment    string `json:"comment"`
	CustomerID string `json:"customer_id"`
	ID         string `json:"id"`
	Name       string `json:"name"`
	Version    struct {
		Active   bool `json:"active"`
		Backends []struct {
			Address             string      `json:"address"`
			AutoLoadbalance     bool        `json:"auto_loadbalance"`
			BetweenBytesTimeout int         `json:"between_bytes_timeout"`
			ClientCert          interface{} `json:"client_cert"`
			Comment             string      `json:"comment"`
			ConnectTimeout      int         `json:"connect_timeout"`
			ErrorThreshold      int         `json:"error_threshold"`
			FirstByteTimeout    int         `json:"first_byte_timeout"`
			Healthcheck         string      `json:"healthcheck"`
			Hostname            string      `json:"hostname"`
			Ipv4                interface{} `json:"ipv4"`
			Ipv6                interface{} `json:"ipv6"`
			MaxConn             int         `json:"max_conn"`
			Name                string      `json:"name"`
			Port                int         `json:"port"`
			RequestCondition    string      `json:"request_condition"`
			Shield              string      `json:"shield"`
			SslCaCert           interface{} `json:"ssl_ca_cert"`
			SslClientCert       interface{} `json:"ssl_client_cert"`
			SslClientKey        interface{} `json:"ssl_client_key"`
			SslHostname         string      `json:"ssl_hostname"`
			UseSsl              bool        `json:"use_ssl"`
			Weight              int         `json:"weight"`
		} `json:"backends"`
		CacheSettings []interface{} `json:"cache_settings"`
		Comment       string        `json:"comment"`
		Conditions    []interface{} `json:"conditions"`
		Deployed      interface{}   `json:"deployed"`
		Directors     []interface{} `json:"directors"`
		Domains       []struct {
			Comment string `json:"comment"`
			Name    string `json:"name"`
		} `json:"domains"`
		Gzips []struct {
			CacheCondition string `json:"cache_condition"`
			ContentTypes   string `json:"content_types"`
			Extensions     string `json:"extensions"`
			Name           string `json:"name"`
		} `json:"gzips"`
		Headers          []interface{} `json:"headers"`
		Healthchecks     []interface{} `json:"healthchecks"`
		InheritServiceID interface{}   `json:"inherit_service_id"`
		Locked           bool          `json:"locked"`
		Matches          []interface{} `json:"matches"`
		Number           int           `json:"number"`
		Origins          []interface{} `json:"origins"`
		RequestSettings  []struct {
			Action           interface{} `json:"action"`
			BypassBusyWait   interface{} `json:"bypass_busy_wait"`
			DefaultHost      string      `json:"default_host"`
			ForceMiss        interface{} `json:"force_miss"`
			ForceSsl         string      `json:"force_ssl"`
			GeoHeaders       interface{} `json:"geo_headers"`
			HashKeys         string      `json:"hash_keys"`
			MaxStaleAge      string      `json:"max_stale_age"`
			Name             string      `json:"name"`
			RequestCondition string      `json:"request_condition"`
			TimerSupport     string      `json:"timer_support"`
			Xff              string      `json:"xff"`
		} `json:"request_settings"`
		ResponseObjects []interface{} `json:"response_objects"`
		ServiceID       string        `json:"service_id"`
		Settings        struct {
			GeneralDefaultHost string `json:"general.default_host"`
			GeneralDefaultPci  int    `json:"general.default_pci"`
			GeneralDefaultTTL  int    `json:"general.default_ttl"`
		} `json:"settings"`
		Staging   interface{}   `json:"staging"`
		Testing   interface{}   `json:"testing"`
		Vcls      []interface{} `json:"vcls"`
		Wordpress []interface{} `json:"wordpress"`
	} `json:"version"`
	Versions []struct {
		Active           bool        `json:"active"`
		Comment          string      `json:"comment"`
		CreatedAt        string      `json:"created_at"`
		DeletedAt        interface{} `json:"deleted_at"`
		Deployed         interface{} `json:"deployed"`
		InheritServiceID interface{} `json:"inherit_service_id"`
		Locked           bool        `json:"locked"`
		Number           int         `json:"number"`
		ServiceID        string      `json:"service_id"`
		Staging          interface{} `json:"staging"`
		Testing          interface{} `json:"testing"`
		UpdatedAt        string      `json:"updated_at"`
	} `json:"versions"`
}
