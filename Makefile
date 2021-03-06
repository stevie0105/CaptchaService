run: binary-build container-build
	docker run -d -p 8666:8666 --name captcha-service captcha-service
	rm captcha-service

container-build:
	docker build -t captcha-service .

binary-build: go-get
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o captcha-service .

go-get:
	go get

complete-remove:
	docker stop captcha-service && docker rm captcha-service && docker rmi captcha-service
