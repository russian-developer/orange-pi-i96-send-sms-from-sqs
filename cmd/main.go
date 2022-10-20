package main

import (
	"github.com/go-co-op/gocron"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
	"github.com/unk2k/orange-pi-i96-send-sms-from-sqs/internal"
	"time"
)

var Logger *logrus.Logger = createLogger()

func main() {
	Logger.Println("Run every 5 seconds")
	scheduler := gocron.NewScheduler(time.UTC)
	_, err := scheduler.Every(5).SingletonMode().Do(fire)
	if err != nil {
		Logger.Fatalf("Cannot start periodic task. %s", err)
	}
	scheduler.StartBlocking()
}

func createLogger() *logrus.Logger {
	logger := logrus.New()

	/*
		TODO: Send logs to yandex cloud logging
	*/
	return logger
}

func fire() {
	sms := internal.InitSMS(Logger)
	sqs := internal.InitSQS(Logger, sms)
	sqs.Run()
}
