package template

var (
	Makefile = `
GOPATH:=$(shell go env GOPATH)

.PHONY: proto test docker

{{if ne .Type "web"}}
proto:
	# protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/example/example.proto
	ironic proto --path "./proto"

build: proto
{{else}}
build:
{{end}}
	GOOS=linux GOARCH=amd64 go build -o {{.Alias}}-{{.Type}} main.go plugin.go
	go build -o {{.Alias}}-{{.Type}}-local main.go plugin.go

test:
	go test -v ./... -cover

docker:
	docker build . -t {{.Alias}}-{{.Type}}:latest
`
)
