build-binary:
	GO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o captcha-service .
