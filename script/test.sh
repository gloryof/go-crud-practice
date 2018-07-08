#!/bin/sh

# Parametes
#  [1]: GOTEST

GOTEST=$1

${GOTEST} -cover ./crud/...