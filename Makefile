.PHONY: all
all: clean lint test redundantimportalias

redundantimportalias: redundantimportalias.go
	@echo ">> Compiling..."
	go build $<

.PHONY: clean
clean:
	@echo ">> Cleaning up..."
	rm -f redundantimportalias

.PHONY: deps
deps:
	@echo ">> Resolving dependencies..."
	go mod download
	go mod tidy

.PHONY: deps-update
deps-update:
	@echo ">> Updating dependencies..."
	go get -d -u -t ./...
	@make deps

.PHONY: tools
tools: tools
	@echo ">> Installing Go tools..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.60.1

.PHONY: lint
lint:
	@echo ">> Linting source code..."
	golangci-lint run ./...

.PHONY: test
test:
	@echo ">> Testing all..."
	go test -race ./...
