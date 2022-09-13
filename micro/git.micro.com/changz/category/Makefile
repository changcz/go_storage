GOPATH:=$(shell go env GOPATH)
.PHONY: proto
proto:
	protoc  -I ./ --go_out=./ --micro_out=./ ./proto/category/category.proto
.PHONY: build
build:
	 CGO_ENABLE=0 GOOS=linux GOARCH=amd64  go build -o category *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build -t category-service:latest .
