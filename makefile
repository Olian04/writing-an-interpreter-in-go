default: test

t: test
test:
	go test -v ./...

l: lint
lint:
	go vet ./...

i: install
install:
	go mod tidy
	go mod verify
	go mod download

