package internal

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"os"
)

type SMS struct {
	logger *logrus.Logger
	ctx    context.Context
}

func InitSMS(logger *logrus.Logger) *SMS {
	ctx := context.Background()
	return &SMS{
		logger: logger,
		ctx:    ctx,
	}
}

func (self *SMS) SendSMS(message *SQSMessageSMSRequest) {
	var destination string = fmt.Sprintf("/var/spool/sms/outgoing/%s", uuid.New().String())

	f, err := os.Create(destination)
	defer f.Close()

	if err != nil {
		self.logger.Errorln(err)
		return
	}

	f.WriteString(fmt.Sprintf("To: %s\nAlphabet: UTF-8\n\n%s", message.Recipient, message.Message))
	f.Sync()

	self.logger.Printf("Message %s was sent to %s", message.Recipient, message.Message)
}
