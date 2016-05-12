package mci

import (
	"github.com/cloudfoundry-community/go-cfenv"
	"net/url"
)

func cloudifyS3(appEnv *cfenv.App, mattermostConfig *MattermostConfig) error {
	var s3Service *cfenv.Service
	var s3Url *url.URL
	var err error
	s3Service, err = getS3Service(appEnv)
	if err != nil {
		return nil
	}
	s3Url, err = parseBasedUrlService(s3Service)
	if err != nil {
		return err
	}
	accessKeyId, err := url.QueryUnescape(s3Url.User.Username())
	if err != nil {
		return err
	}
	secretAccessKey, _ := s3Url.User.Password()
	secretAccessKey, err = url.QueryUnescape(secretAccessKey)
	if err != nil {
		return err
	}
	s3Url, err = parseBasedUrlService(s3Service)
	mattermostConfig.FileSettings.DriverName = "amazons3"
	mattermostConfig.FileSettings.AmazonS3AccessKeyID = accessKeyId
	mattermostConfig.FileSettings.AmazonS3SecretAccessKey = secretAccessKey
	mattermostConfig.FileSettings.AmazonS3Bucket = s3Url.Path[1:]
	mattermostConfig.FileSettings.AmazonS3Endpoint = "https://" + s3Url.Host
	return nil
}

func getS3Service(appEnv *cfenv.App) (*cfenv.Service, error) {
	var service *cfenv.Service
	var services []cfenv.Service
	var err error
	services, err = appEnv.Services.WithTagUsingPattern(".*s3.*")
	if err == nil {
		return &services[0], err
	}
	services, err = appEnv.Services.WithTag(".*riak.*")
	if err == nil {
		return &services[0], err
	}
	service, err = appEnv.Services.WithName("s3-mattermost")
	if err == nil {
		return service, err
	}
	return nil, err
}
