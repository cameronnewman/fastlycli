package fastlyclient

import "errors"

var (
	ErrorNoAPIKeySet             = errors.New("Fastly API key not set. Please export $FASTLYAPIKEY=x")
	ErrorNoServiceWithNameExists = errors.New("No service called with name exists")
)
