default: run

# Build and run project
r: run
run:
	go test -v ./test/...

b: benchmark
benchmark:
	go test -v -bench=. ./test/...

# Install dependencies
i: install
install:
	go mod download
	go mod tidy

# formatting & linting
l: lint
lint:
	go fmt ./**/*.go
	go vet ./**/*.go
