BUILD_VERSION ?= "unknown"

OUTPUT_NAME := "dollar-tool"

clean:
	@rm -rf build/

build: clean
	@GOOS=windows GOARCH=amd64 go build -o ./build/$(OUTPUT_NAME)-windows-amd64.exe -ldflags "-X github.com/matthiasharzer/dollar-tool/cmd/version.version=$(BUILD_VERSION)" ./main.go
	@GOOS=linux GOARCH=amd64 go build -o ./build/$(OUTPUT_NAME)-linux-amd64 -ldflags "-X github.com/matthiasharzer/dollar-tool/cmd/version.version=$(BUILD_VERSION)" ./main.go
	@GOOS=darwin GOARCH=amd64 go build -o ./build/$(OUTPUT_NAME)-darwin-amd64 -ldflags "-X github.com/matthiasharzer/dollar-tool/cmd/version.version=$(BUILD_VERSION)" ./main.go
	@GOOS=darwin GOARCH=arm64 go build -o ./build/$(OUTPUT_NAME)-darwin-arm64 -ldflags "-X github.com/matthiasharzer/dollar-tool/cmd/version.version=$(BUILD_VERSION)" ./main.go

qa: analyze test

analyze:
	@go vet
	@go tool staticcheck --checks=all

test:
	@go test -failfast -cover ./...


.PHONY: clean \
				build \
				qa \
				analyze \
				test
