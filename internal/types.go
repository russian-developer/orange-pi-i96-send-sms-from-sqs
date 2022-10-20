package internal

type SQSMessageSMSRequest struct {
	Recipient string `json:"recipient"`
	Message   string `json:"message"`
}
