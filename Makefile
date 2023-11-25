PHONY: build
build: gen-go
	go build pubsub.go
	go build ./pkg/...

PHONY: examples
examples: gen-go
	go build -o ./bin/publisher ./examples/basic/publisher
	go build -o ./bin/subscriber ./examples/basic/subscriber

PHONY: gen-go
gen-go:
	protoc --go_out=./pkg/proto --go_opt=paths=source_relative \
           --go-grpc_out=./pkg/proto  --go-grpc_opt=paths=source_relative \
           -I./pkg/proto ./pkg/proto/*.proto

PHONY: dep
dep:
	go get ./... && go mod tidy
