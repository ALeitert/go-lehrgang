PROJ_DIR=$(shell pwd)

gen-protobuf:
	protoc \
		--plugin=protoc-gen-go=$(PROJ_DIR)/bin/protoc-gen-go \
		--plugin=protoc-gen-go-grpc=$(PROJ_DIR)/bin/protoc-gen-go-grpc \
		\
		--go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		\
		proto/echo.proto

install-go-protoc:
	GOBIN=$(PROJ_DIR)/bin go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	GOBIN=$(PROJ_DIR)/bin go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
