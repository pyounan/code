package p

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

// SlackColorMap maps
var SlackColorMap = map[string]string{
	"SUCCESS": "#36a64f",
}

type SlackService struct {
	build   Build
	Enabled bool
	Webhook string
	Filter  string
}

// NewSlackService construct a new slack service
func NewSlackService(enabled bool, webhook string, build Build, filter string) *SlackService {
	return &SlackService{
		build:   build,
		Enabled: enabled,
		Webhook: webhook,
		Filter:  filter,
	}
}

// SendSlackNotification will post to an 'Incoming Webook' url setup in Slack Apps. It accepts
// some text and the slack channel is saved within Slack.
func (s *SlackService) SendSlackNotification(msg SlackMessage) error {
	if s.Enabled == false {
		log.Println("Sending notification to slack is disabled. Skipping .. ")
		return nil
	}

	if s.Filter != "" && !strings.Contains(s.Filter, s.build.Status) {
		log.Printf("Sending notification to slack isnot enabled at this status [%s]. Skipping .. \n", s.build.Status)
		return nil
	}

	slackBody, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("cannot marshal slack msg: %v", err)
	}
	req, err := http.NewRequest(http.MethodPost, s.Webhook, bytes.NewBuffer(slackBody))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	if buf.String() != "ok" {
		return errors.New("Non-ok response returned from Slack")
	}
	return nil
}
