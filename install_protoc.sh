#!/bin/bash
# Copyright 2016 Michal Witkowski. All Rights Reserved.
# See LICENSE for licensing terms.
#
# This script installs protobuf compiler `protoc` into PATH.

version=${PROTOBUF_VERSION:-"3.0.0-beta-2"}
dst_dir="${HOME}/soft/protobuf"

# Fail on issues.
set -e

echo "Downloading and installing protoc ${version}"

mkdir -p ${dst_dir}

wget https://github.com/google/protobuf/releases/download/${version}/protoc-${version}-linux-x86_64.zip -O ${dst_dir}/dist.zip

cd ${dst_dir}
unzip dist.zip

echo "Updating \$PATH and setting \$PROTOBUF_DIR"
export PROTOBUF_DIR=${dst_dir}
export PATH=${PATH}:${PROTOBUF_DIR}
