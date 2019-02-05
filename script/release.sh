#!/usr/bin/env bash

XC_ARCH=${XC_ARCH:-386 amd64}
XC_OS=${XC_OS:-linux darwin windows}
BUILD_FLAGS="-X \"main.Version=${VERSION:-v0.0.1}\""

rm -rf pkg/*
gox \
    -ldflags="${BUILD_FLAGS}" \
    -os="${XC_OS}" \
    -arch="${XC_ARCH}" \
    -output "pkg/{{.OS}}-{{.Arch}}/{{.Dir}}"
