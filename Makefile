BUILD_DIR=build

CXX=clang++
CC=clang

.PHONY: all cmd fixed clean

all: cmd fixed

cmd:
	CXX=${CXX} CC=${CC} go build -o ${BUILD_DIR}/detect-cmd example/cmd.go

fixed:
	cp $$GOPATH/src/github.com/Kitt-AI/snowboy/resources/alexa.umdl ${BUILD_DIR}/
	cp $$GOPATH/src/github.com/Kitt-AI/snowboy/resources/common.res ${BUILD_DIR}/
	CXX=${CXX} CC=${CC} go build -o ${BUILD_DIR}/detect-fixed example/fixed.go

clean:
	rm -rf ${BUILD_DIR}/*