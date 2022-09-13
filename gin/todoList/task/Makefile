
GOPATH:=$(shell go env GOPATH)

.PHONY: proto
proto:
	protoc -I internal/service/pb internal/service/pb/*.proto --go_out=plugins=grpc:.
	#protoc -I --proto_path=. --micro_out=${MODIFY}:. --go_out=${MODIFY}:. proto/cartApi/cartApi.proto

.PHONY: build
build: proto

	go build -o cart-service *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t cart-service:latest
