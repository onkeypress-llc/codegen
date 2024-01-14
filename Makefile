.PHONY: test
test:
	go test ./...

.PHONY: coverage
coverage:
	go test ./... -coverprofile=cover.out

.PHONY: coverage-report
coverage-report: coverage
	go tool cover -html=cover.out

.PHONY: fmt
fmt:
	# Fixup modules
	go mod tidy
	# Format the Go sources:
	go fmt ./...
build:
	go build -o gencode
pub: build
	cp ./gencode $$GOPATH/bin/gencode