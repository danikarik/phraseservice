PACKAGES=$(shell go list ./... | grep -v '/simulation')

VERSION := $(shell git describe --tags --long 2>/dev/null || git rev-parse --short HEAD)
COMMIT := $(shell git log -1 --format='%H')

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=phraseservice \
	-X github.com/cosmos/cosmos-sdk/version.ServerName=phraseserviced \
	-X github.com/cosmos/cosmos-sdk/version.ClientName=phraseservicecli \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT)

BUILD_FLAGS := -ldflags '$(ldflags)'

all: prep build

prep:
	@rm -rf build
	@mkdir build

build:
	@echo "--> Building phraseserviced & phraseservicecli"
	@go build $(BUILD_FLAGS) -o build/phraseserviced ./cmd/phraseserviced
	@go build $(BUILD_FLAGS) -o build/phraseservicecli ./cmd/phraseservicecli

build-linux:
	@GOOS=linux GOARCH=amd64 $(MAKE) build

test:
	@go test $(PACKAGES)

image:
	@DOCKER_BUILDKIT=1 docker build -t phraseservice .
