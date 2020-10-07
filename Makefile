all: build

build:
	@mkdir -p build/
	@go build -o build/phraseserviced ./cmd/phraseserviced
	@go build -o build/phraseservicecli ./cmd/phraseservicecli
