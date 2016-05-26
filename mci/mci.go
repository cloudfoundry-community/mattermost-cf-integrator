package mci

import (
	"io/ioutil"
	"encoding/json"
	"os"
	"github.com/cloudfoundry-community/go-cfenv"
	"errors"
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
	var err error
	if !IsInCloudFoundry() {
		return errors.New("Not in Cloud Foundry environment.")
	}
	mattermostConfig.ServiceSettings.ListenAddress = ":" + os.Getenv("PORT")
	appEnv, err := cfenv.Current()
	if err != nil {
		return err
	}
	if mattermostConfig.ServiceSettings.WebsocketPort == 0 {
		mattermostConfig.ServiceSettings.WebsocketPort = 80
	}
	if mattermostConfig.ServiceSettings.WebsocketSecurePort == 0 {
		mattermostConfig.ServiceSettings.WebsocketSecurePort = 443
	}
	err = cloudifyDatabase(appEnv, mattermostConfig)
	if err != nil {
		return err
	}
	err = cloudifySmtp(appEnv, mattermostConfig)
	if err != nil {
		return err
	}
	err = cloudifyS3(appEnv, mattermostConfig)
	if err != nil {
		return err
	}
	return nil
}

func IsInCloudFoundry() bool {
	d := os.Getenv("VCAP_APPLICATION")
	if d != "" {
		return true
	}
	return false
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