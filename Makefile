BUILD_VERSION=dev-snapshot

all: build

build:
	go build -ldflags="-X main.BuildVersion=$(BUILD_VERSION)" -o bin/stratus cmd/stratus/*.go

test:
	go test ./... -v

mocks:
	mockery --name=StateManager --dir internal/state --output internal/state/mocks
	mockery --name=TerraformManager --dir internal/runner --output internal/runner/mocks
	mockery --name=FileSystem --structname FileSystemMock --dir internal/state --output internal/state/mocks