package mci

type ServiceSettings struct {
	SiteURL             string `json:"SiteURL"`
	ListenAddress       string `json:"ListenAddress"`
	WebsocketSecurePort int    `json:"WebsocketSecurePort"`
	WebsocketPort       int    `json:"WebsocketPort"`
}

type SqlSettings struct {
	DriverName string `json:"DriverName"`
	DataSource string `json:"DataSource"`
}

type FileSettings struct {
	DriverName              string `json:"DriverName"`
	AmazonS3AccessKeyID     string `json:"AmazonS3AccessKeyId"`
	AmazonS3SecretAccessKey string `json:"AmazonS3SecretAccessKey"`
	AmazonS3Bucket          string `json:"AmazonS3Bucket"`
	AmazonS3Endpoint        string `json:"AmazonS3Endpoint"`
}
type LogSettings struct {
	EnableFile bool `json:"EnableFile"`
}
type EmailSettings struct {
	SendEmailNotifications   bool   `json:"SendEmailNotifications"`
	RequireEmailVerification bool   `json:"RequireEmailVerification"`
	SMTPUsername             string `json:"SMTPUsername"`
	SMTPPassword             string `json:"SMTPPassword"`
	SMTPServer               string `json:"SMTPServer"`
	SMTPPort                 string `json:"SMTPPort"`
	ConnectionSecurity       string `json:"ConnectionSecurity"`
}

type MattermostConfig struct {
	ServiceSettings ServiceSettings
	SqlSettings     SqlSettings
	FileSettings    FileSettings
	EmailSettings   EmailSettings
	LogSettings     LogSettings
}
