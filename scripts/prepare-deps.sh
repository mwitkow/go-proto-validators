#!/usr/bin/env bash

set -e -u -o pipefail

PROJECT_DIR="$(dirname "${BASH_SOURCE[0]}")/.."
cd "${PROJECT_DIR}"
export PROJECT_DIR

source scripts/includes/deps.sh

install_protobuf
setup_proto_deps
