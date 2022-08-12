.PHONY: all clean

SOURCES = $(shell find . -name '*.go')

all: terrafy
terrafy: ${SOURCES}
	go build -o terrafy cmd/terrafy/main.go
clean:
	rm -rf terrafy
