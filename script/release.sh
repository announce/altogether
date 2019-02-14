#!/usr/bin/env bash

XC_ARCH=${XC_ARCH:-amd64}
XC_OS=${XC_OS:-linux darwin}
VERSION=${VERSION:-v0.0.0}
BUILD_FLAGS="-X \"main.Version=${VERSION}\""
ASSET_DIR="${ASSET_DIR:=pkg}"
echo "ASSET_DIR:${ASSET_DIR}"

#rm -rf pkg/*
gox \
    -ldflags="${BUILD_FLAGS}" \
    -os="${XC_OS}" \
    -arch="${XC_ARCH}" \
    -output "${ASSET_DIR}/{{.OS}}-{{.Arch}}/{{.Dir}}"

find pkg -mindepth 1 -type d -print0 | xargs -0 -I{} \
  zip -r "{}.zip" "{}"
