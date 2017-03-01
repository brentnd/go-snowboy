BUILD_DIR=build

CXX=clang++
CC=clang

.PHONY: all clean

all:
	CXX=${CXX} CC=${CC} go build -o ${BUILD_DIR}/detect-cmd example/cmd.go

clean:
	rm -rf ${BUILD_DIR}/*