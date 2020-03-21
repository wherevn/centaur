
prepare:
	@bash ./scripts/prepare.sh

dep:
	@go mod download
	@go mod vendor

clean:
	@bash ./scripts/clean.sh

lint:
	@golangci-lint run
