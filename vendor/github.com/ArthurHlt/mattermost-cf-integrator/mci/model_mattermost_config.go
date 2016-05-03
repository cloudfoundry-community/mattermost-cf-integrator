package mci

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
