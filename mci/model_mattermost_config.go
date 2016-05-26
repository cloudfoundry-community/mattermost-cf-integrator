package mci

type MattermostConfig struct {
	ServiceSettings   struct {
				  ListenAddress                     string `json:"ListenAddress"`
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
				  AllowCorsFrom                     string `json:"AllowCorsFrom"`
				  SessionLengthWebInDays            int `json:"SessionLengthWebInDays"`
				  SessionLengthMobileInDays         int `json:"SessionLengthMobileInDays"`
				  SessionLengthSSOInDays            int `json:"SessionLengthSSOInDays"`
				  SessionCacheInMinutes             int `json:"SessionCacheInMinutes"`
				  WebsocketSecurePort               int `json:"WebsocketSecurePort"`
				  WebsocketPort	                    int `json:"WebsocketPort"`
			  } `json:"ServiceSettings"`
	TeamSettings      struct {
				  SiteName                  string `json:"SiteName"`
				  MaxUsersPerTeam           int `json:"MaxUsersPerTeam"`
				  EnableTeamCreation        bool `json:"EnableTeamCreation"`
				  EnableUserCreation        bool `json:"EnableUserCreation"`
				  RestrictCreationToDomains string `json:"RestrictCreationToDomains"`
				  RestrictTeamNames         bool `json:"RestrictTeamNames"`
				  EnableTeamListing         bool `json:"EnableTeamListing"`
			  } `json:"TeamSettings"`
	SqlSettings       struct {
				  DriverName         string `json:"DriverName"`
				  DataSource         string `json:"DataSource"`
				  DataSourceReplicas []interface{} `json:"DataSourceReplicas"`
				  MaxIdleConns       int `json:"MaxIdleConns"`
				  MaxOpenConns       int `json:"MaxOpenConns"`
				  Trace              bool `json:"Trace"`
				  AtRestEncryptKey   string `json:"AtRestEncryptKey"`
			  } `json:"SqlSettings"`
	LogSettings       struct {
				  EnableConsole bool `json:"EnableConsole"`
				  ConsoleLevel  string `json:"ConsoleLevel"`
				  EnableFile    bool `json:"EnableFile"`
				  FileLevel     string `json:"FileLevel"`
				  FileFormat    string `json:"FileFormat"`
				  FileLocation  string `json:"FileLocation"`
			  } `json:"LogSettings"`
	FileSettings      struct {
				  DriverName                 string `json:"DriverName"`
				  Directory                  string `json:"Directory"`
				  EnablePublicLink           bool `json:"EnablePublicLink"`
				  PublicLinkSalt             string `json:"PublicLinkSalt"`
				  ThumbnailWidth             int `json:"ThumbnailWidth"`
				  ThumbnailHeight            int `json:"ThumbnailHeight"`
				  PreviewWidth               int `json:"PreviewWidth"`
				  PreviewHeight              int `json:"PreviewHeight"`
				  ProfileWidth               int `json:"ProfileWidth"`
				  ProfileHeight              int `json:"ProfileHeight"`
				  InitialFont                string `json:"InitialFont"`
				  AmazonS3AccessKeyID        string `json:"AmazonS3AccessKeyId"`
				  AmazonS3SecretAccessKey    string `json:"AmazonS3SecretAccessKey"`
				  AmazonS3Bucket             string `json:"AmazonS3Bucket"`
				  AmazonS3Region             string `json:"AmazonS3Region"`
				  AmazonS3Endpoint           string `json:"AmazonS3Endpoint"`
				  AmazonS3BucketEndpoint     string `json:"AmazonS3BucketEndpoint"`
				  AmazonS3LocationConstraint bool `json:"AmazonS3LocationConstraint"`
				  AmazonS3LowercaseBucket    bool `json:"AmazonS3LowercaseBucket"`
			  } `json:"FileSettings"`
	EmailSettings     struct {
				  EnableSignUpWithEmail    bool `json:"EnableSignUpWithEmail"`
				  EnableSignInWithEmail    bool `json:"EnableSignInWithEmail"`
				  EnableSignInWithUsername bool `json:"EnableSignInWithUsername"`
				  SendEmailNotifications   bool `json:"SendEmailNotifications"`
				  RequireEmailVerification bool `json:"RequireEmailVerification"`
				  FeedbackName             string `json:"FeedbackName"`
				  FeedbackEmail            string `json:"FeedbackEmail"`
				  SMTPUsername             string `json:"SMTPUsername"`
				  SMTPPassword             string `json:"SMTPPassword"`
				  SMTPServer               string `json:"SMTPServer"`
				  SMTPPort                 string `json:"SMTPPort"`
				  ConnectionSecurity       string `json:"ConnectionSecurity"`
				  InviteSalt               string `json:"InviteSalt"`
				  PasswordResetSalt        string `json:"PasswordResetSalt"`
				  SendPushNotifications    bool `json:"SendPushNotifications"`
				  PushNotificationServer   string `json:"PushNotificationServer"`
			  } `json:"EmailSettings"`
	RateLimitSettings struct {
				  EnableRateLimiter bool `json:"EnableRateLimiter"`
				  PerSec            int `json:"PerSec"`
				  MemoryStoreSize   int `json:"MemoryStoreSize"`
				  VaryByRemoteAddr  bool `json:"VaryByRemoteAddr"`
				  VaryByHeader      string `json:"VaryByHeader"`
			  } `json:"RateLimitSettings"`
	PrivacySettings   struct {
				  ShowEmailAddress bool `json:"ShowEmailAddress"`
				  ShowFullName     bool `json:"ShowFullName"`
			  } `json:"PrivacySettings"`
	SupportSettings   struct {
				  TermsOfServiceLink string `json:"TermsOfServiceLink"`
				  PrivacyPolicyLink  string `json:"PrivacyPolicyLink"`
				  AboutLink          string `json:"AboutLink"`
				  HelpLink           string `json:"HelpLink"`
				  ReportAProblemLink string `json:"ReportAProblemLink"`
				  SupportEmail       string `json:"SupportEmail"`
			  } `json:"SupportSettings"`
	GitLabSettings    struct {
				  Enable          bool `json:"Enable"`
				  Secret          string `json:"Secret"`
				  ID              string `json:"Id"`
				  Scope           string `json:"Scope"`
				  AuthEndpoint    string `json:"AuthEndpoint"`
				  TokenEndpoint   string `json:"TokenEndpoint"`
				  UserAPIEndpoint string `json:"UserApiEndpoint"`
			  } `json:"GitLabSettings"`
}
