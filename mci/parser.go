package mci

import (
	"github.com/cloudfoundry-community/go-cfenv"
	"net/url"
)

func parseBasedUrlService(service *cfenv.Service) (*url.URL, error) {
	return url.Parse(getUriFromService(service))
}
func getUriFromService(service *cfenv.Service) string {
	return getStringValueFromServiceWithKeys(service, "uri", "url")
}

func getStringValueFromService(service *cfenv.Service, key string) string {

	value, ok := service.Credentials[key].(string)
	if !ok {
		return ""
	}
	return value
}
func getStringValueFromServiceWithKeys(service *cfenv.Service, keys ...string) string {
	for _, key := range keys {
		urlString := getStringValueFromService(service, key)
		if urlString != "" {
			return urlString
		}
	}
	return ""
}