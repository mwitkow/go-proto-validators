#!/usr/bin/env bash
set -e -u -o pipefail

repo_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd -P)"
cd "${repo_root}"

bazel build //...

bazel test //...
