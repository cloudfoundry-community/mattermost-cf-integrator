package mci

import (
	"github.com/cloudfoundry-community/go-cfenv"
)

var knownSmtp map[string]SmtpType

type SmtpType struct {
	Port               string
	ConnectionSecurity string
}
type Smtp struct {
	Hostname string
	Password string
	Username string
	SmtpType SmtpType
}

var defaultSmtpType SmtpType = SmtpType{
	Port: "25",
	ConnectionSecurity: "",
}

func init() {
	knownSmtp = make(map[string]SmtpType)
	knownSmtp["sendgrid"] = SmtpType{
		Port: "587",
		ConnectionSecurity: "STARTTLS",
	}
}
func cloudifySmtp(appEnv *cfenv.App, mattermostConfig *MattermostConfig) error {
	smtp, err := getSmtp(appEnv)
	if err != nil {
		return nil
	}
	mattermostConfig.EmailSettings.SendEmailNotifications = true
	mattermostConfig.EmailSettings.RequireEmailVerification = true
	mattermostConfig.EmailSettings.SMTPServer = smtp.Hostname
	mattermostConfig.EmailSettings.SMTPPassword = smtp.Password
	mattermostConfig.EmailSettings.SMTPUsername = smtp.Username
	mattermostConfig.EmailSettings.SMTPPort = smtp.SmtpType.Port
	mattermostConfig.EmailSettings.ConnectionSecurity = smtp.SmtpType.ConnectionSecurity
	return nil
}
func getSmtp(appEnv *cfenv.App) (Smtp, error) {
	var smtp Smtp
	var smtpType SmtpType
	service, err := getServiceSmtp(appEnv)
	if err != nil {
		return smtp, err
	}
	if val, ok := knownSmtp[service.Label]; ok {
		smtpType = val
	} else {
		smtpType = defaultSmtpType
	}
	smtp = Smtp{
		Hostname: getStringValueFromServiceWithKeys(service, "hostname", "host", "server"),
		Password: getStringValueFromServiceWithKeys(service, "password", "pass"),
		Username: getStringValueFromServiceWithKeys(service, "username", "user"),
		SmtpType: smtpType,
	}
	return smtp, nil
}
func getServiceSmtp(appEnv *cfenv.App) (*cfenv.Service, error) {
	var service *cfenv.Service
	var services []cfenv.Service
	var err error
	services, err = appEnv.Services.WithTag("smtp")
	if err == nil {
		return &services[0], nil
	}
	services, err = appEnv.Services.WithTag("email")
	if err == nil {
		return &services[0], nil
	}
	services, err = appEnv.Services.WithTag("Email")
	if err == nil {
		return &services[0], nil
	}
	service, err = appEnv.Services.WithName("smtp-mattermost")
	if err == nil {
		return service, nil
	}
	return nil, err
}

