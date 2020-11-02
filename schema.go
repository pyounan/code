package p

import "time"

// PubSubMessage is the payload of a Pub/Sub event.
type PubSubMessage struct {
	Data []byte `json:"data"`
}

// Environment defines the envs for main function
type Environment struct {
	SlackConf struct {
		Enabled             bool   `env:"SLACK_ENABLED"`
		FilterBuildStatuses string `env:"SLACK_FILTER_BUILD_STATUSES"`
		Webhook             string `env:"SLACK_WEBHOOK_URL"`
	}
	MailConf struct {
		Enabled             bool   `env:"MAIL_ENABLED"`
		FilterBuildStatuses string `env:"SLACK_FILTER_BUILD_STATUSES"`
		Domain              string `env:"MAIL_DOMAIN"`
		Sender              string `env:"MAIL_SENDER"`
		APIKey              string `env:"MAIL_APIKEY"`
		Recipients          string `env:"RECIPIENTS"`
	}
}

type SlackMessage struct {
	Text        string            `json:"text"`
	Attachments []SlackAttachment `json:"attachments"`
}

type SlackAttachment struct {
	Color     string            `json:"color"`
	Title     string            `json:"title"`
	TitleLink string            `json:"title_link"`
	Fields    []AttachmentField `json:"fields"`
}

type AttachmentField struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

type Build struct {
	Name      string `json:"name"`
	ID        string `json:"id"`
	ProjectID string `json:"projectId"`
	Status    string `json:"status"`
	Source    struct {
		StorageSource struct {
			Bucket string `json:"bucket"`
			Object string `json:"object"`
		} `json:"storageSource"`
	} `json:"source"`
	Steps []struct {
		Name       string   `json:"name"`
		Args       []string `json:"args"`
		ID         string   `json:"id"`
		Entrypoint string   `json:"entrypoint"`
		Volumes    []struct {
			Name string `json:"name"`
			Path string `json:"path"`
		} `json:"volumes"`
		Timing struct {
			StartTime time.Time `json:"startTime"`
			EndTime   time.Time `json:"endTime"`
		} `json:"timing"`
		PullTiming struct {
			StartTime time.Time `json:"startTime"`
			EndTime   time.Time `json:"endTime"`
		} `json:"pullTiming"`
		Status string `json:"status"`
		Dir    string `json:"dir,omitempty"`
	} `json:"steps"`
	Results struct {
		BuildStepImages  []string `json:"buildStepImages"`
		BuildStepOutputs []string `json:"buildStepOutputs"`
	} `json:"results"`
	CreateTime     time.Time `json:"createTime"`
	StartTime      time.Time `json:"startTime"`
	FinishTime     time.Time `json:"finishTime"`
	Timeout        string    `json:"timeout"`
	QueueTTL       string    `json:"queueTtl"`
	LogsBucket     string    `json:"logsBucket"`
	BuildTriggerID string    `json:"buildTriggerId"`
	Options        struct {
		MachineType          string `json:"machineType"`
		SubstitutionOption   string `json:"substitutionOption"`
		DynamicSubstitutions bool   `json:"dynamicSubstitutions"`
		Logging              string `json:"logging"`
	} `json:"options"`
	LogURL        string `json:"logUrl"`
	Substitutions struct {
		BRANCHNAME  string `json:"BRANCH_NAME"`
		COMMITSHA   string `json:"COMMIT_SHA"`
		REPONAME    string `json:"REPO_NAME"`
		REVISIONID  string `json:"REVISION_ID"`
		SHORTSHA    string `json:"SHORT_SHA"`
		SERVICEDIR  string `json:"_SERVICE_DIR"`
		SERVICENAME string `json:"_SERVICE_NAME"`
	} `json:"substitutions"`
	Tags   []string `json:"tags"`
	Timing struct {
		BUILD struct {
			StartTime time.Time `json:"startTime"`
			EndTime   time.Time `json:"endTime"`
		} `json:"BUILD"`
		FETCHSOURCE struct {
			StartTime time.Time `json:"startTime"`
			EndTime   time.Time `json:"endTime"`
		} `json:"FETCHSOURCE"`
	} `json:"timing"`
}
