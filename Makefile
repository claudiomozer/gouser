.PHONY: setup
setup:
	go install go.uber.org/mock/mockgen@latest
	go install github.com/jackc/tern/v2@latest	
	go mod tidy

.PHONY: mocks
mocks:
	rm -rf ./mocks
	go generate ./...

.PHONY: migrations
migrations:
	tern migrate --migrations=./migrations

.PHONY: run
run: 
	go run ./cmd/main.go