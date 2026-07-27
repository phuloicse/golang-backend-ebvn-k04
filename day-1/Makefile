.PHONY: run swagger test

run:
	go run cmd/api/main.go

swagger:
	swag init -g cmd/api/main.go -oupput docs

COVERAGE_EXCLUDE=main.go|test|mocks

test:
	go test ./... -coverprofile=coverage.tmp -coverpkg=./... -covermode=atomic -p 1
	grep -vE "$(COVERAGE_EXCLUDE)" coverage.tmp > coverage.out
	go tool cover -html=coverage.out -o coverage.html