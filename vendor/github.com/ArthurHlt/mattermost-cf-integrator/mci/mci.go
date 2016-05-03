package mci

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
	"os"
	"github.com/cloudfoundry-community/go-cfenv"
	"net/url"
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
	var dbService *cfenv.Service
	var dbUrl *url.URL
	var err error
	if !IsInCloudFoundry() {
		return errors.New("Not in Cloud Foundry environment.")
	}
	mattermostConfig.Servicesettings.Listenaddress = ":" + os.Getenv("PORT")
	appEnv, err := cfenv.Current()
	if err != nil {
		return err
	}
	dbService, err = getPostgresDb(appEnv)
	if err != nil {
		dbService, err = getMysqlDb(appEnv)
	}else {
		mattermostConfig.Sqlsettings.Drivername = "postgres"
	}
	if err != nil {
		return errors.New("Cannot find database service from cloudfoundry")
	}
	dbUrl, err = parseService(dbService)
	if err != nil {
		return err
	}
	if mattermostConfig.Sqlsettings.Drivername == "postgres" {
		mattermostConfig.Sqlsettings.Datasource = formatDataSource(dbUrl, true)
	}else {
		mattermostConfig.Sqlsettings.Datasource = formatDataSource(dbUrl, false)
	}

	return nil
}
func formatDataSource(dbUrl *url.URL, isPostgres bool) string {
	var postgresString = "postgres://%s:%s@%s%s?sslmode=disable&connect_timeout=10"
	var mysqlString = "%s:%s@tcp(%s)%s?charset=utf8mb4,utf8"
	var dataSource string
	if (isPostgres) {
		dataSource = postgresString
	}else {
		dataSource = mysqlString
	}
	password, _ := dbUrl.User.Password()
	return fmt.Sprintf(
		dataSource,
		dbUrl.User.Username(),
		password,
		dbUrl.Host,
		dbUrl.Path)
}
func getPostgresDb(appEnv *cfenv.App) (*cfenv.Service, error) {
	var service *cfenv.Service
	var services []cfenv.Service
	var err error
	services, err = appEnv.Services.WithTag("postgres")
	if err == nil {
		return &services[0], err
	}
	services, err = appEnv.Services.WithTag("postgresql")
	if err == nil {
		return &services[0], err
	}
	service, err = appEnv.Services.WithName("postgres-mattermost")
	if err == nil {
		return service, err
	}
	return nil, err
}
func getMysqlDb(appEnv *cfenv.App) (*cfenv.Service, error) {
	var service *cfenv.Service
	var services []cfenv.Service
	var err error
	services, err = appEnv.Services.WithTag("mysql")
	if err == nil {
		return &services[0], err
	}
	service, err = appEnv.Services.WithName("mysql-mattermost")
	if err == nil {
		return service, err
	}
	return nil, err
}
func parseService(service *cfenv.Service) (*url.URL, error) {
	return url.Parse(getUriFromService(service))
}
func getUriFromService(service *cfenv.Service) string {
	keys := []string{
		"uri",
		"url",
	}
	for _, key := range keys {
		urlString, ok := service.Credentials[key].(string)
		if !ok {
			continue
		}
		if urlString != "" {
			return urlString
		}
	}
	return ""
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
