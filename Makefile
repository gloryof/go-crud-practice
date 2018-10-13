GOCMD=go
GOBUILD="$(GOCMD) build"
GOCLEAN="$(GOCMD) clean"
GOTEST="$(GOCMD) test"
BINARY_NAME=crud
BINARY_DIR=bin

all: dependency clean test build assets
dependency:
	dep ensure
build:
	./make/script/build.sh ${GOBUILD} ${BINARY_DIR} ${BINARY_NAME}
test:
	./make/script/test.sh ${GOTEST} ${BINARY_DIR}
assets:
	./make/script/assets.sh ${BINARY_DIR}
clean:
	./make/script/clean.sh ${GOCLEAN} ${BINARY_DIR}
