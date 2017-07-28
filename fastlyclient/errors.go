package fastlyclient

import "errors"

var (
	//ErrorNoAPIKeySet Error when no API Key set
	ErrorNoAPIKeySet = errors.New("Fastly API key not set. Please export $FASTLYAPIKEY=x")

	//ErrorNoServiceWithNameExists when no service found with described name
	ErrorNoServiceWithNameExists = errors.New("No service called with name exists")

	//ErrorNoObjectSetForPurge when you dont declare an object to be purged
	ErrorNoObjectSetForPurge = errors.New("No object set to purge")

	//ErrorObjectIsNotVaildateURI object declared for purge isnt an valid URI
	ErrorObjectIsNotVaildateURI = errors.New("Object is not a valid URI")

	//ErrorFailedToPurgeService failed to purge all objects from a service
	ErrorFailedToPurgeService = errors.New("Failed to purge the Service")

	//ErrorFailedToPurgeObject failed to purge and object
	ErrorFailedToPurgeObject = errors.New("Failed to purge the Object")

	//ErrorFailedToRetrieveServiceDetails failed to restrive the Service details
	ErrorFailedToRetrieveServiceDetails = errors.New("Failed to retrieve service details")
)
