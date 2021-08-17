#!/usr/bin/env bash
set -e -u -o pipefail

repo_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd -P)"
cd "${repo_root}"

bazel build //:validators_go_proto
cp bazel-bin/validators_go_proto_/github.com/mwitkow/go-proto-validators/validators.pb.go .
chmod 644 validators.pb.go

bazel build //test:validatortest_go_proto
cp bazel-bin/test/validatortest_go_proto_/github.com/mwitkow/go-proto-validators/test/*.pb.go ./test
chmod 644 test/*.pb.go

bazel build //examples:validator_examples_go_proto
cp bazel-bin/examples/validator_examples_go_proto_/github.com/mwitkow/go-proto-validators/examples/*.pb.go ./examples
chmod 644 test/*.pb.go
