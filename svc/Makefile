
.PHONY: proto
proto:
	yu-v3 --proto_path=. --micro_out=. --go_out=:. ./proto/svc/svc.proto

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o svc *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t svc:latest
