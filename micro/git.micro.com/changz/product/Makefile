
GOPATH:=$(shell go env GOPATH)
.PHONY: init
init:
	go install github.com/micro/go-plugins/registry/consul/v2
	go install github.com/micro/micro/v3/cmd/protoc-gen-micro@latest
	go install github.com/micro/micro/v3/cmd/protoc-gen-openapi@latest

.PHONY: api
api:
	protoc --openapi_out=. --proto_path=. proto/git.micro.com/changz/pruduct.proto

.PHONY: proto
proto:
		protoc  -I ./ --go_out=./ --micro_out=./ ./proto/product/product.proto
	
.PHONY: build
build:
	go build -o git.micro.com/changz/pruduct *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t git.micro.com/changz/pruduct:latest
