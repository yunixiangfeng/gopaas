
.PHONY: proto
proto:
	yu-v3 --proto_path=. --micro_out=. --go_out=:. ./proto/volumeApi/volumeApi.proto

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o volumeApi *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t volumeApi:latest
