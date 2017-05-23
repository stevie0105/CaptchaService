build-binary: go-get
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o captcha-service .

go-get:
	go get
