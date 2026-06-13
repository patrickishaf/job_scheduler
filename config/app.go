package config

type AppConfig struct {
	Port               int    `envconfig:"PORT" default:"8080"`
	MaxDLQSize         int    `envconfig:"MAX_DLQ_SIZE" default:"10"`
	MaxJobAttemptCount int    `envconfig:"MAX_JOB_ATTEMPT_COUNT" default:"4"`
	MaxJobRetryCount   int    `envconfig:"MAX_JOB_RETRY_COUNT" default:"3"`
	QueueType          string `envconfig:"QUEUE_TYPE" default:"heap"`
}
