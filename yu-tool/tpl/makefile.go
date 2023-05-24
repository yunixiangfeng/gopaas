package template

var (
	Makefile = `
.PHONY: proto
proto:
	yu-v3 --proto_path=. --micro_out=. --go_out=:. ./proto/{{.Alias}}/{{.Alias}}.proto

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o {{.Alias}} *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t {{.Alias}}:latest
`

//	GenerateFile = `package main
////go:generate make proto
//`
)
