package mci

import (
	"github.com/cloudfoundry-community/gautocloud/connectors/smtp/raw"
	"github.com/cloudfoundry-community/gautocloud/connectors/smtp/smtptype"
	"github.com/cloudfoundry-community/gautocloud/loader"
	"strconv"
	"strings"
)

var knownSmtp map[string]SmtpType

type SmtpType struct {
	ConnectionSecurity string
}

func init() {
	knownSmtp = make(map[string]SmtpType)
	knownSmtp["sendgrid"] = SmtpType{
		ConnectionSecurity: "STARTTLS",
	}
}
func cloudifySmtp(loader *loader.Loader, mattermostConfig *MattermostConfig) error {
	loader.RegisterConnector(raw.NewSmtpRawConnector())
	var svc smtptype.Smtp
	err := loader.Inject(&svc)
	if err != nil {
		return err
	}
	mattermostConfig.EmailSettings.SendEmailNotifications = true
	mattermostConfig.EmailSettings.RequireEmailVerification = true
	mattermostConfig.EmailSettings.SMTPServer = svc.Host
	mattermostConfig.EmailSettings.SMTPPassword = svc.Password
	mattermostConfig.EmailSettings.SMTPUsername = svc.User
	mattermostConfig.EmailSettings.SMTPPort = strconv.Itoa(svc.Port)
	connectionSecurity := ""
	for key, aSmtp := range knownSmtp {
		if strings.Contains(svc.Host, key) {
			connectionSecurity = aSmtp.ConnectionSecurity
			break
		}
	}
	mattermostConfig.EmailSettings.ConnectionSecurity = connectionSecurity
	return nil
}
