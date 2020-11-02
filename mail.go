package p

import (
	"log"
	"strings"

	"github.com/mailgun/mailgun-go"
)

type MailService struct {
	smtpInfo
}

// smtpInfo defines SMTP server vars
type smtpInfo struct {
	Enabled bool
	Domain  string
	Sender  string
	APIKey  string
	build   Build
	filter  string
}

// NewMailService construct a new mail service
func NewMailService(enabled bool, domain string, sender string, apikey string, build Build, filter string) *MailService {
	return &MailService{
		smtpInfo{
			Enabled: enabled,
			Domain:  domain,
			Sender:  sender,
			APIKey:  apikey,
			build:   build,
			filter:  filter,
		},
	}
}

// Send tries to send email with Mailgun
func (s *MailService) Send(addresses []string, subject string, body string) error {
	if s.Enabled == false {
		log.Println("Sending Email is disabled. Skipping .. ")
		return nil
	}

	if s.filter != "" && !strings.Contains(s.filter, s.build.Status) {
		log.Printf("Sending notification to email isnot enabled at this status [%s]. Skipping .. \n", s.build.Status)
		return nil
	}

	log.Printf("Sending notification email to address(es): [%v]: \n", addresses)
	// Don't proceed with sending emails if the addresses list is nil
	if len(addresses) == 0 || addresses == nil {
		log.Println("Skipping sending email because addresses is nil")
		return nil
	}

	mg := mailgun.NewMailgun(s.Domain, s.APIKey)

	message := mg.NewMessage(s.Sender, subject, body, addresses...)

	_, _, err := mg.Send(message)
	return err
}
