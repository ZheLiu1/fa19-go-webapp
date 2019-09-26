default: build

build:
	go get -d ./...
	go build -o bin/webapp cmd/api/api.go

image:
	go get -d ./...
	CGO_ENABLED=0 GOOS=linux go build -a -o bin/webapp cmd/api/api.go

docker:
	sudo docker build -t webapp .
