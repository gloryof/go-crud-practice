#!/bin/sh

# Parametes
#  [1]: GOCLEAN
#  [2]: BINARY_DIR

GOCLEAN=$1
BINARY_DIR=$2

${GOCLEAN}
if [ -e ${BINARY_DIR} ]; then
    rm -rf ${BINARY_DIR}
fi