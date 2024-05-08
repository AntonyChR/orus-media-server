#!/bin/bash

echo "Checking dependencies"

node_version=$(node --version 2>/dev/null)

if [[ -z $node_version ]]; then
    echo "Node is not installed"
    exit 1
fi

node_version_major="$(echo $node_version | cut -d "." -f 1 | cut -d "v" -f 2)"

if [[ $node_version_major -lt 18 ]]; then
    echo "Node version 18.0.0 or higher is required"
    exit 1
fi

golang_version=$(go version 2>/dev/null)

if [[ -z $golang_version ]]; then
    echo "Go is not installed"
    exit 1
fi

golang_version_major="$(echo $golang_version | cut -d " " -f 3 | cut -d "." -f 2)"

if [[ $golang_version_major -lt 20 ]]; then
    echo "Go version 1.20 or higher is required"
    exit 1
fi
