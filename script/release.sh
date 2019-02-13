#!/usr/bin/env bash

XC_ARCH=${XC_ARCH:-amd64}
XC_OS=${XC_OS:-linux darwin windows}
VERSION=${VERSION:-v0.0.0}
BUILD_FLAGS="-X \"main.Version=${VERSION}\""
ASSET_DIR="${ASSET_DIR:=pkg}"

rm -rf pkg/*
gox \
    -ldflags="${BUILD_FLAGS}" \
    -os="${XC_OS}" \
    -arch="${XC_ARCH}" \
    -output "${ASSET_DIR}/{{.OS}}-{{.Arch}}/{{.Dir}}"
