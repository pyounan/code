package p

import (
	"fmt"
)

type MessageService struct {
	build Build
}

func NewMessageService(build Build) *MessageService {
	return &MessageService{
		build: build,
	}
}

func Includebackquote(input string) string {
	return fmt.Sprintf("`%s`", input)
}

// Slack formats the cloudbuild message to slack messsage
func (s *MessageService) Slack() SlackMessage {
	duration := s.build.FinishTime.Sub(s.build.StartTime)
	stdin := `
Build: %s

					Repo: %s

					Branch: %s

					Duration: %s
`
	mainText := fmt.Sprintf(
		stdin,
		Includebackquote(s.build.BuildTriggerID),
		Includebackquote(s.build.Substitutions.REPONAME),
		Includebackquote(s.build.Substitutions.BRANCHNAME),
		Includebackquote(duration.String()),
	)
	msg := SlackMessage{
		Text: mainText,
		Attachments: []SlackAttachment{
			{
				Color:     SlackColorMap[s.build.Status],
				Title:     "Build Logs",
				TitleLink: s.build.LogURL,
				Fields: []AttachmentField{
					{
						Title: "Status",
						Value: s.build.Status,
						Short: false,
					},
				},
			},
		},
	}

	return msg
}

// Mail formats the cloudbuild to mail templ
func (s *MessageService) Mail() string {
	duration := s.build.FinishTime.Sub(s.build.StartTime)

	ret := fmt.Sprintf(`
	RepoName : %s
	BranchName : %s
	BuildStatus : %s
	BuildDuration : %s
	BuildLogs : %s
	`, s.build.Substitutions.REPONAME, s.build.Substitutions.BRANCHNAME, s.build.Status, duration.String(), s.build.LogURL)
	return ret
}
