package mci

import (
	"github.com/cloudfoundry-community/go-cfenv"
	"net/url"
	"errors"
)

type S3Service struct {
	AccessKeyID     string
	SecretAccessKey string
	Bucket          string
	Endpoint        string
}

const DEFAULT_S3_HOST = "s3.amazonaws.com"

func cloudifyS3(appEnv *cfenv.App, mattermostConfig *MattermostConfig) error {
	var s3ServiceFromCf *cfenv.Service
	var s3Url *url.URL
	var err error
	var s3Service S3Service
	s3ServiceFromCf, err = getS3Service(appEnv)
	if err != nil {
		return nil
	}
	s3Url, err = parseBasedUrlService(s3ServiceFromCf)
	if err != nil {
		return err
	}
	if s3Url.Host != "" {
		s3Service, err = generateS3ServiceFromUrl(s3Url)
		if err != nil {
			return err
		}
	} else {
		s3Service, err = generateS3ServiceFromNonUriBasedS3Broker(s3ServiceFromCf)
		if err != nil {
			return err
		}
	}
	mattermostConfig.FileSettings.DriverName = "amazons3"
	mattermostConfig.FileSettings.AmazonS3AccessKeyID = s3Service.AccessKeyID
	mattermostConfig.FileSettings.AmazonS3SecretAccessKey = s3Service.SecretAccessKey
	mattermostConfig.FileSettings.AmazonS3Bucket = s3Service.Bucket
	mattermostConfig.FileSettings.AmazonS3Endpoint = s3Service.Endpoint
	return nil
}
func generateS3ServiceFromUrl(s3Url *url.URL) (S3Service, error) {
	var s3Service S3Service
	accessKeyId, err := url.QueryUnescape(s3Url.User.Username())
	if err != nil {
		return s3Service, err
	}
	secretAccessKey, _ := s3Url.User.Password()
	secretAccessKey, err = url.QueryUnescape(secretAccessKey)
	if err != nil {
		return s3Service, err
	}
	s3Service = S3Service{
		AccessKeyID: accessKeyId,
		SecretAccessKey: secretAccessKey,
		Bucket: s3Url.Path[1:],
		Endpoint: "https://" + s3Url.Host,
	}
	return s3Service, nil
}
func generateS3ServiceFromNonUriBasedS3Broker(s3ServiceFromCf *cfenv.Service) (S3Service, error) {
	var s3Service S3Service
	accessKeyId := getStringValueFromServiceWithKeys(s3ServiceFromCf, "access_key_id", "accessKeyId", "access-key-id")
	if accessKeyId == "" {
		return s3Service, errors.New("Can't parse access key id.")
	}
	secretAccessKey := getStringValueFromServiceWithKeys(s3ServiceFromCf, "secret_access_key", "secretAccessKey", "secret-access-key")
	if secretAccessKey == "" {
		return s3Service, errors.New("Can't parse secret access key.")
	}
	bucket := getStringValueFromServiceWithKeys(s3ServiceFromCf, "bucket")
	if bucket == "" {
		return s3Service, errors.New("Can't parse bucket name.")
	}
	host := getStringValueFromServiceWithKeys(s3ServiceFromCf, "host", "hostname")
	if host == "" {
		host = DEFAULT_S3_HOST
	}
	s3Service = S3Service{
		AccessKeyID: accessKeyId,
		SecretAccessKey: secretAccessKey,
		Bucket: bucket,
		Endpoint: "https://" + host,
	}
	return s3Service, nil
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
