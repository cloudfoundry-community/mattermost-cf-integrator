package mci

type ServiceSettings struct {
	SiteURL                                  string
	LicenseFileLocation                      string
	ListenAddress                            string
	ConnectionSecurity                       string
	TLSCertFile                              string
	TLSKeyFile                               string
	UseLetsEncrypt                           bool
	LetsEncryptCertificateCacheFile          string
	Forward80To443                           bool
	ReadTimeout                              int
	WriteTimeout                             int
	MaximumLoginAttempts                     int
	GoroutineHealthThreshold                 int
	GoogleDeveloperKey                       string
	EnableOAuthServiceProvider               bool
	EnableIncomingWebhooks                   bool
	EnableOutgoingWebhooks                   bool
	EnableCommands                           bool
	EnableOnlyAdminIntegrations              bool
	EnablePostUsernameOverride               bool
	EnablePostIconOverride                   bool
	EnableAPIv3                              bool
	EnableLinkPreviews                       bool
	EnableTesting                            bool
	EnableDeveloper                          bool
	EnableSecurityFixAlert                   bool
	EnableInsecureOutgoingConnections        bool
	EnableMultifactorAuthentication          bool
	EnforceMultifactorAuthentication         bool
	AllowCorsFrom                            string
	SessionLengthWebInDays                   int
	SessionLengthMobileInDays                int
	SessionLengthSSOInDays                   int
	SessionCacheInMinutes                    int
	WebsocketSecurePort                      int
	WebsocketPort                            int
	WebserverMode                            string
	EnableCustomEmoji                        bool
	EnableEmojiPicker                        bool
	RestrictCustomEmojiCreation              string
	RestrictPostDelete                       string
	AllowEditPost                            string
	PostEditTimeLimit                        int
	TimeBetweenUserTypingUpdatesMilliseconds int64
	EnablePostSearch                         bool
	EnableUserTypingMessages                 bool
	EnableChannelViewedMessages              bool
	EnableUserStatuses                       bool
	ClusterLogTimeoutMilliseconds            int
}

type ClusterSettings struct {
	Enable                bool
	ClusterName           string
	OverrideHostname      string
	UseIpAddress          bool
	UseExperimentalGossip bool
	ReadOnlyConfig        bool
	GossipPort            int
	StreamingPort         int
}

type MetricsSettings struct {
	Enable           bool
	BlockProfileRate int
	ListenAddress    string
}

type AnalyticsSettings struct {
	MaxUsersForStatistics int
}

type SSOSettings struct {
	Enable          bool
	Secret          string
	Id              string
	Scope           string
	AuthEndpoint    string
	TokenEndpoint   string
	UserApiEndpoint string
}

type SqlSettings struct {
	DriverName               string
	DataSource               string
	DataSourceReplicas       []string
	DataSourceSearchReplicas []string
	MaxIdleConns             int
	MaxOpenConns             int
	Trace                    bool
	AtRestEncryptKey         string
	QueryTimeout             int
}

type LogSettings struct {
	EnableConsole          bool
	ConsoleLevel           string
	EnableFile             bool
	FileLevel              string
	FileFormat             string
	FileLocation           string
	EnableWebhookDebugging bool
	EnableDiagnostics      bool
}

type PasswordSettings struct {
	MinimumLength int
	Lowercase     bool
	Number        bool
	Uppercase     bool
	Symbol        bool
}

type FileSettings struct {
	EnableFileAttachments   bool
	MaxFileSize             int64
	DriverName              string
	Directory               string
	EnablePublicLink        bool
	PublicLinkSalt          string
	InitialFont             string
	AmazonS3AccessKeyId     string
	AmazonS3SecretAccessKey string
	AmazonS3Bucket          string
	AmazonS3Region          string
	AmazonS3Endpoint        string
	AmazonS3SSL             bool
	AmazonS3SignV2          bool
}

type EmailSettings struct {
	EnableSignUpWithEmail             bool
	EnableSignInWithEmail             bool
	EnableSignInWithUsername          bool
	SendEmailNotifications            bool
	RequireEmailVerification          bool
	FeedbackName                      string
	FeedbackEmail                     string
	FeedbackOrganization              string
	SMTPUsername                      string
	SMTPPassword                      string
	SMTPServer                        string
	SMTPPort                          string
	ConnectionSecurity                string
	InviteSalt                        string
	SendPushNotifications             bool
	PushNotificationServer            string
	PushNotificationContents          string
	EnableEmailBatching               bool
	EmailBatchingBufferSize           int
	EmailBatchingInterval             int
	SkipServerCertificateVerification bool
}

type RateLimitSettings struct {
	Enable           bool
	PerSec           int
	MaxBurst         int
	MemoryStoreSize  int
	VaryByRemoteAddr bool
	VaryByHeader     string
}

type PrivacySettings struct {
	ShowEmailAddress bool
	ShowFullName     bool
}

type SupportSettings struct {
	TermsOfServiceLink       string
	PrivacyPolicyLink        string
	AboutLink                string
	HelpLink                 string
	ReportAProblemLink       string
	AdministratorsGuideLink  string
	TroubleshootingForumLink string
	CommercialSupportLink    string
	SupportEmail             string
}

type AnnouncementSettings struct {
	EnableBanner         bool
	BannerText           string
	BannerColor          string
	BannerTextColor      string
	AllowBannerDismissal bool
}

type TeamSettings struct {
	SiteName                            string
	MaxUsersPerTeam                     int
	EnableTeamCreation                  bool
	EnableUserCreation                  bool
	EnableOpenServer                    bool
	RestrictCreationToDomains           string
	EnableCustomBrand                   bool
	CustomBrandText                     string
	CustomDescriptionText               string
	RestrictDirectMessage               string
	RestrictTeamInvite                  string
	RestrictPublicChannelManagement     string
	RestrictPrivateChannelManagement    string
	RestrictPublicChannelCreation       string
	RestrictPrivateChannelCreation      string
	RestrictPublicChannelDeletion       string
	RestrictPrivateChannelDeletion      string
	RestrictPrivateChannelManageMembers string
	UserStatusAwayTimeout               int64
	MaxChannelsPerTeam                  int64
	MaxNotificationsPerChannel          int64
	TeammateNameDisplay                 string
}

type LdapSettings struct {
	// Basic
	Enable             bool
	LdapServer         string
	LdapPort           int
	ConnectionSecurity string
	BaseDN             string
	BindUsername       string
	BindPassword       string

	// Filtering
	UserFilter string

	// User Mapping
	FirstNameAttribute string
	LastNameAttribute  string
	EmailAttribute     string
	UsernameAttribute  string
	NicknameAttribute  string
	IdAttribute        string
	PositionAttribute  string

	// Syncronization
	SyncIntervalMinutes int

	// Advanced
	SkipCertificateVerification bool
	QueryTimeout                int
	MaxPageSize                 int

	// Customization
	LoginFieldName string
}

type ComplianceSettings struct {
	Enable      bool
	Directory   string
	EnableDaily bool
}

type LocalizationSettings struct {
	DefaultServerLocale string
	DefaultClientLocale string
	AvailableLocales    string
}

type SamlSettings struct {
	// Basic
	Enable  bool
	Verify  bool
	Encrypt bool

	IdpUrl                      string
	IdpDescriptorUrl            string
	AssertionConsumerServiceURL string

	IdpCertificateFile    string
	PublicCertificateFile string
	PrivateKeyFile        string

	// User Mapping
	FirstNameAttribute string
	LastNameAttribute  string
	EmailAttribute     string
	UsernameAttribute  string
	NicknameAttribute  string
	LocaleAttribute    string
	PositionAttribute  string

	LoginButtonText string
}

type NativeAppSettings struct {
	AppDownloadLink        string
	AndroidAppDownloadLink string
	IosAppDownloadLink     string
}

type WebrtcSettings struct {
	Enable              bool
	GatewayWebsocketUrl string
	GatewayAdminUrl     string
	GatewayAdminSecret  string
	StunURI             string
	TurnURI             string
	TurnUsername        string
	TurnSharedKey       string
}

type ElasticSearchSettings struct {
	ConnectionUrl   string
	Username        string
	Password        string
	EnableIndexing  bool
	EnableSearching bool
	Sniff           bool
}

type DataRetentionSettings struct {
	Enable bool
}

type MattermostConfig struct {
	ServiceSettings       ServiceSettings
	TeamSettings          TeamSettings
	SqlSettings           SqlSettings
	LogSettings           LogSettings
	PasswordSettings      PasswordSettings
	FileSettings          FileSettings
	EmailSettings         EmailSettings
	RateLimitSettings     RateLimitSettings
	PrivacySettings       PrivacySettings
	SupportSettings       SupportSettings
	AnnouncementSettings  AnnouncementSettings
	GitLabSettings        SSOSettings
	GoogleSettings        SSOSettings
	Office365Settings     SSOSettings
	LdapSettings          LdapSettings
	ComplianceSettings    ComplianceSettings
	LocalizationSettings  LocalizationSettings
	SamlSettings          SamlSettings
	NativeAppSettings     NativeAppSettings
	ClusterSettings       ClusterSettings
	MetricsSettings       MetricsSettings
	AnalyticsSettings     AnalyticsSettings
	WebrtcSettings        WebrtcSettings
	ElasticSearchSettings ElasticSearchSettings
	DataRetentionSettings DataRetentionSettings
}
