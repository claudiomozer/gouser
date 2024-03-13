.PHONY: setup
setup:
	go install go.uber.org/mock/mockgen@latest
	go install github.com/jackc/tern/v2@latest	
	go mod tidy

.PHONY: mocks
mocks:
	rm -rf ./mocks
	go generate ./...

.PHONY: load_environment
load_environment:
	export $(< .env)

.PHONY: migrations
migrations:	load_environment
	tern migrate --migrations=./migrations

.PHONY: run
run: load_environment
	go run ./cmd/main.go