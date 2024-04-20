.PHONY: build
build:
	env GOOS=linux GOARCH=amd64 \
	go build -o bin/mock-ups-ctrl cmd/mock-ups-ctrl/main.go

.PHONY: test
test:
	go test -race -timeout 30s ./...

.DEFAULT_GOAL :=build
