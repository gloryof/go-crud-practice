#!/bin/sh

# Parametes
#  [1]: GOTEST
#  [2]: BINARY_DIR

GOTEST=$1
BINARY_DIR=$2
TARGET_DIR=./crud/...

if [ ! -e ${BINARY_DIR} ]; then
    mkdir ${BINARY_DIR}
fi


ENV_VARIABLES=GOOS=linux GOARCH=amd64 CGO_ENABLED=0 ${GOTEST} -coverprofile=${BINARY_DIR}/cover.out -coverpkg=${TARGET_DIR} ${TARGET_DIR}