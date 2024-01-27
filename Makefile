.PHONY: tidy initdb

tidy:
	@go mod tidy

initdb:
	@go run ./cmd/initdb

.DEFAULT_GOAL := tidy