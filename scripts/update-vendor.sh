#!/bin/bash

set -eux

project_root="$(cd "$(cd "$( dirname "${BASH_SOURCE[0]}" )" && git rev-parse --show-toplevel)" >/dev/null 2>&1 && pwd)"
readonly project_root


readonly PROTOC_VERSION=3.20.1
readonly PROTOC_GEN_GO_VERSION=1.27.1
readonly PROTOC_GEN_GOGO_VERSION=1.3.2

readonly dependencies=(
  "github.com/gogo/protobuf/protoc-gen-gogo@v${PROTOC_GEN_GOGO_VERSION}"
)


readonly container="go-proto-validator-update-vendor"
readonly build_image="go-proto-validator-build"

build_dir=$(mktemp -d); readonly build_dir

cleanup() {
  local -r retval="$?"
  set +eu

  docker rm "${container}"
  docker rmi "${build_image}"

  rm -rf "${build_dir}"

  set +x

  if [[ "${success:-no}" == "yes" ]]; then
    echo
    echo "*******************************"
    echo "*** Vendor Update Succeeded ***"
    echo "*******************************"
    echo
  else
    echo
    echo "****************************"
    echo "*** Vendor Update Failed ***"
    echo "****************************"
    echo
  fi

  exit "${retval}"
}
trap cleanup EXIT

cd "${project_root}"

git clone --depth 1 file://"${project_root}" "${build_dir}"

docker build \
  --tag "${build_image}" \
  --target "build-env" \
  - < "${project_root}"/Dockerfile

docker run \
  --rm \
  --interactive \
  --volume "${build_dir}":/build/ \
  "${build_image}" \
  /bin/sh - <<EOF
#!/bin/sh
set -eux


cd /build/


rm -rf \
  ./go.mod \
  ./go.sum \
  ./vendor/

go clean -cache
go mod init github.com/rakshasa/go-proto-validator

for dep in ${dependencies[@]}; do
  go get -u -v "\${dep}"
done

go mod tidy -v
go mod vendor -v


set +x
echo
echo "+----------------------+"
echo "| Vendor Files Created |"
echo "+----------------------+"
echo

EOF


rm -rf ./{go.mod,go.sum,vendor}

cp -r "${build_dir}"/{go.mod,go.sum,vendor} ./
 
success="yes"
