package fastlyclient

import "errors"

var (
	ErrorNoAPIKeySet                    = errors.New("Fastly API key not set. Please export $FASTLYAPIKEY=x")
	ErrorNoServiceWithNameExists        = errors.New("No service called with name exists")
	ErrorNoObjectSetForPurge            = errors.New("No object set to purge")
	ErrorFailedToPurgeService           = errors.New("Failed to purge the Service")
	ErrorFailedToPurgeObject            = errors.New("Failed to purge the Object")
	ErrorFailedToRetrieveServiceDetails = errors.New("Failed to retrieve service details")
)
