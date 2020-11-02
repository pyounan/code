package p

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/Netflix/go-env"
)

// HelloPubSub consumes a Pub/Sub message.
// https://api.slack.com/apps/A01D75SGC30/incoming-webhooks?success=1
func HelloPubSub(ctx context.Context, m PubSubMessage) error {
	var opts Environment

	_, err := env.UnmarshalFromEnviron(&opts)
	if err != nil {
		return fmt.Errorf("cannot get envs: %w", err)
	}

	build := Build{}

	err = json.Unmarshal(m.Data, &build)
	if err != nil {
		return fmt.Errorf("cannot unmarshal build: %w", err)
	}

	mailService := NewMailService(
		opts.MailConf.Enabled,
		opts.MailConf.Domain,
		opts.MailConf.Sender,
		opts.MailConf.APIKey,
		build,
		opts.MailConf.FilterBuildStatuses,
	)

	slackService := NewSlackService(
		opts.SlackConf.Enabled,
		opts.SlackConf.Webhook,
		build,
		opts.SlackConf.FilterBuildStatuses,
	)

	messageService := NewMessageService(build)

	err = mailService.Send(
		strings.Split(opts.MailConf.Recipients, ","),
		fmt.Sprintf("Cloud Build %s", build.ID),
		messageService.Mail(),
	)
	if err != nil {
		log.Println("error during sending email: %w", err)
	}

	err = slackService.SendSlackNotification(messageService.Slack())
	if err != nil {
		log.Println("error during sending email: %w", err)

	}

	return err
}
