FROM scratch
ADD captcha-service /
EXPOSE 8666:8666
CMD ["/captcha-service"]
