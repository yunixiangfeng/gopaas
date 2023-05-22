
.PHONY: proto
proto:
	sudo yu-v3 --proto_path=. --micro_out=. --go_out=:. ./proto/base/base.proto

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o base *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t base:latest
