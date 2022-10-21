run:
	GOOS=linux GOARCH=arm go build -o board/bin/sms-sqs cmd/main.go
	tar -zcvf archive.tar.gz -C board .