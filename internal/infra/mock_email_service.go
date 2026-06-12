package infra

type SendEmailPayload struct {
	ToEmailAddress   string
	FromEmailAddress string
	Subject          string
	Title            string
	Body             string
}

func SendMockEmail(payload *SendEmailPayload) error {
	return nil
}
