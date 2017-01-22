package mci

import (
	"io/ioutil"
	"encoding/json"
	"os"
	"errors"
	"strings"
	"github.com/cloudfoundry-community/gautocloud/loader"
	"log"
	"github.com/cloudfoundry-community/gautocloud/logger"
	"github.com/cloudfoundry-community/gautocloud"
)

func PushConfig(mattermostConfig *MattermostConfig, configFilePath string) error {
	configData, err := json.Marshal(mattermostConfig)
	if err != nil {
		return nil
	}
	ioutil.WriteFile(configFilePath, configData, 0644)
	return nil
}
func CloudifyConfig(mattermostConfig *MattermostConfig) error {
	ld := loader.NewLoaderWithLogger(
		gautocloud.CloudEnvs(),
		log.New(os.Stdout, "", log.Ldate | log.Ltime),
		logger.Linfo,
	)
	var err error
	if !ld.IsInACloudEnv() {
		return errors.New("Not in any cloud environment.")
	}
	mattermostConfig.ServiceSettings.ListenAddress = ":" + os.Getenv("PORT")
	if err != nil {
		return err
	}
	if mattermostConfig.ServiceSettings.WebsocketPort == 0 {
		mattermostConfig.ServiceSettings.WebsocketPort = 80
	}
	if mattermostConfig.ServiceSettings.WebsocketSecurePort == 0 {
		mattermostConfig.ServiceSettings.WebsocketSecurePort = 443
	}
	err = cloudifyDatabase(ld, mattermostConfig)
	if err != nil {
		return err
	}
	err = cloudifySmtp(ld, mattermostConfig)
	if err != nil  && !strings.Contains(err.Error(), "cannot be found") {
		return err
	}
	err = cloudifyS3(ld, mattermostConfig)
	if err != nil && !strings.Contains(err.Error(), "cannot be found") {
		return err
	}
	return nil
}

func ExtractConfig(configFilePath string) (*MattermostConfig, error) {
	var err error
	var mattermostConfig MattermostConfig
	configData, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(configData, &mattermostConfig)
	if err != nil {
		return nil, err
	}
	return &mattermostConfig, nil
}