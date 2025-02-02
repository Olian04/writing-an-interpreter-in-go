default: run

# Build and run tests
r: run
run:
	go test -v ./test/...

# Build, test and benchmark project
b: benchmark
benchmark:
	go test -v -bench=. ./test/...

# Install dependencies
i: install
install:
	go mod download
	go mod tidy

# Formatting & linting
l: lint
lint:
	go fmt ./...
	go vet ./...
