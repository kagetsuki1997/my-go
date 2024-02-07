doge:
	go run cmd/doge/main.go

lint:
	golangci-lint run

PROTO_FILES := $(shell find . -type f -name "*.proto")

proto: $(patsubst %.proto,%.pb.go,$(PROTO_FILES))

%.pb.go: %.proto
	protoc --proto_path=. \
	--go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative $<
