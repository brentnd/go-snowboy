BUILD_DIR=build

CXX=clang++
CC=clang

.PHONY: all clean cmd api

all: cmd api

cmd:
	CXX=${CXX} CC=${CC} go build -o ${BUILD_DIR}/snowboy-detect example/detect.go

api:
	CXX=${CXX} CC=${CC} go build -o ${BUILD_DIR}/snowboy-api example/api.go

clean:
	rm -rf ${BUILD_DIR}/*