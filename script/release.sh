#!/usr/bin/env bash

XC_ARCH=${XC_ARCH:-amd64}
XC_OS=${XC_OS:-linux darwin}
VERSION=${VERSION:-v0.0.0}
BUILD_FLAGS="-X \"main.Version=${VERSION}\""
ASSET_DIR="${ASSET_DIR:=pkg}"
echo "ASSET_DIR:${ASSET_DIR}"

rm -vrf "${ASSET_DIR:?}/"*

gox \
    -ldflags="${BUILD_FLAGS}" \
    -os="${XC_OS}" \
    -arch="${XC_ARCH}" \
    -output "${ASSET_DIR}/{{.OS}}-{{.Arch}}/{{.Dir}}"

type sha256sum awk
compress () {
  cat << EOS
    {
      TARGET_PATH="@"
      DIRNAME="\$(basename "\${TARGET_PATH}")"
      TARBALL="\${DIRNAME}.tar.gz"
      cd "\${TARGET_PATH}/.."
      tar cfvz "\${TARBALL}" "\${DIRNAME}"
      FINGERPRINT="\$(sha256sum -b "\${TARBALL}" | awk '{print \$1}')"
      echo "\${FINGERPRINT}" > "\${DIRNAME}-\${FINGERPRINT:0:7}.txt"
    }
EOS
}

find "${ASSET_DIR}" -mindepth 1 -type d -print0 | xargs -0 -I @ bash -c "$(compress)"

ls -la "${ASSET_DIR}"
