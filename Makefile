.PHONY: test

install:
	go install -mod=readonly -v ./internal/...

mod:
	go mod vendor
	go mod tidy

run:
	go run internal/main.go
