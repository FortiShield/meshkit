include build/Makefile.core.mk
include build/Makefile.show-help.mk

check:
	golangci-lint run -c .golangci.yml -v ./... --fix

test:
	go test --short ./... -race -coverprofile=coverage.txt -covermode=atomic

tidy:
	go mod tidy
	git diff --exit-code go.mod go.sum

errorutil:
	go run github.com/khulnasoft/meshkit/cmd/errorutil -d . update --skip-dirs meshplay -i ./helpers -o ./helpers

errorutil-analyze:
	go run github.com/khulnasoft/meshkit/cmd/errorutil -d . analyze --skip-dirs meshplay -i ./helpers -o ./helpers

build-errorutil:
	go build -o errorutil cmd/errorutil/main.go
