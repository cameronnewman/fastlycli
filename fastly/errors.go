package fastly

import (
	"errors"
)

var (
	//ErrNoAPIKeySet is returned when no API Key set
	ErrNoAPIKeySet = errors.New("Fastly API key not set. Please export $FASTLYAPIKEY=x")

	//ErrAPIKeyDoesntHaveAccess is returned when the API Keys doesnt have access to the account or services
	ErrAPIKeyDoesntHaveAccess = errors.New("API Key doesnt have access to account or Service. Please be aware that the API key needs to have `Read-only access (global:read)` to lookup services by name")

	//ErrNoServiceWithNameExists is returned when no service found with described name
	ErrNoServiceWithNameExists = errors.New("No service called with name exists")

	//ErrNoObjectSetForPurge is returned when no object is declare
	ErrNoObjectSetForPurge = errors.New("No object set to purge")

	//ErrObjectIsNotVaildateURI is returned when the object declared isnt an valid URI
	ErrObjectIsNotVaildateURI = errors.New("Object is not a valid URI")

	//ErrFailedToPurgeService is returned when a purge failed for a service
	ErrFailedToPurgeService = errors.New("Failed to purge the Service")

	//ErrFailedToPurgeObject is returned when a purge failed for an object
	ErrFailedToPurgeObject = errors.New("Failed to purge the Object")

	//ErrFailedToRetrieveServiceDetails is returned when a call failed to restrive a Service
	ErrFailedToRetrieveServiceDetails = errors.New("Failed to retrieve service details")

	errHTTPRequestFailed = errors.New("HTTP request failed")
)
