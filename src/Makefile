.DEFAULT_GOAL := build

fmt:
	go fmt TimeServer.go
.PHONY:fmt

lint: fmt
	golint TimeServer.go
.PHONY:lint

vet: fmt
	go vet TimeServer.go
.PHONY:vet

build: vet
	go build TimeServer.go
.PHONY:build