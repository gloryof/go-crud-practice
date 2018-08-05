#!/bin/sh

# Parameters
#  [1]: DEP_PATH

DEP_PATH=$1

if [ ! -e  ${DEP_PATH} ]; then
    curl -fsSL -o DEP_PATH https://github.com/golang/dep/releases/download/v${DEP_VERSION}/dep-linux-amd64
    chmod +x DEP_PATH
fi

dep ensure