setup:
	go get -u github.com/golang/lint/golint

clean:
	go clean


prebuild:
	go fmt ./...
	go vet ./...
#	golint ./...

build: prebuild
	go install -tags pprof

.DEFAULT_GOAL := build

.PHONY: setup clean prebuild build
