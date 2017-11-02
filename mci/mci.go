package mci

import (
	"encoding/json"
	"errors"
	"github.com/cloudfoundry-community/gautocloud"
	"github.com/cloudfoundry-community/gautocloud/loader"
	"github.com/cloudfoundry-community/gautocloud/logger"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func PushConfig(mattermostConfig *MattermostConfig, configFilePath string) error {
	var currentConfig map[string]interface{}
	configData, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(configData, &currentConfig)
	if err != nil {
		return err
	}

	var newConfig map[string]interface{}
	configData, err = json.Marshal(mattermostConfig)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(configData, &newConfig)
	if err != nil {
		return err
	}

	finalConfig := mergeMaps(currentConfig, newConfig)

	configData, err = json.Marshal(finalConfig)
	if err != nil {
		return nil
	}
	ioutil.WriteFile(configFilePath, configData, 0644)
	return nil
}
func mergeMaps(parent, partial map[string]interface{}) map[string]interface{} {
	for k, v := range partial {
		if _, ok := parent[k]; !ok {
			parent[k] = v
			continue
		}
		if vMap, ok := v.(map[string]interface{}); ok {
			parent[k] = mergeMaps(parent[k].(map[string]interface{}), vMap)
			continue
		}
		parent[k] = v
	}
	return parent
}

func CloudifyConfig(mattermostConfig *MattermostConfig) error {
	ld := loader.NewLoaderWithLogger(
		gautocloud.CloudEnvs(),
		log.New(os.Stdout, "", log.Ldate|log.Ltime),
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
	if err != nil && !strings.Contains(err.Error(), "cannot be found") {
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
