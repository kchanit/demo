mock:
	rm -rf internal/core/port/mocks && \
	mockery --dir=internal/core/port --all --output=internal/core/port/mocks && \
	rm -rf internal/core/port/mocks/Service.go && \
	go mod tidy

dev:
	set -a && \
		source configs/local.env && \
	set +a && \
	nodemon --exec go run --tags dynamic $(shell pwd)/cmd/server/main.go --signal SIGTERM
