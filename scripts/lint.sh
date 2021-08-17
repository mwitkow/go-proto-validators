#!/usr/bin/env bash
set -e -u -o pipefail

repo_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd -P)"
cd "${repo_root}"

function check_diff() {
  if ! git diff --quiet; then
    git diff
    echo >&2 "---------------"
    echo >&2 "Linting failure. A diff was detected. See above for the exact diff"
    exit 1
  fi
}

echo "Verifying go.mod is up-to-date."
go mod tidy
check_diff

echo "Ensuring Bazel build configuration is up-to-date."
bazel run --run_under="cd $(pwd) && " @bazel_gazelle//cmd/gazelle -- update-repos -from_file=go.mod -to_macro=go_deps.bzl%go_repositories
bazel run //:gazelle
check_diff

echo "Ensuring generated files are up-to-date."
./scripts/regenerate.sh

check_diff
