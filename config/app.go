package config

type AppConfig struct {
	Port               int `envconfig:"PORT" default:"8080"`
	MaxJobAttemptCount int `envconfig:"MAX_JOB_ATTEMPT_COUNT" default:"3"`
	MaxJobRetryCount   int `envconfig:"MAX_JOB_RETRY_COUNT" default:"3"`
}
