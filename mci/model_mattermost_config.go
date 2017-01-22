package mci

type MattermostConfig struct {
	ServiceSettings      struct {
				     SiteURL                           string `json:"SiteURL"`
				     ListenAddress                     string `json:"ListenAddress"`
				     ConnectionSecurity                string `json:"ConnectionSecurity"`
				     TLSCertFile                       string `json:"TLSCertFile"`
				     TLSKeyFile                        string `json:"TLSKeyFile"`
				     UseLetsEncrypt                    bool `json:"UseLetsEncrypt"`
				     LetsEncryptCertificateCacheFile   string `json:"LetsEncryptCertificateCacheFile"`
				     Forward80To443                    bool `json:"Forward80To443"`
				     ReadTimeout                       int `json:"ReadTimeout"`
				     WriteTimeout                      int `json:"WriteTimeout"`
				     MaximumLoginAttempts              int `json:"MaximumLoginAttempts"`
				     SegmentDeveloperKey               string `json:"SegmentDeveloperKey"`
				     GoogleDeveloperKey                string `json:"GoogleDeveloperKey"`
				     EnableOAuthServiceProvider        bool `json:"EnableOAuthServiceProvider"`
				     EnableIncomingWebhooks            bool `json:"EnableIncomingWebhooks"`
				     EnableOutgoingWebhooks            bool `json:"EnableOutgoingWebhooks"`
				     EnableCommands                    bool `json:"EnableCommands"`
				     EnableOnlyAdminIntegrations       bool `json:"EnableOnlyAdminIntegrations"`
				     EnablePostUsernameOverride        bool `json:"EnablePostUsernameOverride"`
				     EnablePostIconOverride            bool `json:"EnablePostIconOverride"`
				     EnableTesting                     bool `json:"EnableTesting"`
				     EnableDeveloper                   bool `json:"EnableDeveloper"`
				     EnableSecurityFixAlert            bool `json:"EnableSecurityFixAlert"`
				     EnableInsecureOutgoingConnections bool `json:"EnableInsecureOutgoingConnections"`
				     EnableMultifactorAuthentication   bool `json:"EnableMultifactorAuthentication"`
				     EnforceMultifactorAuthentication  bool `json:"EnforceMultifactorAuthentication"`
				     AllowCorsFrom                     string `json:"AllowCorsFrom"`
				     SessionLengthWebInDays            int `json:"SessionLengthWebInDays"`
				     SessionLengthMobileInDays         int `json:"SessionLengthMobileInDays"`
				     SessionLengthSSOInDays            int `json:"SessionLengthSSOInDays"`
				     SessionCacheInMinutes             int `json:"SessionCacheInMinutes"`
				     WebsocketSecurePort               int `json:"WebsocketSecurePort"`
				     WebsocketPort                     int `json:"WebsocketPort"`
				     WebserverMode                     string `json:"WebserverMode"`
				     EnableCustomEmoji                 bool `json:"EnableCustomEmoji"`
				     RestrictCustomEmojiCreation       string `json:"RestrictCustomEmojiCreation"`
				     RestrictPostDelete                string `json:"RestrictPostDelete"`
				     AllowEditPost                     string `json:"AllowEditPost"`
				     PostEditTimeLimit                 int `json:"PostEditTimeLimit"`
			     } `json:"ServiceSettings"`
	TeamSettings         struct {
				     SiteName                         string `json:"SiteName"`
				     MaxUsersPerTeam                  int `json:"MaxUsersPerTeam"`
				     EnableTeamCreation               bool `json:"EnableTeamCreation"`
				     EnableUserCreation               bool `json:"EnableUserCreation"`
				     EnableOpenServer                 bool `json:"EnableOpenServer"`
				     RestrictCreationToDomains        string `json:"RestrictCreationToDomains"`
				     EnableCustomBrand                bool `json:"EnableCustomBrand"`
				     CustomBrandText                  string `json:"CustomBrandText"`
				     CustomDescriptionText            string `json:"CustomDescriptionText"`
				     RestrictDirectMessage            string `json:"RestrictDirectMessage"`
				     RestrictTeamInvite               string `json:"RestrictTeamInvite"`
				     RestrictPublicChannelCreation    string `json:"RestrictPublicChannelCreation"`
				     RestrictPrivateChannelCreation   string `json:"RestrictPrivateChannelCreation"`
				     RestrictPublicChannelManagement  string `json:"RestrictPublicChannelManagement"`
				     RestrictPrivateChannelManagement string `json:"RestrictPrivateChannelManagement"`
				     RestrictPublicChannelDeletion    string `json:"RestrictPublicChannelDeletion"`
				     RestrictPrivateChannelDeletion   string `json:"RestrictPrivateChannelDeletion"`
				     UserStatusAwayTimeout            int `json:"UserStatusAwayTimeout"`
				     MaxChannelsPerTeam               int `json:"MaxChannelsPerTeam"`
				     MaxNotificationsPerChannel       int `json:"MaxNotificationsPerChannel"`
			     } `json:"TeamSettings"`
	SqlSettings          struct {
				     DriverName         string `json:"DriverName"`
				     DataSource         string `json:"DataSource"`
				     DataSourceReplicas []interface{} `json:"DataSourceReplicas"`
				     MaxIdleConns       int `json:"MaxIdleConns"`
				     MaxOpenConns       int `json:"MaxOpenConns"`
				     Trace              bool `json:"Trace"`
				     AtRestEncryptKey   string `json:"AtRestEncryptKey"`
			     } `json:"SqlSettings"`
	LogSettings          struct {
				     EnableConsole          bool `json:"EnableConsole"`
				     ConsoleLevel           string `json:"ConsoleLevel"`
				     EnableFile             bool `json:"EnableFile"`
				     FileLevel              string `json:"FileLevel"`
				     FileFormat             string `json:"FileFormat"`
				     FileLocation           string `json:"FileLocation"`
				     EnableWebhookDebugging bool `json:"EnableWebhookDebugging"`
				     EnableDiagnostics      bool `json:"EnableDiagnostics"`
			     } `json:"LogSettings"`
	PasswordSettings     struct {
				     MinimumLength int `json:"MinimumLength"`
				     Lowercase     bool `json:"Lowercase"`
				     Number        bool `json:"Number"`
				     Uppercase     bool `json:"Uppercase"`
				     Symbol        bool `json:"Symbol"`
			     } `json:"PasswordSettings"`
	FileSettings         struct {
				     MaxFileSize             int `json:"MaxFileSize"`
				     DriverName              string `json:"DriverName"`
				     Directory               string `json:"Directory"`
				     EnablePublicLink        bool `json:"EnablePublicLink"`
				     PublicLinkSalt          string `json:"PublicLinkSalt"`
				     ThumbnailWidth          int `json:"ThumbnailWidth"`
				     ThumbnailHeight         int `json:"ThumbnailHeight"`
				     PreviewWidth            int `json:"PreviewWidth"`
				     PreviewHeight           int `json:"PreviewHeight"`
				     ProfileWidth            int `json:"ProfileWidth"`
				     ProfileHeight           int `json:"ProfileHeight"`
				     InitialFont             string `json:"InitialFont"`
				     AmazonS3AccessKeyID     string `json:"AmazonS3AccessKeyId"`
				     AmazonS3SecretAccessKey string `json:"AmazonS3SecretAccessKey"`
				     AmazonS3Bucket          string `json:"AmazonS3Bucket"`
				     AmazonS3Region          string `json:"AmazonS3Region"`
				     AmazonS3Endpoint        string `json:"AmazonS3Endpoint"`
				     AmazonS3SSL             bool `json:"AmazonS3SSL"`
			     } `json:"FileSettings"`
	EmailSettings        struct {
				     EnableSignUpWithEmail    bool `json:"EnableSignUpWithEmail"`
				     EnableSignInWithEmail    bool `json:"EnableSignInWithEmail"`
				     EnableSignInWithUsername bool `json:"EnableSignInWithUsername"`
				     SendEmailNotifications   bool `json:"SendEmailNotifications"`
				     RequireEmailVerification bool `json:"RequireEmailVerification"`
				     FeedbackName             string `json:"FeedbackName"`
				     FeedbackEmail            string `json:"FeedbackEmail"`
				     FeedbackOrganization     string `json:"FeedbackOrganization"`
				     SMTPUsername             string `json:"SMTPUsername"`
				     SMTPPassword             string `json:"SMTPPassword"`
				     SMTPServer               string `json:"SMTPServer"`
				     SMTPPort                 string `json:"SMTPPort"`
				     ConnectionSecurity       string `json:"ConnectionSecurity"`
				     InviteSalt               string `json:"InviteSalt"`
				     PasswordResetSalt        string `json:"PasswordResetSalt"`
				     SendPushNotifications    bool `json:"SendPushNotifications"`
				     PushNotificationServer   string `json:"PushNotificationServer"`
				     PushNotificationContents string `json:"PushNotificationContents"`
				     EnableEmailBatching      bool `json:"EnableEmailBatching"`
				     EmailBatchingBufferSize  int `json:"EmailBatchingBufferSize"`
				     EmailBatchingInterval    int `json:"EmailBatchingInterval"`
			     } `json:"EmailSettings"`
	RateLimitSettings    struct {
				     Enable           bool `json:"Enable"`
				     PerSec           int `json:"PerSec"`
				     MaxBurst         int `json:"MaxBurst"`
				     MemoryStoreSize  int `json:"MemoryStoreSize"`
				     VaryByRemoteAddr bool `json:"VaryByRemoteAddr"`
				     VaryByHeader     string `json:"VaryByHeader"`
			     } `json:"RateLimitSettings"`
	PrivacySettings      struct {
				     ShowEmailAddress bool `json:"ShowEmailAddress"`
				     ShowFullName     bool `json:"ShowFullName"`
			     } `json:"PrivacySettings"`
	SupportSettings      struct {
				     TermsOfServiceLink string `json:"TermsOfServiceLink"`
				     PrivacyPolicyLink  string `json:"PrivacyPolicyLink"`
				     AboutLink          string `json:"AboutLink"`
				     HelpLink           string `json:"HelpLink"`
				     ReportAProblemLink string `json:"ReportAProblemLink"`
				     SupportEmail       string `json:"SupportEmail"`
			     } `json:"SupportSettings"`
	GitLabSettings       struct {
				     Enable          bool `json:"Enable"`
				     Secret          string `json:"Secret"`
				     ID              string `json:"Id"`
				     Scope           string `json:"Scope"`
				     AuthEndpoint    string `json:"AuthEndpoint"`
				     TokenEndpoint   string `json:"TokenEndpoint"`
				     UserAPIEndpoint string `json:"UserApiEndpoint"`
			     } `json:"GitLabSettings"`
	GoogleSettings       struct {
				     Enable          bool `json:"Enable"`
				     Secret          string `json:"Secret"`
				     ID              string `json:"Id"`
				     Scope           string `json:"Scope"`
				     AuthEndpoint    string `json:"AuthEndpoint"`
				     TokenEndpoint   string `json:"TokenEndpoint"`
				     UserAPIEndpoint string `json:"UserApiEndpoint"`
			     } `json:"GoogleSettings"`
	Office365Settings    struct {
				     Enable          bool `json:"Enable"`
				     Secret          string `json:"Secret"`
				     ID              string `json:"Id"`
				     Scope           string `json:"Scope"`
				     AuthEndpoint    string `json:"AuthEndpoint"`
				     TokenEndpoint   string `json:"TokenEndpoint"`
				     UserAPIEndpoint string `json:"UserApiEndpoint"`
			     } `json:"Office365Settings"`
	LdapSettings         struct {
				     Enable                      bool `json:"Enable"`
				     LdapServer                  string `json:"LdapServer"`
				     LdapPort                    int `json:"LdapPort"`
				     ConnectionSecurity          string `json:"ConnectionSecurity"`
				     BaseDN                      string `json:"BaseDN"`
				     BindUsername                string `json:"BindUsername"`
				     BindPassword                string `json:"BindPassword"`
				     UserFilter                  string `json:"UserFilter"`
				     FirstNameAttribute          string `json:"FirstNameAttribute"`
				     LastNameAttribute           string `json:"LastNameAttribute"`
				     EmailAttribute              string `json:"EmailAttribute"`
				     UsernameAttribute           string `json:"UsernameAttribute"`
				     NicknameAttribute           string `json:"NicknameAttribute"`
				     IDAttribute                 string `json:"IdAttribute"`
				     PositionAttribute           string `json:"PositionAttribute"`
				     SyncIntervalMinutes         int `json:"SyncIntervalMinutes"`
				     SkipCertificateVerification bool `json:"SkipCertificateVerification"`
				     QueryTimeout                int `json:"QueryTimeout"`
				     MaxPageSize                 int `json:"MaxPageSize"`
				     LoginFieldName              string `json:"LoginFieldName"`
			     } `json:"LdapSettings"`
	ComplianceSettings   struct {
				     Enable      bool `json:"Enable"`
				     Directory   string `json:"Directory"`
				     EnableDaily bool `json:"EnableDaily"`
			     } `json:"ComplianceSettings"`
	LocalizationSettings struct {
				     DefaultServerLocale string `json:"DefaultServerLocale"`
				     DefaultClientLocale string `json:"DefaultClientLocale"`
				     AvailableLocales    string `json:"AvailableLocales"`
			     } `json:"LocalizationSettings"`
	SamlSettings         struct {
				     Enable                      bool `json:"Enable"`
				     Verify                      bool `json:"Verify"`
				     Encrypt                     bool `json:"Encrypt"`
				     IdpURL                      string `json:"IdpUrl"`
				     IdpDescriptorURL            string `json:"IdpDescriptorUrl"`
				     AssertionConsumerServiceURL string `json:"AssertionConsumerServiceURL"`
				     IdpCertificateFile          string `json:"IdpCertificateFile"`
				     PublicCertificateFile       string `json:"PublicCertificateFile"`
				     PrivateKeyFile              string `json:"PrivateKeyFile"`
				     FirstNameAttribute          string `json:"FirstNameAttribute"`
				     LastNameAttribute           string `json:"LastNameAttribute"`
				     EmailAttribute              string `json:"EmailAttribute"`
				     UsernameAttribute           string `json:"UsernameAttribute"`
				     NicknameAttribute           string `json:"NicknameAttribute"`
				     LocaleAttribute             string `json:"LocaleAttribute"`
				     PositionAttribute           string `json:"PositionAttribute"`
				     LoginButtonText             string `json:"LoginButtonText"`
			     } `json:"SamlSettings"`
	NativeAppSettings    struct {
				     AppDownloadLink        string `json:"AppDownloadLink"`
				     AndroidAppDownloadLink string `json:"AndroidAppDownloadLink"`
				     IosAppDownloadLink     string `json:"IosAppDownloadLink"`
			     } `json:"NativeAppSettings"`
	ClusterSettings      struct {
				     Enable                 bool `json:"Enable"`
				     InterNodeListenAddress string `json:"InterNodeListenAddress"`
				     InterNodeUrls          []interface{} `json:"InterNodeUrls"`
			     } `json:"ClusterSettings"`
	MetricsSettings      struct {
				     Enable           bool `json:"Enable"`
				     BlockProfileRate int `json:"BlockProfileRate"`
				     ListenAddress    string `json:"ListenAddress"`
			     } `json:"MetricsSettings"`
	AnalyticsSettings    struct {
				     MaxUsersForStatistics int `json:"MaxUsersForStatistics"`
			     } `json:"AnalyticsSettings"`
	WebrtcSettings       struct {
				     Enable              bool `json:"Enable"`
				     GatewayWebsocketURL string `json:"GatewayWebsocketUrl"`
				     GatewayAdminURL     string `json:"GatewayAdminUrl"`
				     GatewayAdminSecret  string `json:"GatewayAdminSecret"`
				     StunURI             string `json:"StunURI"`
				     TurnURI             string `json:"TurnURI"`
				     TurnUsername        string `json:"TurnUsername"`
				     TurnSharedKey       string `json:"TurnSharedKey"`
			     } `json:"WebrtcSettings"`
}
