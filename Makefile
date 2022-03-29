ROOT_PATH := $(PWD)
BIN_PATH := bin

.PHONY: \
	all \
	run \
	run_dev \
	devrun \
	build-faf-content-service \
	tools \
	imports \
	vendor

all: build-faf-content-service

run:
	$(BIN_PATH)/faf-content-service -config=configs/config.yaml

run_dev:
	$(BIN_PATH)/faf-content-service -config=configs/config.yaml

devrun: build-faf-content-service
	$(BIN_PATH)/faf-content-service -config=configs/config.yaml -addr=127.0.0.1

build-faf-content-service:
	go build -o $(BIN_PATH)/faf-content-service ./cmd/

tools:
	go mod download golang.org/x/tools
	go install golang.org/x/tools/cmd/goimports@latest

imports:
	goimports -l -w .

vendor:
	go mod tidy && go mod vendor
