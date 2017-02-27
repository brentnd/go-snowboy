# Name of the binary output
BINARY=detect
BUILD_DIR=build

CXX=clang++
CC=clang

.PHONY: all

all:
	CXX=${CXX} CC=${CC} go build -o ${BUILD_DIR}/${BINARY} example/main.go

clean:
	rm -rf ${BUILD_DIR}/*