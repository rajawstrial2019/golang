clean:
	go clean
build:
	go build httpd/newsfeed_service.go
install:
	go install httpd/newsfeed_service.go
run:
	go run httpd/newsfeed_service.go
