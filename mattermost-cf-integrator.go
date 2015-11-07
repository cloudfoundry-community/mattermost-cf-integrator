package main

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"github.com/cloudfoundry-community/go-cfenv"
	"net/url"
	"errors"
	"os/exec"
)

type MattermostConfig struct {
	Servicesettings   struct {
						  Listenaddress              string `json:"ListenAddress"`
						  Maximumloginattempts       int `json:"MaximumLoginAttempts"`
						  Segmentdeveloperkey        string `json:"SegmentDeveloperKey"`
						  Googledeveloperkey         string `json:"GoogleDeveloperKey"`
						  Enableoauthserviceprovider bool `json:"EnableOAuthServiceProvider"`
						  Enableincomingwebhooks     bool `json:"EnableIncomingWebhooks"`
						  Enableoutgoingwebhooks     bool `json:"EnableOutgoingWebhooks"`
						  Enablepostusernameoverride bool `json:"EnablePostUsernameOverride"`
						  Enableposticonoverride     bool `json:"EnablePostIconOverride"`
						  Enabletesting              bool `json:"EnableTesting"`
						  Enablesecurityfixalert     bool `json:"EnableSecurityFixAlert"`
					  } `json:"ServiceSettings"`
	Teamsettings      struct {
						  Sitename                  string `json:"SiteName"`
						  Maxusersperteam           int `json:"MaxUsersPerTeam"`
						  Enableteamcreation        bool `json:"EnableTeamCreation"`
						  Enableusercreation        bool `json:"EnableUserCreation"`
						  Restrictcreationtodomains string `json:"RestrictCreationToDomains"`
						  Restrictteamnames         bool `json:"RestrictTeamNames"`
						  Enableteamlisting         bool `json:"EnableTeamListing"`
					  } `json:"TeamSettings"`
	Sqlsettings       struct {
						  Drivername         string `json:"DriverName"`
						  Datasource         string `json:"DataSource"`
						  Datasourcereplicas []interface{} `json:"DataSourceReplicas"`
						  Maxidleconns       int `json:"MaxIdleConns"`
						  Maxopenconns       int `json:"MaxOpenConns"`
						  Trace              bool `json:"Trace"`
						  Atrestencryptkey   string `json:"AtRestEncryptKey"`
					  } `json:"SqlSettings"`
	Logsettings       struct {
						  Enableconsole bool `json:"EnableConsole"`
						  Consolelevel  string `json:"ConsoleLevel"`
						  Enablefile    bool `json:"EnableFile"`
						  Filelevel     string `json:"FileLevel"`
						  Fileformat    string `json:"FileFormat"`
						  Filelocation  string `json:"FileLocation"`
					  } `json:"LogSettings"`
	Filesettings      struct {
						  Drivername              string `json:"DriverName"`
						  Directory               string `json:"Directory"`
						  Enablepubliclink        bool `json:"EnablePublicLink"`
						  Publiclinksalt          string `json:"PublicLinkSalt"`
						  Thumbnailwidth          int `json:"ThumbnailWidth"`
						  Thumbnailheight         int `json:"ThumbnailHeight"`
						  Previewwidth            int `json:"PreviewWidth"`
						  Previewheight           int `json:"PreviewHeight"`
						  Profilewidth            int `json:"ProfileWidth"`
						  Profileheight           int `json:"ProfileHeight"`
						  Initialfont             string `json:"InitialFont"`
						  Amazons3Accesskeyid     string `json:"AmazonS3AccessKeyId"`
						  Amazons3Secretaccesskey string `json:"AmazonS3SecretAccessKey"`
						  Amazons3Bucket          string `json:"AmazonS3Bucket"`
						  Amazons3Region          string `json:"AmazonS3Region"`
					  } `json:"FileSettings"`
	Emailsettings     struct {
						  Enablesignupwithemail    bool `json:"EnableSignUpWithEmail"`
						  Sendemailnotifications   bool `json:"SendEmailNotifications"`
						  Requireemailverification bool `json:"RequireEmailVerification"`
						  Feedbackname             string `json:"FeedbackName"`
						  Feedbackemail            string `json:"FeedbackEmail"`
						  Smtpusername             string `json:"SMTPUsername"`
						  Smtppassword             string `json:"SMTPPassword"`
						  Smtpserver               string `json:"SMTPServer"`
						  Smtpport                 string `json:"SMTPPort"`
						  Connectionsecurity       string `json:"ConnectionSecurity"`
						  Invitesalt               string `json:"InviteSalt"`
						  Passwordresetsalt        string `json:"PasswordResetSalt"`
						  Applepushserver          string `json:"ApplePushServer"`
						  Applepushcertpublic      string `json:"ApplePushCertPublic"`
						  Applepushcertprivate     string `json:"ApplePushCertPrivate"`
					  } `json:"EmailSettings"`
	Ratelimitsettings struct {
						  Enableratelimiter bool `json:"EnableRateLimiter"`
						  Persec            int `json:"PerSec"`
						  Memorystoresize   int `json:"MemoryStoreSize"`
						  Varybyremoteaddr  bool `json:"VaryByRemoteAddr"`
						  Varybyheader      string `json:"VaryByHeader"`
					  } `json:"RateLimitSettings"`
	Privacysettings   struct {
						  Showemailaddress bool `json:"ShowEmailAddress"`
						  Showfullname     bool `json:"ShowFullName"`
					  } `json:"PrivacySettings"`
	Gitlabsettings    struct {
						  Enable          bool `json:"Enable"`
						  Secret          string `json:"Secret"`
						  ID              string `json:"Id"`
						  Scope           string `json:"Scope"`
						  Authendpoint    string `json:"AuthEndpoint"`
						  Tokenendpoint   string `json:"TokenEndpoint"`
						  Userapiendpoint string `json:"UserApiEndpoint"`
					  } `json:"GitLabSettings"`
}

func main() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	configFilePath := path.Join(wd, "config", "config.json")
	config, err := extractConfig(configFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	err = cloudifyConfig(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	pushConfig(config, configFilePath)
	mattermostExec := exec.Command(path.Join(wd, "bin", "platform"))
	mattermostExec.Stdout = os.Stdout
	mattermostExec.Stderr = os.Stderr
	err = mattermostExec.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

}
func pushConfig(mattermostConfig *MattermostConfig, configFilePath string) error {
	configData, err := json.Marshal(mattermostConfig)
	if err != nil {
		return nil
	}
	ioutil.WriteFile(configFilePath, configData, 0644)
	return nil
}
func cloudifyConfig(mattermostConfig *MattermostConfig) error {
	var dbService cfenv.Service
	var dbUrl *url.URL
	var err error
	if !isInCloudFoundry() {
		return nil
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
	var postgresString = "postgres://%s:%s@%s/%s?sslmode=disable&connect_timeout=10"
	var mysqlString = "%s:%s@tcp(%s)/%s?charset=utf8mb4,utf8"
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
func getPostgresDb(appEnv *cfenv.App) (cfenv.Service, error) {
	var service []cfenv.Service
	var err error
	service, err = appEnv.Services.WithTag("postgres")
	if err == nil {
		return service[0], err
	}
	service, err = appEnv.Services.WithTag("postgresql")
	if err == nil {
		return service[0], err
	}
	service, err = appEnv.Services.WithTag("postgres-mattermost")
	if err == nil {
		return service[0], err
	}
	return cfenv.Service{}, err
}
func getMysqlDb(appEnv *cfenv.App) (cfenv.Service, error) {
	var service []cfenv.Service
	var err error
	service, err = appEnv.Services.WithTag("mysql")
	if err == nil {
		return service[0], err
	}
	service, err = appEnv.Services.WithTag("mysql-mattermost")
	if err == nil {
		return service[0], err
	}
	return cfenv.Service{}, err
}
func parseService(service cfenv.Service) (*url.URL, error) {
	return url.Parse(getUriFromService(service))
}
func getUriFromService(service cfenv.Service) string {
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
func isInCloudFoundry() bool {
	d := os.Getenv("VCAP_APPLICATION")
	if d != "" {
		return true
	}
	return false
}
func extractConfig(configFilePath string) (*MattermostConfig, error) {
	var err error
	var mattermostConfig MattermostConfig
	configData, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return &mattermostConfig, err
	}
	err = json.Unmarshal(configData, &mattermostConfig)
	if err != nil {
		return nil, err
	}
	return &mattermostConfig, nil
}