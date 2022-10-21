# Build

> GOOS=linux GOARCH=arm go build -o build/sms-sqs cmd/main.go

# Testing

To send test message use the next command line:
> AWS_PAGER="" aws sqs send-message --endpoint-url 'https://message-queue.api.cloud.yandex.net' --queue-url 'https://message-queue.api.cloud.yandex.net/<queue_id>/<queue_name>' --message-body "{\"recipient\": \"79262471221\", \"message\": \"Привет, сообщение №$(uuidgen)\"}"
