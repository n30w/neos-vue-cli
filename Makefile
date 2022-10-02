BINARY_NAME=neos-vue-cli
THIS_DIR := $(dir $(abspath $(firstword $(MAKEFILE_LIST))))

all: build run

build:
	# GOARCH=arm64 GOOS=darwin go build -ldflags="-s -w" ${THIS_DIR}.
	# GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o ${BINARY_NAME}-linux main.go
	# GOARCH=amd64 GOOS=window go build -ldflags="-s -w" -o ${BINARY_NAME}-windows main.go

run:
	${THIS_DIR}./${BINARY_NAME} -v

clean:
	go clean
	rm ${THIS_DIR}${BINARY_NAME}
	rm ${BINARY_NAME}-linux
	rm ${BINARY_NAME}-windows
	rm -rf ${THIS_DIR}Default
