package fastly

import (
	"errors"
	"net/url"
)

//PurgeObject an object from the service
func (f *Fastly) PurgeObject(serviceName string, object string) (string, error) {
	service, err := f.getServiceByName(serviceName)
	if err != nil {
		return "", err
	}

	if f.isStringEmpty(object) {
		return "", ErrNoObjectSetForPurge
	}

	_, err = url.ParseRequestURI(object)
	if err != nil {
		return "", ErrObjectIsNotVaildateURI
	}

	//PURGE /service/ekjhsdfkjhsdfouejk/purge??
	_, err = f.purge(fastlyAPIEndPoint + "/service/" + service.ID + "/" + object)
	if err != nil {
		if errors.Is(err, errHTTPRequestFailed) {
			return "", ErrFailedToPurgeService
		}
		return "", err
	}

	return RespSucessfullyPurgedObject, nil
}

//PurgeAllObjects for service
func (f *Fastly) PurgeAllObjects(serviceName string) (string, error) {
	service, err := f.getServiceByName(serviceName)
	if err != nil {
		return "", err
	}

	//POST /service/{serviceID}/purge_all
	_, err = f.post(fastlyAPIEndPoint+"/service/"+service.ID+"/purge_all", nil)
	if err != nil {
		if errors.Is(err, errHTTPRequestFailed) {
			return "", ErrFailedToPurgeService
		}
		return "", err
	}

	return RespSucessfullyPurgedAllObjects, err
}
