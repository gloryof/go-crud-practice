#!/bin/sh

# Parametes
#  [1]: GOBUILD
#  [2]: BINARY_DIR
#  [3]: BINARY_NAME

GOBUILD=$1
BINARY_DIR=$2
BINARY_NAME=$3

if [ ! -e ${BINARY_DIR} ]; then
    mkdir ${BINARY_DIR}
fi

echo ${GOBUILD}
echo ${BINARY_DIR}
echo ${BINARY_NAME}


GOOS=linux GOARCH=amd64 CGO_ENABLED=0 ${GOBUILD} -tags netgo -installsuffix netgo  -o ${BINARY_DIR}/${BINARY_NAME} -v