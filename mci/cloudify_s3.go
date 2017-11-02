package mci

import (
	"github.com/cloudfoundry-community/gautocloud/connectors/objstorage/objstoretype"
	"github.com/cloudfoundry-community/gautocloud/connectors/objstorage/raw"
	"github.com/cloudfoundry-community/gautocloud/loader"
	"strconv"
)

const DEFAULT_S3_HOST = "s3.amazonaws.com"

func cloudifyS3(loader *loader.Loader, mattermostConfig *MattermostConfig) error {
	loader.RegisterConnector(raw.NewS3RawConnector())
	var svc objstoretype.S3
	err := loader.Inject(&svc)
	if err != nil {
		return err
	}
	mattermostConfig.FileSettings.DriverName = "amazons3"
	mattermostConfig.FileSettings.AmazonS3AccessKeyID = svc.AccessKeyID
	mattermostConfig.FileSettings.AmazonS3SecretAccessKey = svc.SecretAccessKey
	mattermostConfig.FileSettings.AmazonS3Bucket = svc.Bucket
	endpoint := svc.Host
	if endpoint == "" {
		endpoint = DEFAULT_S3_HOST
		svc.UseSsl = true
	}
	if svc.UseSsl {
		endpoint = "https://" + endpoint
		mattermostConfig.FileSettings.AmazonS3SSL = true
	} else {
		endpoint = "http://" + endpoint
		mattermostConfig.FileSettings.AmazonS3SSL = false
	}
	if svc.Port != 0 {
		endpoint += ":" + strconv.Itoa(svc.Port)
	}
	mattermostConfig.FileSettings.AmazonS3Endpoint = endpoint
	return nil
}
