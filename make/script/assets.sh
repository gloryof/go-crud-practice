#!/bin/sh

# Parametes
#  [1]: BINARY_DIR
#  [3]: BINARY_NAME

BINARY_DIR=$1
BINARY_NAME=$3

if [ ! -e ${BINARY_DIR} ]; then
    mkdir ${BINARY_DIR}
fi

tar -czvf ${BINARY_DIR}/crud-assets.tar.gz static