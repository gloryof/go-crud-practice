#!/bin/sh

# Parametes
#  [1]: GOTEST
#  [2]: BINARY_DIR

GOTEST=$1
BINARY_DIR=$2

if [ ! -e ${BINARY_DIR} ]; then
    mkdir ${BINARY_DIR}
fi

${GOTEST} -coverprofile=${BINARY_DIR}/cover.out -covermode count ./crud/...