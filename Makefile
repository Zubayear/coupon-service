GOPATH:=$(shell go env GOPATH)

.PHONY: init
init:
	@go get -u google.golang.org/protobuf/proto
	@go install github.com/golang/protobuf/protoc-gen-go@latest
	@go install github.com/asim/go-micro/cmd/protoc-gen-micro/v4@latest

.PHONY: proto
proto:
	@protoc --proto_path=. --micro_out=. --go_out=:. proto/Coupon.proto

.PHONY: update
update:
	@go get -u

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: build
build:
	@go build -o Coupon *.go

.PHONY: test
test:
	@go test -v ./... -cover

.PHONY: docker
docker:
	@docker build -t Coupon:latest .

.PHONY: grpc-gateway
grpc-gateway:
	@protoc --proto_path=./proto --micro_out=. --grpc-gateway_out=logtostderr=true,register_func_suffix=UT:. --openapiv2_out=./proto --openapiv2_opt=logtostderr=true --openapiv2_opt=use_go_templates=true --go_out=plugins=grpc:. ./proto/Coupon.proto

.PHONY: run_service
DEFAULT_PORT=60009
run_service:
	@go run main.go --server_address=localhost:$(DEFAULT_PORT)

.PHONY: run-gateway
run-gateway:
	@go run gateway/gateway.go

.PHONY: di
di:
	@wire di/wire.go