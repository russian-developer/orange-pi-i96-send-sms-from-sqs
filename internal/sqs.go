package internal

import (
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type SQS struct {
	SQS_URL       string
	SQS_QUEUE_URL string
	SQS_REGION    string
	logger        *logrus.Logger
	sms           *SMS
	ctx           context.Context
}

func InitSQS(logger *logrus.Logger, sms *SMS) *SQS {
	ctx := context.Background()
	return &SQS{
		logger:        logger,
		ctx:           ctx,
		sms:           sms,
		SQS_URL:       os.Getenv("SQS_URL"),
		SQS_QUEUE_URL: os.Getenv("SQS_QUEUE_URL"),
		SQS_REGION:    os.Getenv("SQS_REGION"),
	}
}

func (self *SQS) getClient() *sqs.Client {
	resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL:           self.SQS_URL,
			SigningRegion: self.SQS_REGION,
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(
		self.ctx,
		config.WithEndpointResolverWithOptions(resolver),
	)
	if err != nil {
		self.logger.Fatalln(err)
	}

	return sqs.NewFromConfig(cfg)
}

func (self *SQS) Run() {
	self.logger.Println("Run SQS")

	client := self.getClient()

	received, err := client.ReceiveMessage(self.ctx, &sqs.ReceiveMessageInput{
		QueueUrl:            &self.SQS_QUEUE_URL,
		MaxNumberOfMessages: 10,
		WaitTimeSeconds:     1,
		VisibilityTimeout:   1,
	})
	if err != nil {
		self.logger.Println(err)
		return
	}

	for _, v := range received.Messages {
		self.logger.Printf("Message received\nID: %s\nBody: %s\n", *v.MessageId, *v.Body)

		if _, err := client.DeleteMessage(
			self.ctx,
			&sqs.DeleteMessageInput{
				QueueUrl:      &self.SQS_QUEUE_URL,
				ReceiptHandle: v.ReceiptHandle,
			},
		); err != nil {
			self.logger.Println(err)
			return
		}

		var message SQSMessageSMSRequest

		err := json.Unmarshal([]byte(*v.Body), &message)
		if err != nil {
			self.logger.Printf("Cannot unmarshall SQS message detail data. %s", err)
			continue
		}

		self.logger.Println("Send SMS")
		go self.sms.SendSMS(&message)
	}
}
