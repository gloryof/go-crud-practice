GOCMD=go
GOBUILD="$(GOCMD) build"
GOCLEAN="$(GOCMD) clean"
GOTEST="$(GOCMD) test"
GOGET="$(GOCMD) get"
DEP_VERSION=0.4.1
DEP_PATH=${GOPATH}/bin/dep
BINARY_NAME=crud
BINARY_DIR=bin

all: dependency clean test build
dependency:
	./script/dependency.sh ${DEP_PATH}
build:
	./script/build.sh ${GOBUILD} ${BINARY_DIR} ${BINARY_NAME}
test:
	./script/test.sh ${GOTEST}
clean:
	./script/clean.sh ${GOCLEAN} ${BINARY_DIR}
